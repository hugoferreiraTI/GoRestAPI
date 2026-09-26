package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" //baixando a conexao
)

const (
	host     = "go_db"
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbnamme  = "postgres"
)

func ConnectDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+ //usamos a const para passar os dados do database via placeholder
		"password= %s dbname=%s sslmode=disable",
		host, port, user, password, dbnamme) //constantes

	db, err := sql.Open("postgres", psqlInfo) //abre a conexao

	if err != nil { //verificamos se está tendo conexao via ping
		panic(err)
	}

	err = db.Ping()
	if err != nil { //verificamos novamente se teve um erro
		panic(err)
	}

	fmt.Println("Connected to " + dbnamme)
	return db, nil
}
