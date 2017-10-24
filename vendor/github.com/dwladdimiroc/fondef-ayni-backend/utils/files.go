package utils

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func CreateFolder(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.Mkdir(path, os.FileMode(0755))
		if err != nil {
			return err
		}
	}

	return nil
}

func assignFilename(dir string, filename string, index int) string {
	base := filename
	ext := filepath.Ext(filename)

	var name string
	var newFilename string
	if index != 0 {
		name = strings.Replace(base, ext, "", -1)
		newFilename = name + "(" + strconv.Itoa(index) + ")" + ext
	} else {
		newFilename = filename
	}

	if _, err := os.Stat(filepath.Join(dir, newFilename)); os.IsNotExist(err) {
		return newFilename
	} else {
		newFilename = assignFilename(dir, filename, index+1)
	}

	return newFilename
}

func CreateFilename(dir string, filename string) string {
	return assignFilename(dir, filename, 0)
}
