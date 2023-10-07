package koreainvestment

import (
	"testing"
)

func TestKoreaInvestmentAccount(t *testing.T) {
	const validFormat = "123456-78"
	result, err := ConvertToKoreaInvestmentAccount(validFormat)
	if err != nil {
		t.Errorf(`unwanted error`)
	}
	if result.CANO != "123456" {
		t.Errorf(`cannot parse CANO`)
	}
	if result.ACNT_PRDT_CD != "78" {
		t.Errorf(`cannot parse ACNT_PRDT_CD`)
	}
}

func TestKoreaInvestmentAccountInvalidFormat(t *testing.T) {
	const validFormat = "12345678"
	_, err := ConvertToKoreaInvestmentAccount(validFormat)
	if err != nil {
		t.Errorf(`unwanted error`)
	}
}
