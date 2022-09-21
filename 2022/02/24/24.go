package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	viperDemo()
}

func viperDemo() {
	viper.AddConfigPath("conf")
	viper.SetConfigName("test")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w\n", err))
	}
	//viper.ReadConfig(os.Stdout)
	fmt.Println(viper.Get("goroot"))
	fmt.Println(viper.Get("goproxy"))
}
