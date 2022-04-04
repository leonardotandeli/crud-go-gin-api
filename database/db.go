package database

import (
	"gogin/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComOBancoDeDados() {
	stringDeConexao := "host=localhost user=postgres password=1234 dbname=Alura_Loja sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("erro ao conectar no db")
	}
	DB.AutoMigrate(&models.Aluno{})
}
