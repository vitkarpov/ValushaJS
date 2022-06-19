package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vitkarpov/ValushaJS/backend/db"
)

func getComments(c *gin.Context) {
	dbInstance := c.Keys["DB"].(db.Database)
	response := dbInstance.GetComments()
	status := http.StatusOK
	if response.Error != "" {
		status = http.StatusInternalServerError
	}
	c.IndentedJSON(status, response)
}

func GetCommentsHandler(router *gin.Engine) *gin.Engine {
	router.GET("/comments", getComments)
	return router
}
