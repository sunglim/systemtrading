package koreainvestment

import "errors"

type token = string

// A store to get an always valid token.
type TokenStore struct {
	tokenCache *TokenCache
	api        *ApiGetAccessToken
}

func NewTokenStore(api *ApiGetAccessToken) *TokenStore {
	return &TokenStore{tokenCache: NewTokenCache(), api: api}
}

func (t *TokenStore) GetToken() (token, error) {
	token, err := t.tokenCache.Get()
	if err != nil {
		tokenResponse := t.api.Call()
		if tokenResponse.IsFailed() {
			return "", errors.New("failed to issue a token")
		}

		t.tokenCache.Set(tokenResponse.AccessToken)
		return tokenResponse.AccessToken, nil
	}

	return token, nil
}
