package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func download(ctx *gin.Context) {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		logrus.Error(err)
		return
	}
	var req struct {
		ID int64 `json:"id"`
	}
	if err := json.Unmarshal(body, &req); err != nil {
		logrus.Error(err)
		return
	}
	logrus.Infof("%+v", req)

	var resp = struct {
		Name string `json:"name"`
		City string `json:"city"`
	}{
		Name: "aaa",
		City: "HZ",
	}
	b, err := json.Marshal(&resp)
	if err != nil {
		logrus.Error(err)
		return
	}
	ctx.Writer.Header().Set("Content-Disposition", "attachment; filename=0205-1.json")
	if _, err := ctx.Writer.Write(b); err != nil {
		logrus.Error(err)
		return
	}
}

func main() {
	server := gin.Default()
	server.POST("/download", download)
	if err := server.Run("127.0.0.1:9999"); err != nil {
		logrus.Error(err)
		return
	}
}
