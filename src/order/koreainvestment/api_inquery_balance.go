package koreainvestment

// Implementatino of
//
import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type ApiInqueryBalance struct {
}

func (api ApiInqueryBalance) url() string {
	url := productionUrl + "/uapi/domestic-stock/v1/trading/inquire-balance?CANO=" + accountInfo.CANO
	url = fmt.Sprintf("%s&ACNT_PRDT_CD=%s", url, "01")
	url = fmt.Sprintf("%s&AFHR_FLPR_YN=%s", url, "N")
	url = fmt.Sprintf("%s&OFL_YN=%s", url, "")
	url = fmt.Sprintf("%s&INQR_DVSN=%s", url, "02")
	url = fmt.Sprintf("%s&UNPR_DVSN=%s", url, "01")
	url = fmt.Sprintf("%s&FUND_STTL_ICLD_YN=%s", url, "N")
	url = fmt.Sprintf("%s&FNCG_AMT_AUTO_RDPT_YN=%s", url, "N")
	url = fmt.Sprintf("%s&PRCS_DVSN=%s", url, "00")
	url = fmt.Sprintf("%s&CTX_AREA_FK100=", url)
	url = fmt.Sprintf("%s&CTX_AREA_NK100=%s", url, "")

	fmt.Println(url)
	return url
}

func (api ApiInqueryBalance) buildRequestBody() *bytes.Buffer {
	body := []byte(``)

	return bytes.NewBuffer(body)
}

type ApiInqueryBalanceResponseOutput struct {
	PdNo     string `json:"pdno"`
	PrdtName string `json:"prdt_name"`
	// current stock price
	Prpr string `json:"prpr"`
	// Average purchase price
	PchsAvgPric string `json:"pchs_avg_pric"`
	// prps - pchsavgpric
	EvluPflsAmt string `json:"evlu_pfls_amt"`
}

type ApiInqueryBalanceResponse struct {
	// is success.
	RtCd    string                            `json:"rt_cd"`
	Msg1    string                            `json:"msg1"`
	Output1 []ApiInqueryBalanceResponseOutput `json:"output1"`
	// response time
	ResponseTime time.Time
}

func (api ApiInqueryBalance) Call() *ApiInqueryBalanceResponse {
	r, err := http.NewRequest("GET", api.url(), api.buildRequestBody())
	if err != nil {
		panic(err)
	}
	r.Header.Add("Content-Type", "application/json")
	validToekn := getAlwaysValidAccessToken()
	r.Header.Add("authorization", validToekn)
	r.Header.Add("appkey", appKey)
	r.Header.Add("appsecret", appSecret)
	r.Header.Add("tr_id", "TTTC8434R")

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	post := &ApiInqueryBalanceResponse{}
	derr := json.NewDecoder(res.Body).Decode(post)
	if derr != nil {
		panic(derr)
	}
	post.ResponseTime = time.Now()
	return post
}
