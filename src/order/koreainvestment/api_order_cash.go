package koreainvestment

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type ApiOrdeCash struct {
	stockCode string
}

func (api ApiOrdeCash) url() string {
	return productionUrl + "/uapi/domestic-stock/v1/trading/order-cash"
}

func (api ApiOrdeCash) buildRequestBody() *bytes.Buffer {
	body := []byte(fmt.Sprintf(`{
		"grant_type": "client_credentials",
		"appkey": "%s",
		"appsecret": "%s"
	}`, appKey, appSecret))

	return bytes.NewBuffer(body)
}

type ApiOrdeCashResponse struct {
	// is success.
	RtCd string `json:"rt_cd"`
}

func (api ApiOrdeCash) Call() *ApiOrdeCashResponse {
	r, err := http.NewRequest(postMethod, api.url(), api.buildRequestBody())
	if err != nil {
		panic(err)
	}
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("authorization", getAlwaysValidAccessToken())
	r.Header.Add("appkey", appKey)
	r.Header.Add("appsecret", appSecret)
	// order cash
	r.Header.Add("tr_id", "TTTC0802U")
	r.Header.Add("custtype", "P")

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	post := &ApiOrdeCashResponse{}
	derr := json.NewDecoder(res.Body).Decode(post)
	if derr != nil {
		panic(derr)
	}
	return post
}
