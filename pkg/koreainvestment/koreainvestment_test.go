package koreainvestment

import (
	"testing"
)

func TestToHanaVersions(t *testing.T) {
	forTesting = true

	api := NewKoreaInvestmentTokenRefresh(Credential{
		AppKey:    "..",
		AppSecret: "..",
	}, 1)
	api.InitializeToken()
}
