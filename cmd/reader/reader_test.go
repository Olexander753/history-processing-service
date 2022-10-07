package reader

import "testing"

func TestReadTransactions(t *testing.T) {
	testTable := []struct {
		name    string
		url     string
		wantErr bool
	}{
		{
			name:    "OK",
			url:     "https://raw.githubusercontent.com/CryptoRStar/GasPriceTestTask/main/gas_price.json",
			wantErr: false,
		},
		{
			name:    "Error bad url",
			url:     "https://raw.githubusercontent.com/CryptoRStar/GasPriceTestTas",
			wantErr: true,
		},
		{
			name:    "Error empty url",
			url:     "",
			wantErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			_, err := ReadTransactions(testCase.url)
			if err != nil && !testCase.wantErr {
				t.Fatalf("Unexpected error: %s", err.Error())
			}

			if err == nil && testCase.wantErr {
				t.Fatal("Error was expected, but got nil")
			}
		})
	}
}
