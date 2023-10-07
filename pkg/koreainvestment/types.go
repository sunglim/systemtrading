package koreainvestment

import (
	"fmt"
	"strings"
)

type KoreaInvestmentAccount struct {
	// 종합계좌번호; 계좌번호 체계(8-2)의 앞 8자리
	CANO string

	// 계좌상품코드; 계좌번호 체계(8-2)의 뒤 2자리
	ACNT_PRDT_CD string
}

func ConvertToKoreaInvestmentAccountNoError(rawString string) KoreaInvestmentAccount {
	account, _ := ConvertToKoreaInvestmentAccount(rawString)
	return account
}

func ConvertToKoreaInvestmentAccount(rawString string) (KoreaInvestmentAccount, error) {
	splitted := strings.Split(rawString, "-")
	if len(splitted) != 2 {
		return KoreaInvestmentAccount{}, fmt.Errorf("invalid string format %s", rawString)
	}

	return KoreaInvestmentAccount{
		CANO:         splitted[0],
		ACNT_PRDT_CD: splitted[1],
	}, nil
}
