package commands

import "os"

func InitRepo(path string) error {
	repoPath := path + "/.minigit"
	if err := os.Mkdir(repoPath, 0755); err != nil {
		return err
	}
	if err := os.Mkdir(repoPath + "/objects", 0755); err != nil {
		return err
	}
	if err := os.Mkdir(repoPath + "/refs", 0755); err != nil {
		return err
	}
	file, err := os.Create(repoPath + "/HEAD")
	if err != nil {
		return err
	}
	file.Close()
	return nil
}
