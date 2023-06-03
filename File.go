package csy

import (
	"log"
	"net/http"
	"os"
	"strings"
)

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

// IsFile is_file()
func IsFile(filename string) bool {
	_, err := os.Stat(filename)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

// IsDir is_dir()
func IsDir(filename string) (bool, error) {
	fd, err := os.Stat(filename)
	if err != nil {
		return false, err
	}
	fm := fd.Mode()
	return fm.IsDir(), nil
}

// ISIMG
func IsImg(filename string) (bool, error) {
	// Open File
	f, err := os.Open(filename)
	if err != nil {
		return false, err
	}
	defer f.Close()

	buffer := make([]byte, 512)

	_, err = f.Read(buffer)
	if err != nil {
		return false, err
	}

	contentType := http.DetectContentType(buffer)
	log.Println(contentType)
	if !strings.HasPrefix(contentType, "im") {
		return false, nil
	}
	return true, nil
}
