package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float32 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "t1", Artist: "sqh", Price: 12.34},
}

func main() {
	//fmt.Println("hello, world!")
	router := gin.Default()
	router.GET("/albums", getAlbums)

	router.Run("localhost:8080")
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}
