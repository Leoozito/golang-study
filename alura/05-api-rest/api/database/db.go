package database

import (
	"log"

	"github.com/api/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	err error
)

func ConnectDB() {
	connection := "host:localhost user:root password:root dbname:root port:5432"

	DB, err := gorm.Open(postgres.Open(connection))

	if (err != nil) {
		log.Fatal("Erro ao conectar com o banco de dados")
	}

	DB.AutoMigrate(&model.Student{})
}