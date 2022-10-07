package transaction

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
)

// структура для хранения списка транзакций
type Transactions struct {
	Transaction []transaction `json:"transactions"`
}

type transaction struct {
	Time           string  `json:"time"`
	GasPrice       float64 `json:"gasPrice"`
	GasValue       float64 `json:"gasValue"`
	Average        float64 `json:"average"`
	MaxGasPrice    float64 `json:"maxGasPrice"`
	MedianGasPrice float64 `json:"medianGasPrice"`
}

// 1. Сколько было потрачено gas помесячно;
func (t *Transactions) HowMuchWasSpentMonthly() ([]byte, error) {
	// проверка на поустоту
	if len(t.Transaction) > 0 {
		// создание карты для удобства сбора информации
		m := make(map[string]float64)
		// цикл по списку транзакций
		for _, val := range t.Transaction {
			// если в карте уже есть записть по данному месяцу то к значению прибаляется GasPrice
			if value, ok := m[string(val.Time[:5])]; ok {
				m[string(val.Time[:5])] = value + val.GasPrice
			} else {
				// если такой записи нет то записываем
				m[string(val.Time[:5])] = val.GasPrice
			}
		}

		// перевод данных в json формат
		data, err := json.Marshal(m)
		if err != nil {
			return nil, err
		}

		return data, nil
	} else {
		return nil, fmt.Errorf("Transactions is empty!")
	}
}

// 2. Среднюю цену gas за день;
func (t *Transactions) AveragePricePerDay(date string) ([]byte, error) {
	if len(t.Transaction) > 0 {
		// инициализация переменных для вычисления средней цены
		var averagePrice, i float64
		averagePrice = 0
		i = 0
		// цикл по элементам списка транзакций
		for _, val := range t.Transaction {
			// ищет по дате все соответсвия
			if strings.Contains(val.Time, date) {
				// считает сумму за день и количество транзакций по данной дате
				averagePrice += val.GasPrice
				i++
			}
		}
		if i == 0 {
			return nil, fmt.Errorf("Bad date!")
		}

		// перевод данных в json формат
		data, err := json.Marshal(map[string]float64{
			date: averagePrice / i,
		})
		if err != nil {
			return nil, err
		}

		return data, nil
	} else {
		return nil, fmt.Errorf("Transactions is empty!")
	}
}

// 3. Частотное распределение цены по часам;
func (t *Transactions) FrequencyDistributionOfPricesByHours() ([]byte, error) {
	if len(t.Transaction) > 0 {
		// реализация время - цены
		// // создание массива структур для хранения частотного распределения цены по часам за весь период
		// frequencyDistribution := make(map[string][]float64)
		// // цикл по элементам списка транзакций
		// for _, val := range t.Transaction {
		// 	// в массив структур для частотного распределения записывается по элементно дата и цена
		// 	frequencyDistribution[val.Time[9:]] = append(frequencyDistribution[val.Time[9:]], val.GasPrice)
		// }
		// // перевод данных в json формат
		// data, err := json.Marshal(frequencyDistribution)
		// if err != nil {
		// 	return nil, err
		// }

		// реализация время -> даты и цены
		// создание массива структур для хранения частотного распределения цены по часам за весь период
		frequencyDistribution := make(map[string][]struct {
			Date  string  `json:"date"`
			Price float64 `json:"price"`
		})
		// цикл по элементам списка транзакций
		for _, val := range t.Transaction {
			// в массив структур для частотного распределения записывается по элементно дата и цена
			frequencyDistribution[val.Time[9:]] = append(frequencyDistribution[val.Time[9:]], struct {
				Date  string  `json:"date"`
				Price float64 `json:"price"`
			}{Date: val.Time[:8], Price: val.GasPrice})
		}
		// перевод данных в json формат
		data, err := json.Marshal(frequencyDistribution)
		if err != nil {
			return nil, err
		}

		return data, nil
	} else {
		return nil, fmt.Errorf("Transactions is empty!")
	}
}

// 4. Сколько заплатили за весь период;
func (t *Transactions) HowMuchDidPay() ([]byte, error) {
	if len(t.Transaction) > 0 {
		// инициализация переменной для подсчёта суммы
		sum := big.NewFloat(0)
		// цикл по всем элементам списка транзакций
		for _, val := range t.Transaction {
			// посчёт суммы
			sum.Add(sum, big.NewFloat(val.GasPrice*val.GasValue))
		}

		// перевод данных в json формат
		data, err := json.Marshal(map[string]*big.Float{
			"За весь период было потрачено ": sum,
		})
		if err != nil {
			return nil, err
		}

		return data, nil
	} else {
		return nil, fmt.Errorf("Transactions is empty!")
	}
}
