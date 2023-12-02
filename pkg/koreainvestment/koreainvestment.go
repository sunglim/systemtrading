package koreainvestment

import (
	"fmt"
	"log"
)

var forTesting bool

const (
	ProductionDomain = "https://openapi.koreainvestment.com:9443"
	TestingDomain    = "https://openapivts.koreainvestment.com:29443"
)

// Credential crednetials.
type Credential struct {
	AppKey    string
	AppSecret string
}

type KoreaInvestment struct {
	user       Credential
	logger     *log.Logger
	tokenStore *TokenStore
}

func (f *KoreaInvestment) GetCredential() Credential {
	return f.user
}

func (f *KoreaInvestment) GetBearerAccessToken() string {
	token, err := f.tokenStore.GetToken()
	if err != nil {
		fmt.Println("Get token failed: " + err.Error())
		return ""
	}

	return "Bearer " + token
}

func NewKoreaInvestment(user Credential, logger *log.Logger) *KoreaInvestment {
	return &KoreaInvestment{user: user, tokenStore: NewTokenStore(NewApiGetAccessToken(user)), logger: logger}
}
