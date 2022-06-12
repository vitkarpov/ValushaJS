package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vitkarpov/ValushaJS/backend/db"
)

var dbInstance db.Database

func getComments(c *gin.Context) {
	response := dbInstance.GetComments()
	status := http.StatusOK
	if response.Error != "" {
		status = http.StatusInternalServerError
	}
	c.IndentedJSON(status, response)
}

func GetCommentsHandler(router *gin.Engine, db db.Database) *gin.Engine {
	dbInstance = db
	router.GET("/comments", getComments)
	return router
}
