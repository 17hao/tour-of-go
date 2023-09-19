package main

import (
	"fmt"
	"sync"
	"time"
)

type db struct {
	store map[string]string
}

func (d *db) Add(k, v string) {
	d.store[k] = v
}

func (d *db) Get(k string) string {
	if v, ok := d.store[k]; ok {
		return v
	}
	return ""
}

type s struct {
	mutex sync.RWMutex
	db
}

func write(s1 *s) {
	s1.mutex.Lock()
	defer s1.mutex.Unlock()
	s1.db.Add("k2", "v2")
	fmt.Printf("write to db\n")
	// time.Sleep(3 * time.Second)
}

func read(s1 *s) {
	s1.mutex.RLock()
	defer s1.mutex.RUnlock()
	time.Sleep(3 * time.Second)
	v := s1.db.Get("k2")
	fmt.Printf("k2=%s\n", v)
}

func main() {
	s1 := &s{
		mutex: sync.RWMutex{},
		db:    db{store: map[string]string{}},
	}
	s1.db.Add("k1", "v1")
	fmt.Printf("k1=%s\n", s1.db.Get("k1"))


	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(s1 *s) {
		defer wg.Done()
		read(s1)
	}(s1)

	write(s1)

	wg.Wait()
}
