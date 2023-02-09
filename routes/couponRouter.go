package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func couponRoutes(routerGroup *gin.RouterGroup) {
	couponRouter := routerGroup.Group("/coupon")
	{
		couponRouter.GET("/:id", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "success"})
		})
	}
}
