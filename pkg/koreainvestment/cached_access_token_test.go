package koreainvestment

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"
)

type FakeApiGetAccessToken struct {
	credential Credential
	client     *http.Client
}

func (api *FakeApiGetAccessToken) Call() *GetAccessTokenResponse {
	return nil
}

func TestGetToken(t *testing.T) {
	jsonBytes, _ := json.Marshal(testResponse)
	fakeClient := &http.Client{Transport: fakeService(func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
			Body: io.NopCloser(strings.NewReader(string(jsonBytes))),
		}, nil
	})}
	//api := ApiGetAccessToken{Credential{}, fakeClient}
	api := FakeApiGetAccessToken{Credential{}, fakeClient}
	token := &Token{
		api: &api,
	}
	cachedToken := CachedAccessToken{cachedToken: token}
	if cachedToken.GetToken() != testResponse.AccessToken {
		t.Errorf("fail %s, %s", cachedToken.GetToken(), testResponse.AccessToken)
	}
}
