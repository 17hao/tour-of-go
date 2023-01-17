package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	err := os.Setenv("HTTPS_PROXY", "127.0.0.1:8889")
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Get("https://www.okx.com/api/v5/market/tickers?instType=SWAP&uly=BTC-USDT")
	if err != nil {
		log.Fatal(err)
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d\n", resp.StatusCode)
	fmt.Printf("%s\n", string(respBody))
}
