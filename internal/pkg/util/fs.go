package util

import (
	"path"
	"path/filepath"
	"runtime"
)

var (
	_, currentFile, _, _ = runtime.Caller(0)
	currentDir           = filepath.Dir(currentFile)
	projectRoot          = path.Join(currentDir, "..", "..", "..")
)

// GetProjectRoot returns absolute path of project root.
func GetProjectRoot() string {
	return projectRoot
}
