package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	r := gin.Default()

	r.StaticFile("/webboard.html", "./webboard.html")

	r.GET("/api/boards", func(c *gin.Context) {
		c.JSON(http.StatusOK, questionnaire)
	})
	r.POST("/api/boards", func(c *gin.Context) {
		var question board
		if err := c.ShouldBindJSON(&question); err != nil {
			questionnaire = append(questionnaire, board{
				Message: err.Error(),
			})
			c.JSON(http.StatusBadRequest, questionnaire)
			return
		}
		questionnaire = append(questionnaire, question)
		c.JSON(http.StatusOK, questionnaire)
	})

	r.Run()
}

type table struct {
	Name string `json:"name"`
}

type profile struct {
	Name    string
	Email   string
	Message string
}

type board struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

var questionnaire = []board{}
