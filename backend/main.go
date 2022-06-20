package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/vitkarpov/ValushaJS/backend/db"
	"github.com/vitkarpov/ValushaJS/backend/handlers"
	"github.com/vitkarpov/ValushaJS/backend/models"
)

func main() {
	router := gin.Default()
	config := models.NewConfig()
	dbInstance, err := db.NewDatabase(config)
	router.Use(db.SetContextMiddlware(dbInstance))
	if err != nil {
		log.Fatalln(fmt.Printf("Cannot connect to database: %s", err.Error()))
	}
	defer dbInstance.Conn.Close()
	handlers.GetCommentsHandler(router)
	router.Run(config.BindAddr)
}
