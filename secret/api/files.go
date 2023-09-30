package api

import (
	"os"
)

func CreateSecretDirIfNotExists(directoryPath string) string {
	if _, err := os.Stat(directoryPath); err != nil {
		os.Mkdir(directoryPath, 0755)
	}
	return directoryPath
}

func FileExists(filepath string) bool {
	if _, err := os.Stat(filepath); err != nil {
		return false
	}
	return true
}

func WriteSecretToFile(secret string, filepath string) bool {
	f, err := os.Create(filepath)
	if err != nil {
		return false
	}
	defer f.Close()
	f.WriteString(secret)
	return true
}

func ReadSecretFromFile(filepath string) (string, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer f.Close()
	fi, err := f.Stat()
	if err != nil {
		return "", err
	}
	size := fi.Size()
	data := make([]byte, size)
	f.Read(data)
	return string(data), nil
}
