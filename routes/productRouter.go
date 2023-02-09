package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func productRoutes(routerGroup *gin.RouterGroup) {
	productRouter := routerGroup.Group("/product")
	{
		productRouter.GET("/:id", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "success"})
		})
	}
}
