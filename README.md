# history-processing-service

## Задача: реализовать сервис обработки истории цены gas в сети ethereum.

### Вводные: существует исторический массив данных по цене - 
https://github.com/CryptoRStar/GasPriceTestTask/blob/main/gas_price.json 

### Необходимо посчитать:
1. Сколько было потрачено gas помесячно.
2. Среднюю цену gas за день.
3. Частотное распределение цены по часам(за весь период).
4. Сколько заплатили за весь период (gas price * value).

### Требования к сервису:
1. Данные должны загружаться удаленно.
2. Сервис должен вернуть все значения в виде json файла.
3. Данные должны быть посчитаны максимально быстро.

## Инструкция

- Сборка программы `make build`
- Сборка и запуск программы `make run`
- Запуск тестов `make test`
- Запуск тестов и отображение покрытия кода тестами `make cover`

После запуска приложения будет предложено выбрать действие, потребуется ввести цифру соответствующую действию. Все результаты будут выведены в консоль и сохранены в json файлах в директории `results`. 
### Имена файлов в соответствии действию:
- `1)HowMuchWasSpentMonthly.json` - Сколько было потрачено gas помесячно.
- `2)AveragePricePerDay.json` Среднюю цену gas за день.
- `3)FrequencyDistributionOfPricesByHours.json` Частотное распределение цены по часам(за весь период).
- `4)HowMuchDidPay.json` Сколько заплатили за весь период (gas price * value).

