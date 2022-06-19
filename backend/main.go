package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/vitkarpov/ValushaJS/backend/db"
	"github.com/vitkarpov/ValushaJS/backend/handlers"
)

func main() {
	godotenv.Load()
	router := gin.Default()
	dbInstance, err := db.New()
	router.Use(db.SetContextMiddlware(dbInstance))
	if err != nil {
		log.Fatalln(fmt.Printf("Cannot connect to database: %s", err.Error()))
	}
	defer dbInstance.Conn.Close()
	handlers.GetCommentsHandler(router)
	router.Run()
}
