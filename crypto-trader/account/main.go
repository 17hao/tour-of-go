package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	// TODO read config from local file
	apiKey := ""
	secretKey := ""
	timestamp := time.Now().UTC().Format("2006-01-02T15:04:05.999Z")
	passphrase := ""

	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(timestamp + "GET" + "/api/v5/account/balance"))
	sign := base64.StdEncoding.EncodeToString(h.Sum(nil))

	//fmt.Println(sign)
	//fmt.Println(timestamp)

	req, err := http.NewRequest("GET", "https://www.okx.com/api/v5/account/balance", nil)
	if err != nil {
		log.Fatalf("%v", err)
	}
	req.Header.Set("OK-ACCESS-KEY", apiKey)
	req.Header.Set("OK-ACCESS-SIGN", sign)
	req.Header.Set("OK-ACCESS-TIMESTAMP", timestamp)
	req.Header.Set("OK-ACCESS-PASSPHRASE", passphrase)
	client := http.Client{}
	err = os.Setenv("HTTPS_PROXY", "127.0.0.1:8889")
	if err != nil {
		log.Fatalf("setenv err: %v", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Printf("status code=%d\n", resp.StatusCode)
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Printf("response body=%s\n", respBody)
	//fmt.Printf("%+v", *resp)
}
