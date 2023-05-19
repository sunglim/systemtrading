package koreainvestment

import (
	"fmt"
	"time"
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

func DemoCallFunction() {
	fmt.Printf("called %s\n", time.Now().String())

	/*
		accessToken := ApiGetAccessToken{}
		token := accessToken.Call()
	*/
	accessToken := ApiInqueryPrice{}
	token := accessToken.Call("005930")

	fmt.Print("AccessToken: " + token)
}

func getAlwaysValidAccessToken() string {
	if accessToken == "" {
		accessToken = ApiGetAccessToken{}.Call()
	}

	return "Bearer " + accessToken
}
