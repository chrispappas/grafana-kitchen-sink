package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
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
	{ID: "1", Title: "Dookie", Artist: "Green Day", Price: 13.99},
	{ID: "2", Title: "Aenima", Artist: "Tool", Price: 15.99},
	{ID: "3", Title: "Road Apples", Artist: "The Tragically Hip", Price: 2.50},
}

var startTime time.Time

// map of endpoints to request counts
var endpointReqsMap = make(map[string]int)

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getHome(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"status": "Running, GET /albums to get started"})
}

func getStats(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"status":       "running",
		"uptime":       fmt.Sprintf("%v", time.Now().Sub(startTime)),
		"endpointHits": endpointReqsMap,
	})
}

func statsCollector() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := fmt.Sprintf("%s %s", c.Request.Method, c.FullPath())
		endpointReqsMap[key] += 1

		c.Next()
	}
}

func main() {
	startTime = time.Now()

	router := gin.Default()

	// add middleware to collect stats on all routes
	router.Use(statsCollector())

	router.GET("/", getHome)
	router.GET("/stats", getStats)
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run("0.0.0.0:8080")
}
