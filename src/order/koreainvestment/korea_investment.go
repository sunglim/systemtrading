package koreainvestment

import (
	"fmt"
	"time"

	gocron "github.com/go-co-op/gocron"
	"sunglim.github.com/sunglim/koreaexchange"
)

var productionUrl string
var appKey string
var appSecret string
var accessToken string

func Initialize(url, applicationKey, applicationSecret string) {
	productionUrl = url
	appKey = applicationKey
	appSecret = applicationSecret
}

func setAccessToken() {
	accessToken = ApiGetAccessToken{}.Call().AccessToken
	fmt.Printf("set token %s: %s\n", time.Now().String(), accessToken)
}

func refreshToken() {
	setAccessToken()

	s := gocron.NewScheduler(time.UTC).Every(1).Hour()
	s.Do(setAccessToken)
	s.StartAsync()
}

func DemoCallFunction() {
	fmt.Printf("called %s\n", time.Now().String())

	refreshToken()

	price := ApiInqueryPrice{}
	token := price.Call(koreaexchange.Code삼성전자)

	fmt.Print("price: " + token)

	time.Sleep(time.Hour)
}

func getAlwaysValidAccessToken() string {
	return "Bearer " + accessToken
}
