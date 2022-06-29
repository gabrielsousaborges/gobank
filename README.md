### Passo 1:
Suba os containers:
```sh
docker-compose up
```

### Passo 2:
Após a inicialização do docker-compose rode o seguinte comando:
```sh
go run router.go
```

- *Após efetuar esses passos, sua aplicação deverá estar no ar!*

# Métodos
Requisições para a API devem seguir os padrões:
| Método | Descrição |
|---|---|
| `POST` | Insere uma nova transação no Banco de Dados|

## Notas

Endpoint para testar as transações - http://localhost:8080/transaction

Payload: 
```
{
    "value": ,
    "payee": ,
    "payer": 
}
```