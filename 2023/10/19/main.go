package main

import (
	"fmt"
	"os"
	"os/exec"
)

func run() {
	fmt.Printf("running...\n")
	cmd := exec.Cmd{}
	cmd.Path = "/proc/self/exe"
	cmd.Args = []string{"child"}
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	if err := cmd.Start(); err != nil {
		fmt.Printf("ERROR %v", err)
		os.Exit(1)
	}
}

func child() {
	fmt.Printf("child")
}

func main() {
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		fmt.Print("unsupported command")
	}
}
