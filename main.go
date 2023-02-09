package main

import (
	"cleanCode/db"
	"cleanCode/routes"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func setup() *gin.Engine {
	app := gin.Default()
	router := app.Group("/api")
	{
		router.GET("/health-check", func(c *gin.Context) {
			c.JSON(http.StatusOK, "working")
		})
	}
	routes.AddRoutes(router)
	return app
}

func main() {
	db.Init()
	r := setup()
	err := r.Run()
	if err != nil {
		log.Fatal(err)
	}

}
