package service

import (
	"encoding/json"
	"fmt"
	transation "gobank/domains/transaction"
	"gobank/domains/transaction/model"
	"gobank/domains/transaction/repository"
	"log"
	"net/http"
)

func MakeTransactionRequest(transaction *model.TransactionRequest) error {

	db, err := transation.ConnectDataBase()
	if err != nil {
		log.Printf("Erro ao conectar no banco de dados")
		return err
	}

	mock, err := GetMock()
	if err != nil {
		log.Printf("Erro ao realizar o GetMock")
		return err
	}

	if !mock {
		log.Printf("Erro no mock")
		return err
	}

	payer, err := repository.GetUser(db, transaction.Payer)
	if err != nil {
		log.Printf("Erro ao realizar o GetUserPayer")

		return err
	}

	payee, err := repository.GetUser(db, transaction.Payer)
	if err != nil {
		log.Printf("Erro ao realizar o GetUserPayee")

		return err
	}

	if payer.Type != "comum" {
		log.Printf("Erro ao realizar o GetUserPayer")

		return fmt.Errorf("Pagador do tipo logista")
	}

	if payer.Balance < transaction.Value {
		return fmt.Errorf("Erro ao realizar pagamento, saldo insuficiente")
	}

	transc := model.Transaction{IDOrigin: payer.ID, IDDestiny: payee.ID, Value: transaction.Value}

	if err := repository.MakeTransaction(db, payer, payee, transc); err != nil {
		return err
	}
	return nil
}

func CreateUserRequest(user *model.User) error {

	db, err := transation.ConnectDataBase()
	if err != nil {

		return err
	}

	if err := repository.CreateUser(db, user); err != nil {
		return err
	}

	return nil

}

func GetMock() (bool, error) {

	var (
		mockResponse = model.JSONMock{}
		endpointAPI  = "https://run.mocky.io/v3/d02168c6-d88d-4ff2-aac6-9e9eb3425e31"
	)

	req, err := http.NewRequest(http.MethodGet, endpointAPI, nil)
	if err != nil {
		log.Printf("Erro ao realizar o getMock")

		return false, err
	}
	req.Header.Add("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&mockResponse)
	if err != nil {
		return false, err
	}
	return mockResponse.Authorization, nil

}
