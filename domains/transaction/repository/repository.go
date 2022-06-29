package repository

import (
	"database/sql"
	"fmt"
	"gobank/domains/transaction/model"
	"time"
)

func CreateUser(db *sql.DB, user *model.User) error {

	defer db.Close()

	fmt.Print(user)

	err := db.QueryRow("INSERT INTO Users (Name, type, CPFCNPJ, Email, Password, Balance) values (%v, %v, %v, %v, %v, %v);", user.Name, user.Type, user.CPFCNPJ, user.Email, user.Password, user.Balance)
	if err != nil {
		return fmt.Errorf("falha ao criar usuario")
	}
	return nil

}

func GetUser(db *sql.DB, id int64) (model.User, error) {
	var user model.User

	userDB, err := db.Query("SELECT * FROM Users WHERE id = %v", id)
	if err != nil {
		return model.User{}, fmt.Errorf("falha na execução da busca de user no postgres: %v", err)
	}

	userDB.Next()

	err = userDB.Scan(&user.ID, &user.CPFCNPJ, &user.Email, &user.Password, &user.Balance, &user.Type)
	if err != nil {
		return model.User{}, nil
	}
	return user, nil
}

func MakeTransaction(db *sql.DB, payer model.User, payee model.User, transaction model.Transaction) error {

	dateFormat := fmt.Sprintf("'%v'", time.Now().Format("2006-01-02 15:04:05"))
	transaction.DateTime = dateFormat

	err := db.QueryRow("INSERT INTO Transactions (Value, IDOrigin, IDDestiny, DateTime) values (%v, %v, %v, %v)", transaction.Value, transaction.IDOrigin, transaction.IDDestiny, transaction.DateTime)
	if err != nil {
		return fmt.Errorf("falha ao realizar transacao")
	}

	err = db.QueryRow("UPDATE Users set Balance = (Balance-%v) WHERE ID = %v", transaction.Value, transaction.IDOrigin)
	if err != nil {
		return fmt.Errorf("falha ao realizar Update do saldo")
	}

	err = db.QueryRow("UPDATE Users set Balance = (Balance+%v) WHERE ID = %v", transaction.Value, transaction.IDDestiny)
	if err != nil {
		return fmt.Errorf("falha ao realizar Update do saldo")
	}

	return nil

}
