package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	// Import the generated docs (replace 'project_1/web-service-gin' with your actual module name from go.mod)

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Blue Train Record API
// @version         1.0
// @description     This is a sample Gin server for managing record albums.
// @host      localhost:8080
// @BasePath  /

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()

	// Swagger route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}

// getAlbums godoc
// @Summary      Get all albums
// @Description  Responds with the list of all albums as JSON.
// @Produce      json
// @Success      200  {array}  album
// @Router       /albums [get]
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums godoc
// @Summary      Add an album
// @Description  Takes a JSON body and adds it to the albums list.
// @Accept       json
// @Produce      json
// @Param        album  body      album  true  "Add album"
// @Success      201    {object}  album
// @Router       /albums [post]
func postAlbums(c *gin.Context) {
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID godoc
// @Summary      Get album by ID
// @Description  Returns a single album by its ID.
// @Produce      json
// @Param        id   path      string  true  "Album ID"
// @Success      200  {object}  album
// @Failure      404  {object}  map[string]string
// @Router       /albums/{id} [get]
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}