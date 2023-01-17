package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type placeOrderRequest struct {
	InstID  string `json:"instId"`       // instrument ID, e.g. BTC-USDT
	TdMode  string `json:"tdMode"`       // trade mode, non-margin mode cash
	Side    string `json:"side"`         // order side, buy or sell
	OrdType string `json:"ordType"`      // order type, market/limit/post_only...
	Sz      string `json:"sz"`           // quantity to buy
	Px      string `json:"px,omitempty"` // order price
}

func main() {
	// TODO read config from local file
	apiKey := ""
	secretKey := ""
	timestamp := time.Now().UTC().Format("2006-01-02T15:04:05.999Z")
	passphrase := ""

	reqBody := placeOrderRequest{
		InstID:  "BTC-USDT",
		TdMode:  "cash",
		Side:    "buy",
		OrdType: "limit",
		Sz:      "10",
		Px:      "100",
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		log.Fatal(err)
	}
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(timestamp + "POST" + "/api/v5/trade/order" + string(body)))
	sign := base64.StdEncoding.EncodeToString(h.Sum(nil))

	req, err := http.NewRequest("POST", "https://www.okx.com/api/v5/trade/order", bytes.NewBuffer(body))
	if err != nil {
		log.Fatalf("%v", err)
	}
	req.Header.Set("OK-ACCESS-KEY", apiKey)
	req.Header.Set("OK-ACCESS-SIGN", sign)
	req.Header.Set("OK-ACCESS-TIMESTAMP", timestamp)
	req.Header.Set("OK-ACCESS-PASSPHRASE", passphrase)
	req.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	err = os.Setenv("HTTPS_PROXY", "127.0.0.1:8889")
	if err != nil {
		log.Fatalf("setenv err: %v", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d\n", resp.StatusCode)
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", respBody)
}
