package utils

import (
	"io/fs"
	"path/filepath"
)

// GetFilesName 获取文件夹下的所有文件名
func GetFilesName(dirPath string) ([]string, error) {
	var fileNames []string
	err := filepath.Walk(dirPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		fileNames = append(fileNames, filepath.Join(dirPath, info.Name()))
		return nil
	})
	return fileNames, err
}
