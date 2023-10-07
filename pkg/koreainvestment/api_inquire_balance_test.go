package koreainvestment

import (
	"testing"
)

func TestApiInquireBalance(t *testing.T) {
	api := ApiInquireBalance{
		KoreaInvestmentAccount: KoreaInvestmentAccount{
			CANO:         "123456",
			ACNT_PRDT_CD: "01",
		},
		Credential:  Credential{},
		accessToken: "",
	}
	api.url()
}
