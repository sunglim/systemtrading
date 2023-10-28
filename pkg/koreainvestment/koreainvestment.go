package koreainvestment

import (
	"log"
	"strconv"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/sunglim/systemtrading/internal/metrics"
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
	user             Credential
	token            string
	tokenExpire      time.Time
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
	if f.tokenExpire.Before(time.Now()) {
		f.logger.Printf("\nToken expire: %v", f.tokenExpire)
		return true
	}

	metrics.IssueToken()
	response := NewApiGetAccessToken(f.user).Call()
	f.token = response.AccessToken
	i, _ := strconv.ParseInt(response.ExpiresIn, 10, 64)
	f.tokenExpire = time.Now().Add(time.Second * time.Duration(i))
	if f.logger != nil {
		f.logger.Printf("\nset token %s: %s\n", time.Now().String(), f.token)
	}

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
