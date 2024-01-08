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
func CreateApiOrderCash(stockCode string, credential Credential,
	accountInfo KoreaInvestmentAccount, accessToken string) *ApiOrderCash {
	return &ApiOrderCash{
		stockCode: stockCode, amount: 1,
		KoreaInvestmentAccount: accountInfo,
		Credential:             credential,
		accessToken:            accessToken,
	}
}

func NewApiOrderCash(stockCode string, amount int, credential Credential,
	accountInfo KoreaInvestmentAccount, accessToken string) *ApiOrderCash {
	return &ApiOrderCash{
		stockCode: stockCode, amount: amount,
		KoreaInvestmentAccount: accountInfo,
		Credential:             credential,
		accessToken:            accessToken,
	}
}

type ApiOrderCash struct {
	stockCode string
	amount    int
	KoreaInvestmentAccount
	Credential
	accessToken string
}

func (api ApiOrderCash) StockCode() string {
	return api.stockCode
}

func (api ApiOrderCash) url() string {
	return ProductionDomain + "/uapi/domestic-stock/v1/trading/order-cash"
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
		CANO:         api.KoreaInvestmentAccount.CANO,
		ACNT_PRDT_CD: api.KoreaInvestmentAccount.ACNT_PRDT_CD,
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
	RtCd  string `json:"rt_cd"`
	Msg1  string `json:"msg1"`
	MsgCd string `json:"msg_cd"`
	// response time
	ResponseTime time.Time
}

func (response ApiOrderCashResponse) IsSuccess() bool {
	return response.RtCd == "0"
}

func (api ApiOrderCash) Call() *ApiOrderCashResponse {
	req, err := http.NewRequest(http.MethodPost, api.url(), api.buildRequestBody())
	if err != nil {
		panic(err)
	}
	req.Close = true
	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", api.accessToken)
	req.Header.Add("appkey", api.Credential.AppKey)
	req.Header.Add("appsecret", api.Credential.AppSecret)
	// order cash
	req.Header.Add("tr_id", "TTTC0802U")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		print("Order failed:", "reason", err.Error())
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
