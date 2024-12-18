package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type show struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Year   int     `json:"year"`
	Type   string  `json:"type"`
	Rating float64 `json:"rating"`
}

var shows = []show{
	{ID: "1", Title: "Manchester by the Sea", Year: 2016, Type: "movie", Rating: 7.8},
	{ID: "2", Title: "How I Met Your Mother", Year: 2005, Type: "series", Rating: 8.3},
	{ID: "3", Title: "Titanic", Year: 1997, Type: "movie", Rating: 7.9},
	{ID: "4", Title: "Inception", Year: 2010, Type: "movie", Rating: 8.8},
	{ID: "5", Title: "The Crown", Year: 2016, Type: "series", Rating: 8.7},
}

func getShows(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, shows)
}

func showById(c *gin.Context) {
	id := c.Param("id")
	show, err := getShowById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Show not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, show)
}

func getShowById(id string) (*show, error) {
	for i, s := range shows {
		if s.ID == id {
			return &shows[i], nil
		}
	}
	return nil, errors.New("show not found")
}

func addShow(c *gin.Context) {
	var newShow show
	if err := c.BindJSON(&newShow); err != nil {
		return
	}
	shows = append(shows, newShow)
	c.IndentedJSON(http.StatusCreated, newShow)
}

func removeShow(c *gin.Context) {
	id := c.Param("id")

	for i, s := range shows {
		if s.ID == id {
			shows = append(shows[:i], shows[i+1:]...) //slicedan siler
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Show deleted"})
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Show not found"})
}

func main() {
	router := gin.Default()
	router.GET("/shows", getShows)
	router.GET("/shows/:id", showById)
	router.POST("/shows", addShow)
	router.DELETE("/shows/:id", removeShow)
	router.Run("localhost:8000")
}
