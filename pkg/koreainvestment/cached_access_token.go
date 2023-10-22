package koreainvestment

import "time"

func NewCachedAccessToken(credential Credential) *CachedAccessToken {
	token := Token{
		value:           "",
		lastUpdatedTime: time.Time{},
		credential:      credential,
	}
	return &CachedAccessToken{cachedToken: token}
}

type CachedAccessToken struct {
	cachedToken Token
}

func (c *CachedAccessToken) GetToken() string {
	return c.cachedToken.value
}

func NewToken(credential Credential) *Token {
	return &Token{
		// Sets very old date to issue a new token when Get() is called for the first time.
		lastUpdatedTime: time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC),
		credential:      credential,
	}

}

type Token struct {
	value           string
	lastUpdatedTime time.Time
	credential      Credential
}

func (t *Token) Get() string {
	now := time.Now()
	if now.Sub(t.lastUpdatedTime).Minutes() < 6 {
		return t.value
	}

	t.refresh()
	return t.value
}

func (t *Token) refresh() {
	api := NewApiGetAccessToken(t.credential)
	response := api.Call()
	t.set(response.AccessToken)
}

func (t *Token) set(newTokenString string) {
	t.value = newTokenString
	t.lastUpdatedTime = time.Now()
}
