package model

type User struct {
	ID       int64
	Name     string
	CPFCNPJ  string
	Email    string
	Password string
	Balance  float64
	Type     string
}

type Transaction struct {
	ID        int64
	Value     float64
	IDOrigin  int64
	IDDestiny int64
	DateTime  string
}

type TransactionRequest struct {
	Value float64
	Payer int64
	Payee int64
}

type JSONMock struct {
	Authorization bool
}
