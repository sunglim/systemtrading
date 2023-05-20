package koreainvestment

// Implementation of https://apiportal.koreainvestment.com/apiservice/apiservice-domestic-stock#L_aade4c72-5fb7-418a-9ff2-254b4d5f0ceb

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func CreateApiOrderCash(stockCode string) *ApiOrderCash {
	return &ApiOrderCash{stockCode: stockCode}
}

type ApiOrderCash struct {
	stockCode string
}

func (api ApiOrderCash) url() string {
	return productionUrl + "/uapi/domestic-stock/v1/trading/order-cash"
}

func (api ApiOrderCash) buildRequestBody() *bytes.Buffer {
	body := struct {
		CANO         string
		ACNT_PRDT_CD string
		PDNO         string
		ORD_DVSN     string
		ORD_QTY      string
		ORD_UNPR     string
	}{
		CANO:         accountInfo.CANO,
		ACNT_PRDT_CD: accountInfo.ACNT_PRDT_CD,
		PDNO:         api.stockCode,
		ORD_DVSN:     "01",
		ORD_QTY:      "1",
		ORD_UNPR:     "0",
	}
	b, _ := json.Marshal(body)
	return bytes.NewBuffer(b)
}

type ApiOrdeCashResponse struct {
	// is success.
	RtCd string `json:"rt_cd"`
}

func (api ApiOrderCash) Call() *ApiOrdeCashResponse {
	url := api.url()
	r, err := http.NewRequest(postMethod, url, api.buildRequestBody())
	if err != nil {
		panic(err)
	}
	token := getAlwaysValidAccessToken()
	r.Header.Add("content-type", "application/json")
	r.Header.Add("authorization", token)
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
