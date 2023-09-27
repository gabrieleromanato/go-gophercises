package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type File struct {
	Name  string
	IsDir bool
}

func listDirRecursively(dir string) ([]File, error) {
	var files []File
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		files = append(files, File{path, info.IsDir()})
		return nil
	})
	if err != nil {
		return files, err
	}
	return files, nil
}

func parsePattern(pattern string, f func(string) string) (string, error) {
	validPatterns := []string{"(filename)", "(counter)"}
	isValid := true
	for _, validPattern := range validPatterns {
		if !strings.Contains(pattern, validPattern) {
			isValid = false
			break
		}
	}
	if !isValid {
		return "", fmt.Errorf("Invalid pattern")
	}
	return f(pattern), nil
}

func renameFilesByPattern(files []File, pattern string) bool {
	for i, file := range files {
		if file.IsDir {
			continue
		}
		newName, _ := parsePattern(pattern, func(pattern string) string {
			s := strings.ReplaceAll(pattern, "(filename)", file.Name)
			s = strings.ReplaceAll(s, "(counter)", fmt.Sprintf("%d", i))
			return s
		})
		err := os.Rename(file.Name, newName)
		if err != nil {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("Please specify a pattern for renaming files:")
	fmt.Println("Available variables: (filename) (counter)")
	var pattern string
	fmt.Scanln(&pattern)

	files, err := listDirRecursively("./sample")
	if err != nil {
		panic(err)
	}
	if !renameFilesByPattern(files, pattern) {
		panic("Failed to rename files")
	}
	fmt.Println("Done")
	os.Exit(0)

}
