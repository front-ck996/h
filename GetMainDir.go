package csy

import (
	"os"
)

func GetMainDir() string {
	dir, _ := os.Getwd()
	return dir
}
