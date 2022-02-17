package utils

import "os"

func PathExists(path string) (bool, error) { //检查路径是否存在
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
