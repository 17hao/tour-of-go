package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func displayDirTree(path string) {
	walkErr := filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			controllersbytes, readErr := os.ReadFile(filepath.Join(path, "cgroup.controllers"))
			if readErr != nil {
				fmt.Printf("%+v", readErr)
				os.Exit(1)
			}
			fmt.Printf("%s\n\t%s", path, string(controllersbytes))
			if len(controllersbytes) == 0 {
				return nil
			}
			pidsbytes, pidsErr := os.ReadFile(filepath.Join(path, "cgroup.procs"))
			if pidsErr != nil {
				fmt.Printf("%+v", pidsErr)
				os.Exit(1)
			}
			pids := strings.Split(string(pidsbytes), "\n")
			if len(pids) > 0 {
				fmt.Printf("\t%v\n", pids)
			}
		}

		return nil
	})
	if walkErr != nil {
		fmt.Printf("%+v", walkErr)
		os.Exit(1)
	}
}

func main() {
	flag.Parse()

	if len(flag.Args()) == 0 {
		fmt.Print("Usage: ./main <cgroup-dir-path>")
		return
	}

	path := flag.Args()[0]
	displayDirTree(path)
}
