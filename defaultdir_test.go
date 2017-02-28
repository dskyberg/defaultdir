package defaultdir

import (
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"
)

const (
	testBase = "test_base"
	testKey  = "DEFAULTDIR_TEST_ENV_DIR"
)

var createdTestFolder = false

func TestStart(t *testing.T) {
	// Create a test base folder in the current working directory
	if err := os.Mkdir(testBase, 0777); err != nil {
		if !strings.HasSuffix(err.Error(), "file exists") {
			t.Log("Could not create test folder", err)
			t.FailNow()
		}
	} else {
		createdTestFolder = true
	}
}

func TestEmptyCwd(t *testing.T) {
	var dir *string
	var err error
	expected, _ := os.Getwd()

	dir, err = New().
		Cwd().
		Dir()

	if err != nil {
		t.Log("Expected nil error, got ", err)
		t.Fail()
	}

	if expected != *dir {
		t.Logf("Expected %s - got %s\n", expected, *dir)
		t.Fail()
	}
}

func TestCwdWithBase(t *testing.T) {
	root, _ := os.Getwd()
	expected := path.Join(root, testBase)

	var dir *string
	var err error
	dir, err = New().
		Base(testBase).
		Cwd().
		Dir()

	if err != nil {
		t.Log("Expected nil error, got ", err)
		t.Fail()
	}

	if expected != *dir {
		t.Logf("Expected %s - got %s\n", expected, *dir)
		t.Fail()
	}
}

func TestEmptyBin(t *testing.T) {
	expected, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	var dir *string
	var err error
	dir, err = New().
		Bin().
		Dir()

	if err != nil {
		t.Log("Expected nil error, got ", err)
		t.Fail()
	}

	if expected != *dir {
		t.Logf("Expected %s - got %s\nDid you remember to use -o?", expected, *dir)
		t.Fail()
	}
}

func TestBinWithBase(t *testing.T) {
	root, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	expected := path.Join(root, testBase)

	var dir *string
	var err error

	dir, err = New().
		Base(testBase).
		Bin().
		Dir()

	if err != nil {
		t.Log("Expected nil error, got ", err)
		t.Fail()
	}
	if dir == nil {
		t.Log("Expected bin with base directory")
		t.Fail()
	}

	if expected != *dir {
		t.Logf("Expected %s - got %s\nDid you remember to use -o?", expected, *dir)
		t.Fail()
	}
}

func TestBadEnv(t *testing.T) {
	var dir *string
	var err error
	dir, err = New().
		Env("bogus").
		Dir()

	if nil != dir {
		t.Logf("Expected nil - got %s\n", *dir)
		t.Fail()
	}

	if err == nil {
		t.Log("Expected error error, got nil")
		t.Fail()
	}
}

func TestEmptyEnv(t *testing.T) {
	expected, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	os.Setenv(testKey, expected)

	var dir *string
	var err error
	dir, err = New().
		Env(testKey).
		Dir()

	if err != nil {
		t.Log("Expected nil error, got ", err)
		t.Fail()
	}

	if dir == nil {
		t.Log("Expected a directory")
		t.Fail()
	}

	if expected != *dir {
		t.Logf("Expected %s - got %s\n", expected, *dir)
		t.Fail()
	}
}

func TestEnvWithBase(t *testing.T) {
	root, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	expected := path.Join(root, testBase)
	os.Setenv(testKey, root)

	var dir *string
	var err error
	dir, err = New().
		Base(testBase).
		Env(testKey).
		Dir()

	if err != nil {
		t.Log("Expected nil error, got ", err)
		t.Fail()
	}

	if dir == nil {
		t.Log("Expected a directory")
		t.Fail()
	}

	if expected != *dir {
		t.Logf("Expected %s - got %s\n", expected, *dir)
		t.Fail()
	}
}

func TestChain(t *testing.T) {
	root, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	expected := path.Join(root, testBase)

	var dir *string
	var err error

	dir, err = New().
		Base("bogus").
		Cwd().
		Env("bogus").
		Base(testBase).
		Bin().
		Dir()

	if err != nil {
		t.Log("Expected nil error, got ", err)
		t.Fail()
	}

	if dir == nil {
		t.Log("Expected a directory")
		t.Fail()
	}

	if expected != *dir {
		t.Logf("Expected %s - got %s\n", expected, *dir)
		t.Fail()
	}

}

func TestCleanup(t *testing.T) {
	// Create a test base folder in the current working directory
	if createdTestFolder {
		if err := os.Remove(testBase); err != nil {
			t.Log("Could not remove test folder", err)
			t.Fail()
		}
	}
}
