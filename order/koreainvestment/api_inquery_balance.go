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

/*
{
            "pdno": "001750",
            "prdt_name": "한양증권",
            "trad_dvsn_name": "현금",
            "bfdy_buy_qty": "0",
            "bfdy_sll_qty": "0",
            "thdt_buyqty": "0",
            "thdt_sll_qty": "19",
            "hldg_qty": "1",
            "ord_psbl_qty": "1",
            "pchs_avg_pric": "8671.0000",
            "pchs_amt": "8671",
            "prpr": "8820",
            "evlu_amt": "8820",
            "evlu_pfls_amt": "149",
            "evlu_pfls_rt": "1.71",
            "evlu_erng_rt": "0.00000000",
            "loan_dt": "",
            "loan_amt": "0",
            "stln_slng_chgs": "0",
            "expd_dt": "",
            "fltt_rt": "1.37931034",
            "bfdy_cprs_icdc": "120",
            "item_mgna_rt_name": "100%",
            "grta_rt_name": "불가",
            "sbst_pric": "6090",
            "stck_loan_unpr": "0.0000"
        },
*/

type ApiInqueryBalanceResponseOutput struct {
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
	HldgQty string `json:"hldg_qty"`
}

type ApiInqueryBalanceResponse struct {
	// is success.
	RtCd    string                            `json:"rt_cd"`
	Msg1    string                            `json:"msg1"`
	Output1 []ApiInqueryBalanceResponseOutput `json:"output1"`
	// response time
	ResponseTime time.Time
}

func (response ApiInqueryBalanceResponse) IsSucess() bool {
	return response.RtCd == "0"
}

func (api ApiInqueryBalance) Call() *ApiInqueryBalanceResponse {
	r, err := http.NewRequest(http.MethodGet, api.url(), api.buildRequestBody())
	if err != nil {
		panic(err)
	}
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("authorization", ki_package.GetBearerAccessToken())
	r.Header.Add("appkey", ki_package.GetCredential().AppKey)
	r.Header.Add("appsecret", ki_package.GetCredential().AppSecret)
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
