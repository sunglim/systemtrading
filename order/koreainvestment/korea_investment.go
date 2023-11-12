package koreainvestment

import (
	"fmt"
	"log"

	ki "github.com/sunglim/systemtrading/pkg/koreainvestment"
)

var productionUrl string

// var accessToken token
var accountInfo ki.KoreaInvestmentAccount

func GetDefaultAccount() ki.KoreaInvestmentAccount {
	return accountInfo
}

// Holds all necessary sub resources such as tokenstore.
var ki_package *ki.KoreaInvestment

func GetDefaultKoreaInvestmentInstance() *ki.KoreaInvestment {
	return ki_package
}

func Initialize(url, applicationKey, applicationSecret string, account ki.KoreaInvestmentAccount) {
	productionUrl = url
	accountInfo = account

	ki_package = ki.NewKoreaInvestment(ki.Credential{
		AppKey:    applicationKey,
		AppSecret: applicationSecret,
	}, log.Default())

	fmt.Printf("Initialize Korea investment trading.\n ProductionUrl[%s], AppKey[%s], AppSecret[%s], AccountInfo[%v]\n",
		productionUrl, ki_package.GetCredential().AppKey, ki_package.GetCredential().AppSecret, accountInfo)
}
