package main

import (
	"fmt"
	"os"

	"github.com/artorias305/minigit/commands"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "Usage: minigit <command>\n")
		return
	}
	cmd := os.Args[1]
	switch cmd {
	case "init":
		var path string
		if len(os.Args) >= 3 {
			path = os.Args[2]
		} else {
			path = "."
		}
		commands.InitRepo(path)
	}
}
