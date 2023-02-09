package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func orderRoutes(routerGroup *gin.RouterGroup) {
	orderRouter := routerGroup.Group("/order")
	{
		orderRouter.GET("/:id", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "success"})
		})
	}
}
