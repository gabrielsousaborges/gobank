package transport

import (
	"encoding/json"
	"gobank/config"
	"gobank/domains/transaction/model"
	"gobank/domains/transaction/service"
	"io/ioutil"
	"net/http"
)

func NewTransactionHadler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		config.RespondWithError(w, http.StatusMethodNotAllowed, 0, "Metodo incorreto")
		return
	}

	var req = &model.TransactionRequest{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {

		config.RespondWithError(w, http.StatusBadRequest, 0, "Erro ao receber o body de Transaction")
		return
	}

	err = json.Unmarshal(body, req)
	if err != nil {
		config.RespondWithError(w, http.StatusBadRequest, 0, "Erro ao realizar o unmarshal")
		return
	}

	if err := service.MakeTransactionRequest(req); err != nil {

		config.RespondWithError(w, http.StatusBadRequest, 0, "Erro ao realizar a transação")
		return
	}

	config.RespondWithJSON(w, http.StatusOK, "OK")
	return
}

func NewUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		config.RespondWithError(w, http.StatusMethodNotAllowed, 0, "Metodo incorreto")
		return
	}

	var req = &model.User{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {

		config.RespondWithError(w, http.StatusBadRequest, 0, "Erro ao receber o body de Transaction")
		return
	}

	err = json.Unmarshal(body, req)
	if err != nil {
		config.RespondWithError(w, http.StatusBadRequest, 0, "Erro ao realizar o unmarshal")
	}

	if err := service.CreateUserRequest(req); err != nil {
		config.RespondWithError(w, http.StatusBadRequest, 0, "Erro ao criar user ")
		return
	}

	config.RespondWithJSON(w, http.StatusOK, "OK")
	return

}
