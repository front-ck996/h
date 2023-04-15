package csy

import (
	"fmt"
	"os"
)

func GetMainDir() {
	fmt.Println(os.Getwd())
}
