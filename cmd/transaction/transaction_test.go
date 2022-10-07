package transaction

import (
	"testing"
)

func TestHowMuchWasSpentMonthly(t *testing.T) {
	testTable := []struct {
		name         string
		transactions Transactions
		wantErr      bool
	}{
		{
			name: "OK",
			transactions: Transactions{
				Transaction: []transaction{
					{
						Time:           "22-01-01 00:00",
						GasPrice:       84.24001654397276,
						GasValue:       364.7976616468897,
						Average:        0.006750761716697318,
						MaxGasPrice:    15555.000008704,
						MedianGasPrice: 79.535685632,
					},
					{
						Time:           "22-01-01 01:00",
						GasPrice:       80.49914519853286,
						GasValue:       364.7976616468897,
						Average:        0.006750761716697318,
						MaxGasPrice:    15555.000008704,
						MedianGasPrice: 79.535685632,
					},
				},
			},
			wantErr: false,
		},
		{
			name:         "Error transactions is empty",
			transactions: Transactions{},
			wantErr:      true,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			_, err := testCase.transactions.HowMuchWasSpentMonthly()
			if err != nil && !testCase.wantErr {
				t.Fatalf("Unexpected error: %s", err.Error())
			}

			if err == nil && testCase.wantErr {
				t.Fatal("Error was expected, but got nil")
			}
		})
	}
}

func TestAveragePricePerDay(t *testing.T) {
	testTable := []struct {
		name         string
		transactions Transactions
		date         string
		wantErr      bool
	}{
		{
			name: "OK",
			transactions: Transactions{
				Transaction: []transaction{
					{
						Time:           "22-01-01 00:00",
						GasPrice:       84.24001654397276,
						GasValue:       364.7976616468897,
						Average:        0.006750761716697318,
						MaxGasPrice:    15555.000008704,
						MedianGasPrice: 79.535685632,
					},
				},
			},
			date:    "22-01-01",
			wantErr: false,
		},
		{
			name:         "Error transactions is empty",
			transactions: Transactions{},
			date:         "22-01-01",
			wantErr:      true,
		},
		{
			name: "Error bad date",
			transactions: Transactions{
				Transaction: []transaction{
					{
						Time:           "22-01-01 00:00",
						GasPrice:       84.24001654397276,
						GasValue:       364.7976616468897,
						Average:        0.006750761716697318,
						MaxGasPrice:    15555.000008704,
						MedianGasPrice: 79.535685632,
					},
				},
			},
			date:    "22-09-01",
			wantErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			_, err := testCase.transactions.AveragePricePerDay(testCase.date)
			if err != nil && !testCase.wantErr {
				t.Fatalf("Unexpected error: %s", err.Error())
			}

			if err == nil && testCase.wantErr {
				t.Fatal("Error was expected, but got nil")
			}
		})
	}
}

func TestFrequencyDistributionOfPricesByHours(t *testing.T) {
	testTable := []struct {
		name         string
		transactions Transactions
		wantErr      bool
	}{
		{
			name: "OK",
			transactions: Transactions{
				Transaction: []transaction{
					{
						Time:           "22-01-01 00:00",
						GasPrice:       84.24001654397276,
						GasValue:       364.7976616468897,
						Average:        0.006750761716697318,
						MaxGasPrice:    15555.000008704,
						MedianGasPrice: 79.535685632,
					},
				},
			},
			wantErr: false,
		},
		{
			name:         "Error transactions is empty",
			transactions: Transactions{},
			wantErr:      true,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			_, err := testCase.transactions.FrequencyDistributionOfPricesByHours()
			if err != nil && !testCase.wantErr {
				t.Fatalf("Unexpected error: %s", err.Error())
			}

			if err == nil && testCase.wantErr {
				t.Fatal("Error was expected, but got nil")
			}
		})
	}
}

func TestHowMuchDidPay(t *testing.T) {
	testTable := []struct {
		name         string
		transactions Transactions
		wantErr      bool
	}{
		{
			name: "OK",
			transactions: Transactions{
				Transaction: []transaction{
					{
						Time:           "22-01-01 00:00",
						GasPrice:       84.24001654397276,
						GasValue:       364.7976616468897,
						Average:        0.006750761716697318,
						MaxGasPrice:    15555.000008704,
						MedianGasPrice: 79.535685632,
					},
				},
			},
			wantErr: false,
		},
		{
			name:         "Error transactions is empty",
			transactions: Transactions{},
			wantErr:      true,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			_, err := testCase.transactions.HowMuchDidPay()
			if err != nil && !testCase.wantErr {
				t.Fatalf("Unexpected error: %s", err.Error())
			}

			if err == nil && testCase.wantErr {
				t.Fatal("Error was expected, but got nil")
			}
		})
	}
}
