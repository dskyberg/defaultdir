package defaultdir

import "os"

// isDir combines os.Stat and FileInfo.IsDir
func isDir(value string) bool {
	if s, err := os.Stat(value); os.IsNotExist(err) {
		return false
	} else if err != nil {
		return false
	} else if !s.IsDir() {
		return false
	}
	return true
}
