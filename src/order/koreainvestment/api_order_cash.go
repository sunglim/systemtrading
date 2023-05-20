package koreainvestment

// Implementation of https://apiportal.koreainvestment.com/apiservice/apiservice-domestic-stock#L_aade4c72-5fb7-418a-9ff2-254b4d5f0ceb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type ApiOrderCash struct {
	stockCode string
}

func (api ApiOrderCash) url() string {
	return productionUrl + "/uapi/domestic-stock/v1/trading/order-cash"
}

func (api ApiOrderCash) buildRequestBody() *bytes.Buffer {
	body := []byte(fmt.Sprintf(`{
		"grant_type": "client_credentials",
		"CANO": "%s",
		"ACNT_PRDT_CD": "%s",
		"PDNO": "%s",
		"ORD_DVSN": "01",
		"ORD_QTY": "1",
		"ORD_UNPR": "0",
	}`, accountInfo.CANO, accountInfo.ACNT_PRDT_CD, api.stockCode))

	return bytes.NewBuffer(body)
}

type ApiOrdeCashResponse struct {
	// is success.
	RtCd string `json:"rt_cd"`
}

func (api ApiOrderCash) Call() *ApiOrdeCashResponse {
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
