package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
) 

type album struct{
	ID string `json:"id"`
	Title string `json:"title"`
}
var albums = []album{
	{ID:"", Title: "Harry Potter"},
	{ID:"", Title: "Atomic Habit"},
}


func AllAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func main() {
	router := gin.Default()
	router.GET("/albums", AllAlbums)
	router.Run("localhost:8081")
}