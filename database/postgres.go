package database

import (
	"fmt"
	"log"

	"github.com/Gierdiaz/Echo-golang/config"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func Connect(config *config.Config) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName,
	)

	var err error
	DB, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	log.Println("Conexão com o banco de dados estabelicido com sucesso!")
}