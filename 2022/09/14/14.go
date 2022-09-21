package main

import (
	"strings"
	"sync"
)

var m map[string]string
var sm sync.Map

func WriteToMap() {
	m = make(map[string]string, 0)
	m["k"] = "v"
	if v, ok := m["k"]; ok {
		strings.ToUpper(v)
	}
}

func WriteToSyncMap() {
	sm.Store("k", "v")
	if v, ok := sm.Load("v"); ok {
		strings.ToUpper(v.(string))
	}
}
