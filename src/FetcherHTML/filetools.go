package main

import (
	"os"
	"path/filepath"
)

// DirExists returns true if the given path exists and is a directory.
func DirExists(filename string) bool {
	fileInfo, err := os.Stat(filename)
	return err == nil && fileInfo.IsDir()
}

// DirExists returns true if the given path exists and is a directory.
func FileExists(filename string) bool {
	fileInfo, err := os.Stat(filename)
	return err == nil && !fileInfo.IsDir()
}

func DeleteEmptyFile(saveFolder string) {
	walkDir, err := filepath.Abs(saveFolder)
	if err != nil {
		Error("error:", err)
		return
	}
	err = filepath.Walk(walkDir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			if info.Size() < 5120 {
				os.Remove(path)
				Info("remove file", path)
			}
		}
		if err != nil {
			Error("error:", err)
			return nil
		}
		return nil
	})
	err = filepath.Walk(walkDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			if info.Size() < 5120 {
				os.Remove(path)
				Info("remove dir", path)
			}
		}
		if err != nil {
			Error("error:", err)
			return nil
		}
		return nil
	})
	return
}
