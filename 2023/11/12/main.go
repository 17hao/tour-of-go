package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/vishvananda/netlink"
	"golang.org/x/sys/unix"
)

func main() {
	routes, err := netlink.RouteList(nil, unix.AF_INET)
	if err != nil {
		logrus.Error(err)
		return
	}
	for _, route := range routes {
		fmt.Printf("%+v\n", route)
	}
}