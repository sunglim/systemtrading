package koreainvestment

import (
	"fmt"
	"log"
	"time"

	krxcode "github.com/sunglim/go-korea-stock-code/code"
	ki "sunglim.github.com/sunglim/systemtrading/pkg/koreainvestment"
)

type KoreaInvestmentAccount struct {
	// 종합계좌번호; 계좌번호 체계(8-2)의 앞 8자리
	CANO string

	// 계좌상품코드; 계좌번호 체계(8-2)의 뒤 2자리
	ACNT_PRDT_CD string
}

var productionUrl string

// var accessToken token
var accountInfo KoreaInvestmentAccount

var ki_package *ki.KoreaInvestment

func Initialize(url, applicationKey, applicationSecret string, account KoreaInvestmentAccount) {
	productionUrl = url
	accountInfo = account

	ki_package = ki.NewKoreaInvestment(ki.Credential{
		AppKey:    applicationKey,
		AppSecret: applicationSecret,
	}, log.Default())
	ki_package.InitializeToken()

	fmt.Printf("Initialize Korea investment trading.\n ProductionUrl[%s], AppKey[%s], AppSecret[%s], AccountInfo[%v]\n",
		productionUrl, ki_package.GetCredential().AppKey, ki_package.GetCredential().AppSecret, accountInfo)
}

func DemoCallFunction() {
	/*
		price := ApiInqueryPrice{}
		sam := price.Call(koreaexchange.Code삼성전자)
	*/
	api := ApiOrderCash{
		stockCode: krxcode.Code삼성전자,
	}
	response := api.Call()
	sam := response.RtCd

	fmt.Printf("price: %s \n", sam)

	time.Sleep(time.Hour)
}
