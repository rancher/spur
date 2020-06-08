package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if err := Cleanup("."); err != nil {
		log.Fatal(err)
	}
}

// GeneratedString is the matching substring used when removing files
var GeneratedString = "zz_generated"

// Cleanup removes generated files
func Cleanup(path string) error {
	return filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if path != "." && strings.HasPrefix(info.Name(), ".") {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		if info.Name() == "vendor" && info.IsDir() {
			return filepath.SkipDir
		}
		if strings.Contains(path, GeneratedString) {
			fmt.Printf("Removing %s\n", path)
			if err := os.Remove(path); err != nil {
				return err
			}
		}

		return nil
	})
}
