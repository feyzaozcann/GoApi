package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type show struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Year   int     `json:"year"`
	Type   string  `json:"type"`
	Rating float64 `json:"rating"`
}

func getShows(c *gin.Context) {
	var shows []show

	if err := DB.Select(&shows, "SELECT id, title, year, type, rating FROM shows;"); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to fetch shows"})
		return
	}

	c.IndentedJSON(http.StatusOK, shows)

}

func showById(c *gin.Context) {
	id := c.Param("id")
	show, err := getShowById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "show not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, show)
}

func getShowById(id string) (*show, error) {
	var show show
	if err := DB.Get(&show, "SELECT id, title, year, type, rating FROM shows WHERE id = $1", id); err != nil {
		return nil, err
	}
	return &show, nil
}

func addShow(c *gin.Context) {
	var newShow show

	if err := c.ShouldBindJSON(&newShow); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid input"})
		return
	}

	query := "INSERT INTO shows (title, year, type, rating) VALUES ($1, $2, $3, $4) RETURNING id"

	if err := DB.Get(&newShow.ID, query, newShow.Title, newShow.Year, newShow.Type, newShow.Rating); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to insert show"})
		return
	}

}

func removeShow(c *gin.Context) {
	id := c.Param("id")

	if _, err := DB.Exec("DELETE FROM shows WHERE id = $1", id); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "something went wrong"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "show deleted successfully"})
}
