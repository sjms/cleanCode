package routes

import "github.com/gin-gonic/gin"

func AddRoutes(routerGroup *gin.RouterGroup) {
	couponRoutes(routerGroup)
	orderRoutes(routerGroup)
	customerRoutes(routerGroup)
	productRoutes(routerGroup)
}
