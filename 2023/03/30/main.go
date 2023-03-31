package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"reflect"

	jsoniter "github.com/json-iterator/go"
)

func main() {
	// f1(nil)
	//var s []string = nil
	//f1(s)
	//
	//f2(f3)
	//
	//f4()

	//f5()

	f6()
}

func f1(val interface{}) {
	rv := reflect.ValueOf(val)
	println(rv.Kind())
	println(rv.IsNil())
}

func f2(f func(s string)) {
	f("")
}

func f3(s string) {

}

func f4() {
	var v interface{}
	err := jsoniter.UnmarshalFromString("null", &v)
	if err != nil {
		log.Fatal(err)
	}
	rv := reflect.ValueOf(v)
	fmt.Println(rv.Kind())
	fmt.Println(rv.Len())
	fmt.Println(v)
}

func f5() {
	s := "4HMc2FoxGEajacjw1jquIdvrvXiJ2Uc4xLyX7bbZ5mRLYzvjSVVfl1VOPutdmrJlMbah73kRRQop8fqllJJlTKsYLMfJneX+K72CFHoI/uXYpjRttswHIj2LcYW65TjWHoQV/0aO59bBqnDoVvP2kINVFHK/zmS3BLhK/FSrcfk0mDpkUnMWndkrY/BCSGjtM2jxqWmoj+2tYSuslz4Lc7Sw1J0JtKzsr+q2KSlifiacdMB3RAePqXqLyChWnnWbo171zdajziOYaErbBkNrDft7W0vnwa3jhnzsaXh8tEGE3K4p1qFAYqoF+y7xZjhTbqvzdb14rqVMmQkYaY71SXP6Ua/nHaZM3Ii3LQ2qqilS3vVXG6/aYOyJ6yE55WgYMrqdSEk06KNcdXeNWl9a10krsfZLcFd59+CtSydS8J0QwFjbHNup2xFnM1qrivckQHH52khTfLj9mV67Uy3TOd1FhzvW+JkGDyse91ZC8Q67Z0acicY5F0Ktm1U6D1qXtwiFuklm3m+7l0zKpA6S1GTb8tidVufcVO8AOK2Tn+nuoNLwgt//BxickmkcHUSzB1aAQHisyTLik5m1b58H5/eI00qbIcGPodaybd/95hebRgekmTR05arSodmJQS9sKzWmNmMXuQWG7xPV2nVVBI0qxOLPMY4XEMK+3pHhqVSJY3OGCYCIIM+PTrHAkL7qXtJU2NHkmTYYGDSP37R1cT2TcGCrH4H5KbNhdRTyD/ovb7XcrFZH+VnHuza8WiHmoEj94RgOjZB0g8Imd/N5kLh1jML9qL/lvLOCoKtMx46I25+JUWsZv1LZIVUmZSX35VnfSsBjUzCD0H7n51tyKMMFvOmDupbu6dOUl7liSzvTbsFxlMvRKGhuXrBtczERDh7fC/M8yWFuWNbGCjn/LOY8N63xIx6hnZxNa21CsSVRGRdXYbulP8mby1ofFbwZQ9l8yZbNoMm0KU5S8XBvSBt66PzZ27wbzhYt//zhVRsO0GLNfJi1jvZogGHV5XH7Fxi4LK7PT29GF/PovmWG+fmpo1vTKzUHBpfxrD0uuVQCNd48ryRnQGe+hKjtrLzOOeuo8njAsHAxwnssVvRjp2LhlMM75wCqQJCirVb/yxdF89+8U0z+OFj1Utlfhk0u41S3IYpfrjeeiBXRBjlNcis6DeSbx+yE22BXAsyOWRK7CIx11HFT3RNGUxBYxcWGRzudw87M0gJQ0wQ0RaXI4bgbPVWy0TvJRcjwykZL+OPUDn7Riq1MGJw0MovGFZytWzzRXvy5C37j5sA/4FioB4T5/TrTlE/zRomaGqjvKjDNPKQrvEhG+CRQ+bnwoWc5b1ypB4s+WMZB4aykRD2h2ecVlgSz/wbnhIpDv4PJZi7OV1nTTeHSI5taOkgnRjPeSmS0KUkP7GNown3mTC5Kr4nrVh8Otq6uYy9WeQKSLP7uZ+4MuARscALq3dTmOz/fNo/GJTSmJEBQw5FX5OJ0QnDRnEnbBMGuy+1/kagYcjCNCvYgHgtI5kNgI8TojBah"
	bytes, _ := base64.StdEncoding.DecodeString(s)
	fmt.Println(string(bytes))
}

func f6() {
	s := "{\"f1\":null}"
	record := make(map[string]interface{}, 0)
	_ = jsoniter.UnmarshalFromString(s, &record)
	fmt.Printf("%+v\n", record)

	bytes, _ := jsoniter.Marshal(record["f1"])
	fmt.Println(string(bytes))
}
