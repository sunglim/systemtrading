package koreainvestment

// Implementation of https://apiportal.koreainvestment.com/apiservice/apiservice-domestic-stock#L_aade4c72-5fb7-418a-9ff2-254b4d5f0ceb

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

// deprecated
func CreateApiOrderCash(stockCode string) *ApiOrderCash {
	return &ApiOrderCash{stockCode: stockCode, amount: 1}
}

func NewApiOrderCash(stockCode string, amount int) *ApiOrderCash {
	return &ApiOrderCash{stockCode: stockCode, amount: amount}
}

type ApiOrderCash struct {
	stockCode string
	amount    int
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
		ORD_QTY:      strconv.Itoa(api.amount),
		ORD_UNPR:     "0",
	}
	b, _ := json.Marshal(body)
	return bytes.NewBuffer(b)
}

type ApiOrderCashResponse struct {
	// is success.
	RtCd string `json:"rt_cd"`
	Msg1 string `json:"msg1"`
	// response time
	ResponseTime time.Time
}

func (response ApiOrderCashResponse) IsSuccess() bool {
	return response.RtCd == "0"
}

func (api ApiOrderCash) Call() *ApiOrderCashResponse {
	url := api.url()
	r, err := http.NewRequest(http.MethodPost, url, api.buildRequestBody())
	if err != nil {
		panic(err)
	}
	r.Header.Add("content-type", "application/json")
	r.Header.Add("authorization", ki_package.GetBearerAccessToken())
	r.Header.Add("appkey", ki_package.GetCredential().AppKey)
	r.Header.Add("appsecret", ki_package.GetCredential().AppSecret)
	// order cash
	r.Header.Add("tr_id", "TTTC0802U")

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	post := &ApiOrderCashResponse{}
	derr := json.NewDecoder(res.Body).Decode(post)
	if derr != nil {
		panic(derr)
	}
	post.ResponseTime = time.Now()
	return post
}
