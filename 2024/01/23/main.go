package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func mustSetenv(key, value string) {
	if err := os.Setenv(key, value); err != nil {
		panic(err)
	}
}

func main() {
	url := "https://raw.githubusercontent.com/gfwlist/gfwlist/master/gfwlist.txt"
	mustSetenv("http_proxy", "127.0.0.1:8889")
	mustSetenv("https_proxy", "127.0.0.1:8889")
	cli := http.Client{}
	resp, err := cli.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	content := strings.Builder{}
	scanner := bufio.NewScanner(resp.Body)
	for {
		if !scanner.Scan() {
			break
		}
		content.Write(scanner.Bytes())
	}
	res, err := base64.StdEncoding.DecodeString(content.String())
	if err != nil {
		fmt.Print(err)
		return
	}
	if ioErr := os.WriteFile("gfwlist", res, 0600); ioErr != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf(string(res))
}
