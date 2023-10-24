package koreainvestment

import "time"

func NewCachedAccessToken(credential Credential) *CachedAccessToken {
	token := NewToken(credential)
	return &CachedAccessToken{cachedToken: token}
}

type CachedAccessToken struct {
	cachedToken *Token
}

func (c *CachedAccessToken) GetToken() string {
	return c.cachedToken.Get()
}

func NewToken(credential Credential) *Token {
	return &Token{
		// Sets very old date to issue a new token when Get() is called for the first time.
		lastUpdatedTime: time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC),
		api:             NewApiGetAccessToken(credential),
	}

}

type Token struct {
	value           string
	lastUpdatedTime time.Time
	api             *ApiGetAccessToken
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
	response := t.api.Call()
	t.set(response.AccessToken)
}

func (t *Token) set(newTokenString string) {
	t.value = newTokenString
	t.lastUpdatedTime = time.Now()
}
