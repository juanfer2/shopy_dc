package utilities

import (
	"path/filepath"
	"runtime"
)

func GetRootFolder() string {
	_, b, _, _ := runtime.Caller(0)

	return filepath.Join(filepath.Dir(b), "../../..")
}

func GetRootStaticFiles() string {
	return GetRootFolder() + "/tmp"
}
