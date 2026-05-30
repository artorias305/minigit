package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"os"
	"strconv"
)

type Blob struct {
	Data string
	Hash string
	Size int
}

func NewBlob(data string) *Blob {
	size := len(data)
	canonical := "blob " + strconv.Itoa(size) + "\x00" + data
	hash := hashString(canonical)
	return &Blob{
		Data: data,
		Hash: hash,
		Size: size,
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
