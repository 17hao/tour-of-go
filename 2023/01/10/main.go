package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/ghodss/yaml"
)

type cfg1 struct {
	L1 struct {
		L2 struct {
			L3 string `yaml:"l3"`
		}
	}
}

type cfg2 struct {
	M1 struct {
		M2 struct {
			M3 string `yaml:"m3"`
		}
	}
}

type cfg3 struct {
	N1 struct {
		N2 struct {
			N3 []string `yaml:"n3"`
		}
	}
}

func main() {
	wd, _ := os.Getwd()
	println(wd)
	fp := filepath.Join(wd, "conf/test.yaml")
	f, err := ioutil.ReadFile(fp)
	if err != nil {
		panic(err)
	}

	c1 := cfg1{}
	_ = yaml.Unmarshal(f, &c1)
	fmt.Printf("%+v\n", c1)

	c2 := cfg2{}
	_ = yaml.Unmarshal(f, &c2)
	fmt.Printf("%+v\n", c2)

	c3 := cfg3{}
	_ = yaml.Unmarshal(f, &c3)
	fmt.Printf("%+v\n", c3)
}
