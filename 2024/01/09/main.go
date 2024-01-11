package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

type s struct {
	id int64
}

func f1() {
	m := map[string]map[string][]*s{
		"k1": {
			"k2": {
				{
					id: 1,
				},
			},
		},
	}
	fmt.Printf("%+v", m["k1"]["k2"][0])
}

func main() {
	f1()

	key := "private-key"
	h := hmac.New(sha256.New, []byte(key))

	message := "message"
	h.Write([]byte(message))
	fmt.Print(base64.StdEncoding.EncodeToString(h.Sum(nil)))
}
