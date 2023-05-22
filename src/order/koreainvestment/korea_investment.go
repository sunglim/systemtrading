package koreainvestment

import (
	"fmt"
	"time"

	gocron "github.com/go-co-op/gocron"
	"sunglim.github.com/sunglim/koreaexchange"
)

type KoreaInvestmentAccount struct {
	// 종합계좌번호; 계좌번호 체계(8-2)의 앞 8자리
	CANO string

	// 계좌상품코드; 계좌번호 체계(8-2)의 뒤 2자리
	ACNT_PRDT_CD string
}

var productionUrl string
var appKey string
var appSecret string

var accessToken string
var accountInfo KoreaInvestmentAccount

func Initialize(url, applicationKey, applicationSecret string, account KoreaInvestmentAccount) {
	productionUrl = url
	appKey = applicationKey
	appSecret = applicationSecret
	accountInfo = account

	fmt.Printf("Initialize Korea investment trading.\n ProductionUrl[%s], AppKey[%s], AppSecret[%s], AccountInfo[%v]\n",
		productionUrl, appKey, appSecret, accountInfo)

	refreshToken()
}

func setAccessToken() {
	accessToken = ApiGetAccessToken{}.Call().AccessToken
	fmt.Printf("\nset token %s: %s\n", time.Now().String(), accessToken)
}

func refreshToken() {
	setAccessToken()

	s := gocron.NewScheduler(time.UTC).Every(10).Hour()
	s.Do(setAccessToken)
	s.StartAsync()
}

func DemoCallFunction() {
	refreshToken()

	/*
		price := ApiInqueryPrice{}
		sam := price.Call(koreaexchange.Code삼성전자)
	*/
	api := ApiOrderCash{
		stockCode: koreaexchange.Code삼성전자,
	}
	response := api.Call()
	sam := response.RtCd

	fmt.Printf("price: %s \n", sam)

	time.Sleep(time.Hour)
}

func getAlwaysValidAccessToken() string {
	return "Bearer " + accessToken
}
