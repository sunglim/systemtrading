package koreainvestment

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type ApiGetAccessToken struct {
	credential Credential
}

func NewApiGetAccessToken(credential Credential) *ApiGetAccessToken {
	return &ApiGetAccessToken{credential}
}

func (api *ApiGetAccessToken) url() string {
	url := "/oauth2/tokenP"
	if forTesting {
		return TestingDomain + url
	}
	return ProductionDomain + url
}

func (api *ApiGetAccessToken) buildRequestBody() *bytes.Buffer {
	body := []byte(fmt.Sprintf(`{
		"grant_type": "client_credentials",
		"appkey": "%s",
		"appsecret": "%s"
	}`, api.credential.AppKey, api.credential.AppSecret))

	return bytes.NewBuffer(body)
}

type GetAccessTokenResponse struct {
	AccessToken             string `json:"access_token"`
	AccessTokenTokenExpired string `json:"access_token_token_expired"`
}

func (api *ApiGetAccessToken) Call() *GetAccessTokenResponse {
	request, err := http.NewRequest(http.MethodPost, api.url(), api.buildRequestBody())
	if err != nil {
		// TODO: Return error instead of panic.
		panic(err)
	}
	request.Header.Add("Content-Type", "application/json")

	client := http.Client{}
	res, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	post := &GetAccessTokenResponse{}
	derr := json.NewDecoder(res.Body).Decode(post)
	if derr != nil {
		panic(derr)
	}
	return post
}
