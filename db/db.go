package db

import (
	"cleanCode/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var Connection *gorm.DB

func Init() {
	dsn := "host=localhost user=admin password=1234 dbname=store port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("erro abrindo conexão com o banco de dados: ", err)
	}

	err = db.AutoMigrate(&models.Customers{}, &models.Products{}, &models.Orders{})
	if err != nil {
		log.Fatalln("erro configurando o banco de dados: ", err)
	}

	Connection = db
}
