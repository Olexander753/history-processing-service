package reader

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Olexander753/history-processing-service/cmd/transaction"
)

// функция чтения списка транзакций
func ReadTransactions(url string) (transaction.Transactions, error) {

	// создание пустого объкта типа Transactions
	transactions := transaction.Transactions{}

	// отправка get запроса протокола http
	resp, err := http.Get(url)
	if err != nil {
		return transactions, err
	}

	// декодирование тела ответа запроса
	decoder := json.NewDecoder(resp.Body)
	for {
		// получаем токен
		tok, err := decoder.Token()
		// проверка на ошибку
		if err != nil && err != io.EOF {
			return transactions, err
		}
		// проверка токена
		if tok == nil {
			break
		}
		// декодирование
		decoder.Decode(&transactions)
	}

	return transactions, nil
}
