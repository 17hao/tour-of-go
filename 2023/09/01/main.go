package main

import (
	"errors"
	"fmt"
	"log"
	"time"
)

func f1() error {
	return errors.New("f1 err")
}

func f2() error {
	return nil
}

func f3() error {
	var err error
	if err := f1(); err != nil {
		log.Fatal("f1: " + err.Error())
	}
	if err := f2(); err != nil {
		log.Fatal("f2: " + err.Error())
	}
	return err

	// if err := f1(); err != nil {
	// 	log.Fatal("f1: " + err.Error())
	// }
	// if err := f2(); err != nil {
	// 	log.Fatal("f2: " + err.Error())
	// }
	// return nil
}

func main() {
	// err := f3()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	t := time.Now()
	fmt.Printf("%+v\n", t)
	t = t.Add(-8 * time.Hour)
	fmt.Printf("%+v\n", t)
}
