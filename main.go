package main

import (
	"fmt"
	"os"

	"github.com/artorias305/minigit/commands"
	"github.com/artorias305/minigit/utils"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "Usage: minigit <command>\n")
		return
	}
	cmd := os.Args[1]
	switch cmd {
	case "init":
		path := "."
		if len(os.Args) >= 3 {
			path = os.Args[2]
		}
		if err := commands.InitRepo(path); err != nil {
			fmt.Fprintf(os.Stderr, "init failed: %v\n", err)
			os.Exit(1)
		}
	case "hash-object":
		if len(os.Args) < 3 {
			fmt.Fprintf(os.Stderr, "Usage: minigit hash-object <file> [repo-path]\n")
			os.Exit(1)
		}
		filePath := os.Args[2]
		repoPath := "."
		if len(os.Args) >= 4 {
			repoPath = os.Args[3]
		}
		blob, err := utils.NewBlobFromFile(filePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "hash-object failed reading file: %v\n", err)
			os.Exit(1)
		}
		hash, err := commands.HashObject(repoPath, *blob)
		if err != nil {
			fmt.Fprintf(os.Stderr, "hash-object failed: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(hash)
	case "cat-file":
		if len(os.Args) != 3 {
			fmt.Fprintf(os.Stderr, "Usage: minigit cat-file <hash>\n")
			os.Exit(1)
		}
		hash := os.Args[2]
		content, err := commands.CatFile(hash)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cat-file failed: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(content)
	}
}
