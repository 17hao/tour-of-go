package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"github.com/sirupsen/logrus"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func run() {
	fmt.Printf("running...\n")
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	// cmd.Path = "/proc/self/exe"
	// cmd.Args = append([]string{"child"}, os.Args[2:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}
	if err := cmd.Start(); err != nil {
		logrus.Fatalf("ERR %v", err)
	}
}

func child() {
	fmt.Printf("child process, pid=%d\n", syscall.Getpid())
	fmt.Printf("running %+v\n", os.Args[2:])

	wd, wdErr := os.Getwd()
	if wdErr != nil {
		panic(wdErr)
	}
	logrus.Infof("workdir=%s", wd)
	tmpFSPath := "/tmp/my-tmpfs"
	if _, err := os.Stat(tmpFSPath); errors.Is(err, os.ErrExist) {
		must(os.Mkdir(tmpFSPath, 0777))
	}
	must(syscall.Mount(wd, tmpFSPath, "", syscall.MS_BIND, ""))
	must(syscall.Chroot(tmpFSPath))
	root := "/"
	must(syscall.Chdir(root))

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		logrus.Fatalf("ERR %v", err)
	}
	defer func() {
		must(syscall.Unmount(tmpFSPath, syscall.MS_BIND))
		must(syscall.Unlink(tmpFSPath))
	}()
}

func main() {
	logrus.Infof("os.Args: %+v", os.Args)
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		logrus.Fatal("unsupported command, usage: go run main.go run [command] <args>\n")
	}
}
