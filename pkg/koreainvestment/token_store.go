package koreainvestment

import (
	"errors"

	"github.com/sunglim/systemtrading/pkg/koreainvestment/tokencache"
)

type token = string

// A store to get an always valid token.
type TokenStore struct {
	tokenCache *tokencache.TokenCache
	api        *ApiGetAccessToken
}

func NewTokenStore(api *ApiGetAccessToken) *TokenStore {
	return &TokenStore{tokenCache: tokencache.NewTokenCache(), api: api}
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
