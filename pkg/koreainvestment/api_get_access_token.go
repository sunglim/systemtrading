package koreainvestment

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type ApiGetAccessToken struct {
	credential Credential
	client     *http.Client
}

func NewApiGetAccessToken(credential Credential) *ApiGetAccessToken {
	return &ApiGetAccessToken{credential, http.DefaultClient}
}

func (api *ApiGetAccessToken) url() string {
	url := "/oauth2/tokenP"
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
	req, err := http.NewRequest(http.MethodPost, api.url(), api.buildRequestBody())
	if err != nil {
		// TODO: Return error instead of panic.
		panic(err)
	}
	req.Close = true
	req.Header.Add("Content-Type", "application/json")

	res, err := api.client.Do(req)
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
