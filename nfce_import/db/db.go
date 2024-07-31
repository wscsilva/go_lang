package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "Maker@1"
	dbname   = "quintel"
)

func ConectarBancoDeDados() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Conexão com o banco de dados estabelecida!")
	return db, nil
}

func FecharConexao(db *sql.DB) {
	db.Close()
	fmt.Println("Conexão com o banco de dados finalizada!")
}
