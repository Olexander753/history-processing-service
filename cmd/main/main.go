package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Olexander753/history-processing-service/cmd/config"
	"github.com/Olexander753/history-processing-service/cmd/reader"
)

func main() {
	var n int
	// чтение конфига, в нём лежат ссылка на json файл
	cfg := config.GetConfig()

	// читение и декодирование в массив транзакций json файла
	transactions, err := reader.ReadTransactions(cfg.URL)

	// проверка на ошибку
	if err != nil {
		log.Println(err)
	}

	// бесконечный цикл с выбором действий
	for {
		fmt.Println(`
Выбирете действие (введите цифру):
1. Сколько было потрачено gas помесячно;
2. Среднюю цену gas за день;
3. Частотное распределение цены по часам;
4. Сколько заплатили за весь период;
5. Выход.`)
		fmt.Scan(&n)
		switch n {
		case 1:
			// вызов функции соотвествующей выбраному действию
			data, err := transactions.HowMuchWasSpentMonthly()
			if err != nil {
				fmt.Println(err)
			} else {
				// вывод в консоль
				fmt.Println(string(data))

				// вызов функции записи данных в файл
				err = writeToJson("1)HowMuchWasSpentMonthly.json", data)
				if err != nil {
					fmt.Println(err)
				}
			}

		case 2:
			// ввод даты для вывода средней цены  за день
			var date string
			fmt.Println("Введите дату вида 22-01-01.")
			fmt.Scan(&date)

			// вызов функции соотвествующей выбраному действию
			data, err := transactions.AveragePricePerDay(date)
			if err != nil {
				fmt.Println(err)
			} else {
				// вывод в консоль
				fmt.Println(string(data))

				// вызов функции записи данных в файл
				err = writeToJson("2)AveragePricePerDay.json", data)
				if err != nil {
					fmt.Println(err)
				}
			}

		case 3:
			// вызов функции соотвествующей выбраному действию
			data, err := transactions.FrequencyDistributionOfPricesByHours()
			if err != nil {
				fmt.Println(err)
			} else {
				// вывод в консоль
				fmt.Println(string(data))

				// вызов функции записи данных в файл
				err = writeToJson("3)FrequencyDistributionOfPricesByHours.json", data)
				if err != nil {
					fmt.Println(err)
				}
			}

		case 4:
			// вызов функции соотвествующей выбраному действию
			data, err := transactions.HowMuchDidPay()
			if err != nil {
				fmt.Println(err)
			} else {
				// вывод в консоль
				fmt.Println(string(data))

				// вызов функции записи данных в файл
				err = writeToJson("4)HowMuchDidPay.json", data)
				if err != nil {
					fmt.Println(err)
				}
			}

		case 5:
			// выход из программы
			fmt.Println("Good bye!")
			return
		default:
			fmt.Println("Error: неизветное действие!")
		}
	}
}

// функция записи данных в файл принимает имя файла и данные для записи
func writeToJson(fname string, data []byte) error {
	// открывается файл, если такого файла нет то он будет создан
	fp, err := os.Create(fmt.Sprintf("results/%s", fname))
	if err != nil {
		return err
	}

	// закроется по завершению работы функции
	defer fp.Close()

	// запись данных в файл
	_, err = fp.Write(data)
	if err != nil {
		return err
	}

	return nil
}
