package routes

import (
	"cleanCode/db"
	"cleanCode/models"
	"cleanCode/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func customerRoutes(routerGroup *gin.RouterGroup) {
	customerRouter := routerGroup.Group("/customer")
	{
		customerRouter.GET("/validate-cpf/:cpf", func(c *gin.Context) {
			cpf := c.Params.ByName("cpf")
			isCPF, err := service.IsCPF(cpf)
			if err != nil {
				c.JSON(http.StatusAccepted, gin.H{"isCPF": isCPF, "msg:": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"isCPF": isCPF, "msg:": ""})
		})

		customerRouter.GET("/:id", func(c *gin.Context) {
			var customer models.Customers
			id := c.Param("id")
			if err := db.Connection.First(&customer, "id = ?", id).Error; err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "customer not found"})
				return
			}
			c.JSON(http.StatusOK, customer)
		})

		customerRouter.POST("", func(c *gin.Context) {
			print("asdfasdfasdfsdf")
			service.PostCustomer(c)
		})
	}
}
