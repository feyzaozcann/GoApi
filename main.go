package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	err := OpenDatabase()
	if err != nil {
		log.Printf("error connecting to database %v", err)
	}

	defer CloseDatabase()

	router := gin.Default()
	router.GET("/shows", getShows)
	router.GET("/shows/:id", showById)
	router.POST("/shows", addShow)
	router.DELETE("/shows/:id", removeShow)
	router.Run("localhost:8000")
}
