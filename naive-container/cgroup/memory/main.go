package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	cgorupRoot := "/sys/fs/cgroup"
	systemSlice := filepath.Join(cgorupRoot, "system.slice")
	path := filepath.Join(systemSlice, "naive-container.service")
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		must(os.Mkdir(path, 0700))
	}
	must(os.WriteFile(filepath.Join(path, "memory.max"), []byte("3M"), 0700))
	must(os.WriteFile(filepath.Join(path, "cgroup.procs"), []byte(strconv.FormatInt(int64(os.Getpid()), 10)), 0700))

	total := 0
	for i := 1; ; i++ {
		len := 1024 * 1024
		memory := make([]byte, len)
		// 为byte数组赋值，确保真正分配了内存
		for i := 0; i < len; i++ {
			memory[i] = 0
		}
		total += len
		fmt.Printf("[%d] allocate %d-bytes, total=%d\n", i, len, total)
		time.Sleep(500 * time.Millisecond)
	}
}
