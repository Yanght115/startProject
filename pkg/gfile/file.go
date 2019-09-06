package gfile

import (
	"fmt"
	"os"

	"github.com/toolkits/file"
)

// CreateIfNotExist .
func CreateIfNotExist(filePath string) (*os.File, error) {
	var exist bool
	var err error
	if exist, err = file.IsExist(filePath); err != nil {
		fmt.Errorf("get file(dir) exist error", err)
		return nil, err
	}
	if exist {
		return os.Open(filePath)
	}
	return file.Create(filePath)
}
