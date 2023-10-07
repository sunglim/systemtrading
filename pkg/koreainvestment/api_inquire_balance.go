package koreainvestment

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// Implementation of https://apiportal.koreainvestment.com/apiservice/apiservice-domestic-stock#L_66c61080-674f-4c91-a0cc-db5e64e9a5e6

func NewApiInquireBalance(account KoreaInvestmentAccount, credential Credential, accessToken string) *ApiInquireBalance {
	return &ApiInquireBalance{
		KoreaInvestmentAccount: account,
		Credential:             credential,
		accessToken:            accessToken,
	}
}

type ApiInquireBalance struct {
	KoreaInvestmentAccount
	Credential
	accessToken string
}

func (api ApiInquireBalance) url() string {
	var sb strings.Builder
	sb.WriteString(ProductionDomain + "/uapi/domestic-stock/v1/trading/inquire-balance?CANO=" + api.CANO)
	sb.WriteString(fmt.Sprintf("&ACNT_PRDT_CD=%s", api.ACNT_PRDT_CD))
	sb.WriteString(fmt.Sprintf("&AFHR_FLPR_YN=%s", "N"))
	sb.WriteString(fmt.Sprintf("&OFL_YN=%s", ""))
	sb.WriteString(fmt.Sprintf("&INQR_DVSN=%s", "02"))
	sb.WriteString(fmt.Sprintf("&UNPR_DVSN=%s", "01"))
	sb.WriteString(fmt.Sprintf("&FUND_STTL_ICLD_YN=%s", "N"))
	sb.WriteString(fmt.Sprintf("&FNCG_AMT_AUTO_RDPT_YN=%s", "N"))
	sb.WriteString(fmt.Sprintf("&PRCS_DVSN=%s", "00"))
	sb.WriteString(fmt.Sprintf("&CTX_AREA_FK100=%s", ""))
	sb.WriteString(fmt.Sprintf("&CTX_AREA_NK100=%s", ""))

	return sb.String()
}

func (api ApiInquireBalance) buildRequestBody() *bytes.Buffer {
	return bytes.NewBuffer([]byte(``))
}

type ApiInquireBalanceResponseOutput struct {
	PdNo     string `json:"pdno"`
	PrdtName string `json:"prdt_name"`
	// current stock price
	Prpr string `json:"prpr"`
	// Average purchase price
	PchsAvgPric string `json:"pchs_avg_pric"`
	// prps - pchsavgpric
	EvluPflsAmt string `json:"evlu_pfls_amt"`
	// The percentage of gain
	EvluPflsRt string `json:"evlu_pfls_rt"`
	HldgQty    string `json:"hldg_qty"`
}

type ApiInquireBalanceResponse struct {
	// is success.
	RtCd    string                            `json:"rt_cd"`
	Msg1    string                            `json:"msg1"`
	Output1 []ApiInquireBalanceResponseOutput `json:"output1"`
	// response time
	ResponseTime time.Time
}

func (response ApiInquireBalanceResponse) IsSucess() bool {
	return response.RtCd == "0"
}

// Call() returns the result of API call.
// If failed, reutrns {nil, error}
func (api ApiInquireBalance) Call() (*ApiInquireBalanceResponse, error) {
	req, err := http.NewRequest(http.MethodGet, api.url(), api.buildRequestBody())
	if err != nil {
		panic(err)
	}
	req.Close = true
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("authorization", api.accessToken)
	req.Header.Add("appkey", api.AppKey)
	req.Header.Add("appsecret", api.AppSecret)
	req.Header.Add("tr_id", "TTTC8434R")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	post := &ApiInquireBalanceResponse{}
	derr := json.NewDecoder(res.Body).Decode(post)
	if derr != nil {
		return nil, derr
	}
	post.ResponseTime = time.Now()
	return post, nil
}
