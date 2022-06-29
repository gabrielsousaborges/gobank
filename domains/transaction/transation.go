package transation

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConnectDataBase() (*sql.DB, error) {

	err := godotenv.Load()
	if err != nil {
		log.Printf("Erro ao conectar %v", err)
		log.Fatal(err)

	}

	postgresConnection := fmt.Sprintf("host=localhost port=15432 user=postgres password=12345 dbname=postgresql sslmode=disable")

	db, err := sql.Open("postgres", postgresConnection)
	if err != nil {
		log.Printf("Erro ao abrir %v", err)

		return nil, err
	}

	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Printf("Erro ao conectar no postgress")

		return nil, err
	}

	return db, nil

}

func CreateTableUser(db *sql.DB) error {
	sqlStatement := fmt.Sprintf("CREATE TABLE IF NOT EXISTS Users (ID int Primary Key NOT NULL, Name VARCHAR(50) NOT NULL, CPFCNPJ VARCHAR UNIQUE  NOT NULL, Email VARCHAR(50) NOT NULL, Password VARCHAR(15) NOT NULL, Balance FLOAT , Type TEXT CHECK (Type in('lojista','comum')));")
	_, err := db.Exec(sqlStatement)
	if err != nil {
		return fmt.Errorf("falha ao realizar")
	}
	return nil
}

func CreateTableTransaction(db *sql.DB) error {
	create, err := db.Prepare("CREATE TABLE IF NOT EXISTS Transaction (ID int PRIMARY KEY NOT NULL , Value FLOAT, IDOrigin INT, IDDestiny INT, DateTime TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL);")
	if err != nil {
		return err
	}

	_, err = create.Exec()
	if err != nil {
		return err
	}
	return nil
}
