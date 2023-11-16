package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"syscall"
	"time"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

// Create new process in shell: (sleep 5; printf "bye") &
// pstree -apT
// monitor the thread count of a process: ps -o thcount <pid>
//
// https://docs.kernel.org/admin-guide/cgroup-v2.html#pid
// Note that PIDs used in this controller refer to TIDs, process IDs as used by the kernel.
func main() {
	cgorupRoot := "/sys/fs/cgroup"
	systemSlice := filepath.Join(cgorupRoot, "system.slice")
	path := filepath.Join(systemSlice, "naive-container.service")
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		must(os.Mkdir(path, 0700))
	}
	must(os.WriteFile(filepath.Join(path, "pids.max"), []byte("10"), 0700))
	must(os.WriteFile(filepath.Join(path, "cgroup.procs"), []byte(strconv.FormatInt(int64(os.Getpid()), 10)), 0700))

	for {
		pid, forkErr := syscall.ForkExec("/bin/bash", []string{}, nil)
		if forkErr != nil {
			fmt.Printf("fork failed, err=%v\n", forkErr)
			// os.Exit(1)
		}
		if pid == -1 {
			fmt.Println("fork failed")
		} else if pid == 0 {
			break
		} else {
			fmt.Printf("child pid=%d\n", pid)
		}
	}

	time.Sleep(50 * time.Second)
}
