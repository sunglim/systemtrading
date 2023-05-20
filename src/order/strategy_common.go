package order

import (
	"errors"
	"os"

	"sunglim.github.com/sunglim/order/koreainvestment"
)

// Initialize Korea inestment settings
// Initialize access token, refresh every 1 hour.
func initializeKoreaInvestment() error {
	if len(os.Args) != 5 {
		return errors.New("insufficient arguments")
	}

	url := os.Args[1]
	//url := "https://openapivts.koreainvestment.com:29443"
	appKey := os.Args[2]
	appSecret := os.Args[3]
	account := os.Args[4]

	koreainvestment.Initialize(url, appKey, appSecret, koreainvestment.KoreaInvestmentAccount{
		CANO:         account,
		ACNT_PRDT_CD: "01",
	})

	return nil
}
