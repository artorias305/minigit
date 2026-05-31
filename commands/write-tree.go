package commands

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"

	"github.com/artorias305/minigit/utils"
)

func WriteTree() error {
	hash, err := writeTreeDir(".")
	if err != nil {
		return err
	}
	fmt.Println(hash)
	return nil
}

func writeTreeDir(dir string) (string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return "", err
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Name() < entries[j].Name()
	})

	var content bytes.Buffer

	for _, entry := range entries {
		name := entry.Name()

		if name == ".minigit" {
			continue
		}

		fullPath := filepath.Join(dir, name)

		if entry.IsDir() {
			treeHash, err := writeTreeDir(fullPath)
			if err != nil {
				return "", err
			}
			if err := appendTreeEntry(&content, "40000", name, treeHash); err != nil {
				return "", err
			}
			continue
		}

		info, err := entry.Info()
		if err != nil {
			return "", err
		}
		if !info.Mode().IsRegular() {
			continue
		}

		blob, err := utils.NewBlobFromFile(fullPath)
		if err != nil {
			return "", err
		}
		hash, err := HashObject(".", *blob)
		if err != nil {
			return "", err
		}

		mode := "100644"
		if info.Mode()&0111 != 0 {
			mode = "100755"
		}
		if err := appendTreeEntry(&content, mode, name, hash); err != nil {
			return "", err
		}
	}

	return writeObject("tree", content.Bytes())
}

func appendTreeEntry(buf *bytes.Buffer, mode, name, hash string) error {
	hashBytes, err := hex.DecodeString(hash)
	if err != nil {
		return err
	}
	buf.WriteString(mode)
	buf.WriteByte(' ')
	buf.WriteString(name)
	buf.WriteByte(0)
	buf.Write(hashBytes)
	return nil
}

func writeObject(kind string, content []byte) (string, error) {
	header := kind + " " + strconv.Itoa(len(content)) + "\000"
	canonical := append([]byte(header), content...)

	sum := sha1.Sum(canonical)
	hash := hex.EncodeToString(sum[:])

	subdir := hash[:2]
	file := hash[2:]
	objDir := filepath.Join(".minigit", "objects", subdir)
	if err := os.MkdirAll(objDir, 0755); err != nil {
		return "", err
	}

	objPath := filepath.Join(objDir, file)
	if _, err := os.Stat(objPath); err != nil {
		return hash, nil
	}
	if err := os.WriteFile(objPath, canonical, 0644); err != nil {
		return "", err
	}
	return hash, nil
}
