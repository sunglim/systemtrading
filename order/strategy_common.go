package order

import (
	"os"

	"sunglim.github.com/sunglim/systemtrading/order/koreainvestment"
	ki "sunglim.github.com/sunglim/systemtrading/pkg/koreainvestment"
)

// Initialize Korea inestment settings
// Initialize access token, refresh every 1 hour.
func initializeKoreaInvestment() error {
	url := os.Args[1]
	appKey := os.Args[2]
	appSecret := os.Args[3]
	account := os.Args[4]

	koreainvestment.Initialize(url, appKey, appSecret, ki.KoreaInvestmentAccount{
		CANO:         account,
		ACNT_PRDT_CD: "01",
	})

	return nil
}
