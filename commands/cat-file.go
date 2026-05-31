package commands

import (
	"fmt"
	"os"
)

func CatFile(hash string) (string, error) {
	if len(hash) != 40 {
		return "", fmt.Errorf("invalid hash length: %d", len(hash))
	}
	prefix := hash[:2]
	suffix := hash[2:]
	path := ".minigit/objects/" + prefix + "/" + suffix
	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
