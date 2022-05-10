package main

import (
	"os"
	"path/filepath"
)

func main() {
	curDir, _ := os.Getwd()
	println(curDir)
	parentDir := filepath.Dir(curDir)
	println(parentDir)
}
