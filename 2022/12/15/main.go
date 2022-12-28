package main

import (
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func main() {
	svr := gin.Default()
	svr.POST("/grafana/alert", func(c *gin.Context) {
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			return
		}
		fmt.Println(string(body))
	})
	err := svr.Run("127.0.0.1:9000")
	if err != nil {
		return
	}
}
