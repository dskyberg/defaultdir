// Copyright 2017 David Skyberg.  All rights reserved.
// Use of this source code is governed by the GNU v3
// license that can be found in the LICENSE file.

package defaultdir

import (
	"errors"
	"os"
	"path"
	"path/filepath"
)

// Spec holds internal state for the chain
type Spec struct {
	_dir  *string
	_base *string
	_err  error
}

// New initializes and returns a *Spec
func New() *Spec {
	s := Spec{}
	return &s
}

// Base sets the name of the base folder for methods such as cwd.
func (s *Spec) Base(base string) *Spec {
	if s.breakChain() {
		return s
	}
	s._base = &base
	return s
}

// ClearBase removes any base that was previously set
func (s *Spec) ClearBase() *Spec {
	if s.breakChain() {
		return s
	}
	s._base = nil
	return s
}

// Cwd evaluates wheether the current directory is valid.  Typically, you
// want to set a base before calling cwd
func (s *Spec) Cwd() *Spec {
	if s.breakChain() {
		return s
	}

	var root string
	var err error

	if root, err = os.Getwd(); err != nil {
		s._err = err
		return s
	}
	var dir string
	if s._base != nil {
		dir = path.Join(root, *s._base)
	} else {
		dir = root
	}

	if isDir(dir) {
		s._dir = &dir
	}
	return s
}

// Bin evaluates wheether the directory the app is running from is valid.
// Typically, you want to set a base before calling cwd
func (s *Spec) Bin() *Spec {
	if s.breakChain() {
		return s
	}

	var root string
	var err error

	if root, err = filepath.Abs(filepath.Dir(os.Args[0])); err != nil {
		s._err = err
		return s
	}

	var dir string
	if s._base != nil {
		dir = path.Join(root, *s._base)
	} else {
		dir = root
	}

	if isDir(dir) {
		s._dir = &dir
	}
	return s
}

// Env looks for key in the env variables
func (s *Spec) Env(key string) *Spec {
	if s.breakChain() {
		return s
	}

	var root string
	var exists bool

	if root, exists = os.LookupEnv(key); !exists {
		// The ENV key does not exist
		return s
	}

	var dir string
	if s._base != nil {
		dir = path.Join(root, *s._base)
	} else {
		dir = root
	}

	if isDir(dir) {
		s._dir = &dir
	}
	return s
}

// Dir returns the directory, or the resulting error
func (s *Spec) Dir() (*string, error) {
	if s._err != nil {
		return nil, s._err
	}
	if s._dir == nil {
		return nil, errors.New("No directory was found")
	}
	return s._dir, nil
}

// breakChain determines whether to keep going or break the chain.
func (s *Spec) breakChain() bool {
	if s._dir != nil || s._err != nil {
		return true
	}
	return false
}
