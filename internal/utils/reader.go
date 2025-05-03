package utils

import (
	"encoding/json"
	"math/rand"
	"os"
)

func ReadJSONFile(filePath string, data interface{}) (interface{}, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}
	fileSize := fileInfo.Size()

	byteValue := make([]byte, fileSize)
	_, err = file.Read(byteValue)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(byteValue, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func RandomIndex(max int) int {
	return rand.Intn(max)
}
