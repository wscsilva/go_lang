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

func AbrirConsulta(consulta string, params ...interface{}) (*sql.Rows, error) {
	connection, err := ConectarBancoDeDados()
	if err != nil {
		return nil, err
	}
	stmt, err := connection.Prepare(consulta)
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

func ExecutarConsulta(comando string, params ...interface{}) (int64, error) {
	connection, err := ConectarBancoDeDados()
	if err != nil {
		return 0, err
	}
	stmt, err := connection.Prepare(comando)
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

func ExistemRegistro(stmt *sql.Rows) bool {
	if stmt.Next() {
		return true
	} else {
		return false
	}
}
