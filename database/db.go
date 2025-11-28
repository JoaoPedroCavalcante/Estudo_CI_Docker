package database

import (
	"log"
	"os"

	"github.com/guilhermeonrails/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {

	// HOST
	host := os.Getenv("HOST")
	if host == "" {
		host = "postgres" // fallback para Docker Compose
	}

	// PORT
	port := os.Getenv("PORT")
	if port == "" {
		port = "5432"
	}

	// USER
	user := os.Getenv("USER")
	if user == "" {
		user = "root"
	}

	// PASSWORD
	password := os.Getenv("PASSWORD")
	if password == "" {
		password = "root"
	}

	// DBNAME
	dbname := os.Getenv("DBNAME")
	if dbname == "" {
		dbname = "root"
	}

	// Conexão final
	stringDeConexao :=
		"host=" + host +
			" user=" + user +
			" password=" + password +
			" dbname=" + dbname +
			" port=" + port +
			" sslmode=disable"

	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}

	DB.AutoMigrate(&models.Aluno{})
}
