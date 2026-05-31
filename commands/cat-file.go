package commands

import "os"

func CatFile(hash string) (string, error) {
	prefix := hash[:2]
	suffix := hash[2:]
	path := ".minigit/objects/" + prefix + "/" + suffix
	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
