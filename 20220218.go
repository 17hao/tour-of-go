package main

import (
	"encoding/json"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/emirpasic/gods/maps/treemap"
	"github.com/umpc/go-sortedmap"
	"os"
	"strings"
)

func main() {
	//treeMapDemo()
	//sortedMapDemo()
	//jsonDemo()
	sonicDemo()
}

type Employee struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
}

func sonicDemo() {
	bytes, _ := sonic.Marshal(Employee{1, "test", 2})
	os.Stdout.Write(bytes)
}

func jsonDemo() {
	//bytes, _ := json.Marshal(Employee{2, "test", 1})
	////fmt.Println(bytes)
	//os.Stdout.Write(bytes)
	//var e Employee
	//json.Unmarshal(bytes, &e)
	//fmt.Println(e.Name)

	bytes := []byte(`{"id": 1, "name": "hello", "age": 6}`)
	var employee interface{}
	json.Unmarshal(bytes, &employee)
	for k, v := range employee.(map[string]interface{}) {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, " is string", vv)
		case float64:
			fmt.Println(k, " is float64", vv)
		case int:
			fmt.Println(k, " is int", vv)
		default:
			fmt.Print(k, " is of a type I don't know how to handle.")
		}
	}
}

func treeMapDemo() {
	tm := treemap.NewWith(func(i, j interface{}) int {
		return strings.Compare(i.(string), j.(string))
	})
	tm.Put("zabbix", "1")
	tm.Put("golang", "2")
	tm.Put("docker", "3")
	iterator := tm.Iterator()
	for iterator.Next() {
		fmt.Println(iterator.Key(), " -> ", iterator.Value())
	}
}

func sortedMapDemo() {
	sm := sortedmap.New(4, func(i, j interface{}) bool {
		return strings.Compare(i.(string), j.(string)) > 0
	})
	sm.Insert("zabbix", "1")
	sm.Insert("golang", "2")
	sm.Insert("docker", "3")
	itr, _ := sm.IterCh()
	for r := range itr.Records() {
		fmt.Println(r)
	}
}
