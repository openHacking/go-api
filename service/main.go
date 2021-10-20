package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan", Artist: "Sarah", Price: 38.99},
}

// getAlbums responds with the list of all albums as JSON
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {

	/*

		id := c.Query("id") //Query the parameters spliced after the request URL

		name := c.PostForm("name") //Query parameters from the form

		uuid := c.Param("uuid") //Get the parameters in the URL
	*/
	id := c.Param("id")
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
func main() {
	router := gin.Default()

	/*
		test:

		curl http://localhost:8080/albums \
		--header "Content-Type: application/json" \
		--request "GET"
	*/
	router.GET("/albums", getAlbums)

	/*
		test:

		curl http://localhost:8080/albums/2
	*/
	router.GET("/albums/:id", getAlbumByID)

	/*
		test:

		curl http://localhost:8080/albums \
		--include \
		--header "Content-Type: application/json" \
		--request "POST" \
		--data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'
	*/
	router.POST("/albums", postAlbums)
	router.Run("localhost:8080")
}

////////////////////////////
