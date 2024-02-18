package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type s1 struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type s2 struct {
	ID       *int64  `json:"id"`
	Name     *string `json:"name"`
	TenantID *int64  `json:"tenantID"`
}

func main() {
	m := s1{
		ID:   1,
		Name: "s1",
	}
	b, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("s1=%s\n", string(b))

	var n s2
	if err := json.Unmarshal(b, &n); err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("%+v", n)
}
