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

var db *sql.DB

func init() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
}

// Returna um ResultSet
func AbrirConsulta(consulta string, params ...interface{}) (*sql.Rows, error) {

	stmt, err := db.Prepare(consulta)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(params...)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func ExecutarComandos(comando string, params ...interface{}) (int64, error) {

	stmt, err := db.Prepare(comando)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(params...)
	if err != nil {
		return 0, err
	}

	rowsAfetadas, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAfetadas, nil
}

func ExistemRegistros(rows *sql.Rows) bool {
	return rows.Next()
}
