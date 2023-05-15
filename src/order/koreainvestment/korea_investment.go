package koreainvestment

import (
	"fmt"
	"time"
)

var productionUrl string
var appKey string
var appSecret string

func Initialize(url, applicationKey, applicationSecret string) {
	productionUrl = url
	appKey = applicationKey
	appSecret = applicationSecret
}

func DemoCallFunction() {
	fmt.Printf("called %s\n", time.Now().String())

	accessToken := ApiGetAccessToken{}
	token := accessToken.Call()
	fmt.Print("AccessToken: " + token)
}
