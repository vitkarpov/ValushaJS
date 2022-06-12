package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type comment struct {
	ID         string `json:"id"`
	AuthorName string `json:"authorName"`
	Comment    string `json:"comment"`
	ParentId   string `json:"parentId"`
}

var comments = []comment{
	{ID: "1", Comment: "Hello, world"},
	{ID: "2", Comment: "Yo, how's it going", ParentId: "1"},
	{ID: "2", Comment: "What's going on here?"},
}

func getComments(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, comments)
}

func main() {
	router := gin.Default()
	router.GET("/comments", getComments)

	router.Run("localhost:8080")
}
