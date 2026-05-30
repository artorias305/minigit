package commands

import (
	"os"
	"path/filepath"

	"github.com/artorias305/minigit/utils"
)

func HashObject(path string, object utils.Blob) (string, error) {
	subdirName := object.Hash[:2]
	fileName := object.Hash[2:]
	objectDir := filepath.Join(path, ".minigit", "objects", subdirName)
	if err := os.MkdirAll(objectDir, 0755); err != nil {
		return "", err
	}
	objectPath := filepath.Join(objectDir, fileName)
	if _, err := os.Stat(objectPath); err == nil {
		return object.Hash, nil
	}
	if err := os.WriteFile(objectPath, []byte(object.Data), 0644); err != nil {
		return "", err
	}
	return object.Hash, nil
}
