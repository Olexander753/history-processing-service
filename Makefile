clean:
	rm -f main
	rm -f results/1.HowMuchWasSpentMonthly.json
	rm -f results/2.AveragePricePerDay.json
	rm -f results/3.FrequencyDistributionOfPricesByHours.json
	rm -f results/4.HowMuchDidPay.json

build: clean
	go build cmd/main/main.go

run: build
	./main

test:
	go test -v -count=1 ./...

.PHONY: cover
cover:
	go test -short -count=1 -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
	rm coverage.out