package koreainvestment

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type ApiGetAccessToken struct {
}

func (api ApiGetAccessToken) url() string {
	return productionUrl + "/oauth2/tokenP"
}

func (api ApiGetAccessToken) buildRequestBody() *bytes.Buffer {
	body := []byte(fmt.Sprintf(`{
		"grant_type": "client_credentials",
		"appkey": "%s",
		"appsecret": "%s"
	}`, appKey, appSecret))

	return bytes.NewBuffer(body)
}

type GetAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
}

func (api ApiGetAccessToken) Call() string {
	r, err := http.NewRequest(postMethod, api.url(), api.buildRequestBody())
	if err != nil {
		panic(err)
	}
	r.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	post := &GetAccessTokenResponse{}
	derr := json.NewDecoder(res.Body).Decode(post)
	if derr != nil {
		panic(derr)
	}
	return post.AccessToken
}
