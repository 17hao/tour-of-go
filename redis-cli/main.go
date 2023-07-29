package main

import (
	"bufio"
	"fmt"
	"github.com/sirupsen/logrus"
	"net"
	"os"
	"time"
)

func main() {
	cli, err := NewClient("localhost:6379", "123456")
	if err != nil {
		logrus.Errorf("%v", err)
		os.Exit(1)
	}

	set, err := cli.Set("k2", "v2")
	if err != nil {
		logrus.Errorf("%v", err)
	}
	fmt.Printf("set: %v\n", set)

	get, err := cli.Get("k3")
	if err != nil {
		logrus.Errorf("%v", err)
	}
	fmt.Printf("get: %v\n", get)

	defer cli.Conn.Close()
}

type RedisCli struct {
	Addr     string
	Password string
	net.Dialer
	net.Conn
}

func (c *RedisCli) auth() error {
	w := bufio.NewWriter(c.Conn)
	w.WriteString("*5\r\n")
	w.WriteString("$5\r\nhello\r\n")
	w.WriteString("$1\r\n3\r\n")
	w.WriteString("$4\r\nauth\r\n")
	w.WriteString("$7\r\ndefault\r\n")
	w.WriteString(fmt.Sprintf("$%d\r\n%s\r\n", len(c.Password), c.Password))
	w.Flush()

	reader := bufio.NewReader(c.Conn)
	line, _, err := reader.ReadLine()
	if err != nil {
		return err
	}
	logrus.Debugf("%s", line)
	return nil
}

func NewClient(addr string, passwd string) (*RedisCli, error) {
	res := &RedisCli{
		Addr:     addr,
		Password: passwd,
	}
	res.Timeout = 3 * time.Second
	conn, err := res.Dial("tcp", res.Addr)
	if err != nil {
		return nil, err
	}
	res.Conn = conn
	if err = res.auth(); err != nil {
		return nil, err
	}
	return res, nil
}

func (c *RedisCli) Set(key string, value string) (string, error) {
	w := bufio.NewWriter(c.Conn)
	w.WriteString("*3\r\n")
	w.WriteString("$3\r\nset\r\n")
	w.WriteString(fmt.Sprintf("$%d\r\n%s\r\n", len(key), key))
	w.WriteString(fmt.Sprintf("$%d\r\n%s\r\n", len(value), value))
	w.Flush()

	reader := bufio.NewReader(c.Conn)
	line, _, err := reader.ReadLine()
	if err != nil {
		return "", err
	}
	return string(line), nil
}

func (c *RedisCli) Get(key string) (string, error) {
	w := bufio.NewWriter(c.Conn)
	w.WriteString("*2\r\n")
	w.WriteString("$3\r\nget\r\n")
	w.WriteString(fmt.Sprintf("$%d\r\n%s\r\n", len(key), key))
	w.Flush()

	reader := bufio.NewReader(c.Conn)
	line, _, err := reader.ReadLine()
	if err != nil {
		return "", err
	}
	if line[0] == '$' {
		return string(line[1:]), nil
	}
	return string(line), nil
}
