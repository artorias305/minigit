package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"os"
)

type Blob struct {
	Data string
	Hash string
	Size int
}

func NewBlob(data string) *Blob {
	hash := hashString(data)
	return &Blob{
		Data: data,
		Hash: hash,
		Size: len(data),
	}
}

func NewBlobFromFile(filePath string) (*Blob, error) {
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return NewBlob(string(bytes)), nil
}

func hashString(str string) string {
	hashBytes := sha1.Sum([]byte(str))
	hashString := hex.EncodeToString(hashBytes[:])
	return hashString
}
