package koreainvestment

import (
	"log"
	"time"

	gocron "github.com/go-co-op/gocron"
	"sunglim.github.com/sunglim/systemtrading/internal/metrics"
)

var forTesting bool

const (
	ProductionDomain = "https://openapi.koreainvestment.com:9443"
	TestingDomain    = "https://openapivts.koreainvestment.com:29443"
)

type KoreaInvestmentAccount struct {
	// 종합계좌번호; 계좌번호 체계(8-2)의 앞 8자리
	CANO string

	// 계좌상품코드; 계좌번호 체계(8-2)의 뒤 2자리
	ACNT_PRDT_CD string
}

// Credential crednetials.
type Credential struct {
	AppKey    string
	AppSecret string
}

type KoreaInvestment struct {
	user             Credential
	token            string
	tokenRefreshHour int
	logger           *log.Logger
}

func (f *KoreaInvestment) GetCredential() Credential {
	return f.user
}

func (f *KoreaInvestment) GetBearerAccessToken() string {
	return "Bearer " + f.token
}

func NewKoreaInvestment(user Credential, logger *log.Logger) *KoreaInvestment {
	return &KoreaInvestment{user: user, token: "", tokenRefreshHour: 10, logger: logger}
}

func NewKoreaInvestmentTokenRefresh(user Credential, tokenRefreshHour int) *KoreaInvestment {
	return &KoreaInvestment{user: user, token: "", tokenRefreshHour: tokenRefreshHour}
}

func (f *KoreaInvestment) setAccessToken() bool {
	metrics.IssueToken()
	response := NewApiGetAccessToken(f.user).Call()
	f.token = response.AccessToken
	f.logger.Printf("\nset token %s: %s\n", time.Now().String(), f.token)

	return true
}

func (f *KoreaInvestment) InitializeToken() bool {
	f.setAccessToken()

	// Refresh token periodically.
	s := gocron.NewScheduler(time.UTC).Every(f.tokenRefreshHour).Hour()
	s.Do(f.setAccessToken)
	s.StartAsync()

	return true
}
