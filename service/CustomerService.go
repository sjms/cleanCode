package service

import (
	"cleanCode/db"
	"cleanCode/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func PostCustomer(c *gin.Context) {
	// Validate input
	var input models.Customers
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Print(input)
	customer := models.Customers{Name: input.Name}

	if err := db.Connection.Create(&customer).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "error creating a customer"})
		return
	}
	c.JSON(http.StatusOK, customer)
}
