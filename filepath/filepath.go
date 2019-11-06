// Package filepath provides utility functions for retrieving file paths
package filepath

import (
	"path"
	"path/filepath"
	"runtime"
)

// RelativePath returns full path of the file
// in the same directory as source code file
// that is calling this function.
// This is file system agnostic.
func RelativePath(fileName string) string {
	_, thisFile, _, ok := runtime.Caller(1)
	if !ok {
		panic("No caller information")
	}
	return filepath.FromSlash(path.Dir(thisFile) + "/" + fileName)
}
