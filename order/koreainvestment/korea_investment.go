package koreainvestment

import (
	"fmt"
	"log"
	"time"

	krxcode "github.com/sunglim/go-korea-stock-code/code"
	ki "sunglim.github.com/sunglim/systemtrading/pkg/koreainvestment"
)

var productionUrl string

// var accessToken token
var accountInfo ki.KoreaInvestmentAccount

var ki_package *ki.KoreaInvestment

func Initialize(url, applicationKey, applicationSecret string, account ki.KoreaInvestmentAccount) {
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
