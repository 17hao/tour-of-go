package main

import (
	"errors"
	"fmt"
)

func testFmt(fmtInput string, v ...interface{}) string {
	return fmt.Sprintf(fmtInput, v...)
}

func main() {
	fmt.Println(testFmt("err: %s, key: %s", errors.New("mock err").Error(), "mock key"))
}
