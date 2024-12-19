package main

import (
	"database/sql"
	"errors"
	"log"
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

	rows, err := DB.Query("SELECT id, title, year, type, rating FROM shows;")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for rows.Next() {
		var show show
		if err := rows.Scan(&show.ID, &show.Title, &show.Year, &show.Type, &show.Rating); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "failed to load shows"})
			return
		}
		shows = append(shows, show)
	}

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
	var show show
	err := DB.QueryRow("SELECT id, title, year, type, rating FROM shows WHERE id=$1", id).Scan(&show.ID, &show.Title, &show.Year, &show.Type, &show.Rating)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("show not found")
		}
		return nil, err
	}
	return &show, nil
}

func addShow(c *gin.Context) {
	var newShow show

	if err := c.BindJSON(&newShow); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := "INSERT INTO shows (title, year, type, rating) VALUES ($1, $2, $3, $4) RETURNING id"

	err := DB.QueryRow(query, newShow.Title, newShow.Year, newShow.Type, newShow.Rating).Scan(&newShow.ID)
	if err != nil {
		log.Printf("Error inserting show: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to insert show"})
		return
	}

	c.IndentedJSON(http.StatusCreated, newShow)
}

func removeShow(c *gin.Context) {
	id := c.Param("id")

	result, err := DB.Exec("DELETE FROM shows WHERE id = $1", id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if ra, _ := result.RowsAffected(); ra == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Show not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Show deleted"})
}
