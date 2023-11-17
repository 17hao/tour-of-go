package main

import (
	"bufio"
	"fmt"
	"net"
	"runtime/debug"

	"github.com/sirupsen/logrus"
)

func handleConnection(conn net.Conn) {
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	line, err := reader.ReadString('\n')
	if err != nil {
		logrus.Error(err)
	}
	fmt.Printf("%s", line)

	if _, err := fmt.Fprint(conn, line); err != nil {
		logrus.Error(err)
	}
	if _, err := writer.Write([]byte(line)); err != nil {
		logrus.Error(err)
	}
	if err := writer.Flush(); err != nil {
		logrus.Error(err)
	}
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			logrus.Error(string(debug.Stack()))
		}
	}()
	srv, err := net.Listen("tcp", "localhost:9999")
	if err != nil {
		logrus.Error(err)
		return
	}
	for {
		conn, err := srv.Accept()
		if err != nil {
			logrus.Error(err)
			return
		}
		go handleConnection(conn)
	}
}
