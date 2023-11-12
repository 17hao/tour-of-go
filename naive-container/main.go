package main

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"runtime/debug"
	"syscall"

	"github.com/sirupsen/logrus"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func run() {
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags:   syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWUSER,
		Unshareflags: syscall.CLONE_NEWNS,
		UidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 0,
				HostID:      os.Getuid(),
				Size:        1,
			},
		},
		GidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 0,
				HostID:      os.Getgid(),
				Size:        1,
			},
		},
	}
	if err := cmd.Start(); err != nil {
		logrus.WithField("stack", string(debug.Stack())).Errorf("ERR: %v", err)
		return
	}
	if err := cmd.Wait(); err != nil {
		logrus.WithField("stack", string(debug.Stack())).Errorf("ERR: %v", err)
	}
}

func child() {
	logrus.Infof("child process pid=%d", os.Getpid())
	must(syscall.Sethostname([]byte("container-1")))

	rootfs := "/tmp/rootfs"
	proc := filepath.Join(rootfs, "/proc")
	must(syscall.Mount("proc", proc, "proc", 0, ""))
	must(syscall.Mount(rootfs, rootfs, "", syscall.MS_BIND|syscall.MS_PRIVATE|syscall.MS_REC, ""))

	oldRootFS := filepath.Join(rootfs, "old-rootfs")
	if _, err := os.Stat(oldRootFS); errors.Is(err, os.ErrNotExist) {
		must(os.Mkdir(oldRootFS, 0700))
	}
	must(syscall.PivotRoot(rootfs, oldRootFS))
	must(syscall.Chdir("/"))
	must(syscall.Unmount("/old-rootfs", syscall.MNT_DETACH))
	must(os.RemoveAll(oldRootFS))

	cmd := exec.Command("/bin/sh")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = []string{"PS1=-[naive-container]- # "}
	if err := cmd.Run(); err != nil {
		logrus.WithField("stack", string(debug.Stack())).Errorf("ERR: %v", err)
		return
	}
}

// su
// go build main.go
// ./main run
func main() {
	logrus.Infof("os.Args: %+v", os.Args)
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		logrus.WithField("stack", string(debug.Stack())).Error("unsupported command, usage: go run main.go run [command] <args>\n")
	}
}
