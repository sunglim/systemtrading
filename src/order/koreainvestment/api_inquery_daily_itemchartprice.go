package koreainvestment

// 국내주식기간별시세(일/주/월/년)
// Implementatino of
// /uapi/domestic-stock/v1/quotations/inquire-daily-itemchartprice

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type ApiInqueryDailyItemChartPrice struct {
}

func (api ApiInqueryDailyItemChartPrice) url(iscd string) string {
	url := productionUrl + "/uapi/domestic-stock/v1/quotations/inquire-daily-itemchartprice"
	url += url + "?fid_cond_mrkt_div_code=J"
	url += url + "&FID_INPUT_ISCD=" + iscd
	// start date
	// end date
	url += url + "&FID_INPUT_DATE_1=" + time.Now().AddDate(0, 0, -50).Format("202201")
	url += url + "&FID_INPUT_DATE_2=" + time.Now().Format("202201")
	// Daily
	url += url + "&FID_PERIOD_DIV_CODE=D"
	url += url + "&FID_ORG_ADJ_PRC=0"
	return url
}

func (api ApiInqueryDailyItemChartPrice) buildRequestBody() *bytes.Buffer {
	body := []byte(fmt.Sprintf(`{
		"grant_type": "client_credentials",
		"appkey": "%s",
		"appsecret": "%s"
	}`, appKey, appSecret))

	return bytes.NewBuffer(body)
}

type ApiInqueryDailyItemChartPriceResponseOutput struct {
	StockPresentPrice string `json:"stck_prpr"`
}

type ApiInqueryDailyItemChartPriceResponseOutput2 struct {
	StockPresentPrice string `json:"stck_clpr"`
}

type ApiInqueryDailyItemChartPriceResponse struct {
	Output1 ApiInqueryDailyItemChartPriceResponseOutput  `json:"output1"`
	Output2 ApiInqueryDailyItemChartPriceResponseOutput2 `json:"output2"`
}

func (api ApiInqueryDailyItemChartPrice) Call(iscd string) *ApiInqueryDailyItemChartPriceResponse {
	r, err := http.NewRequest(postMethod, api.url(iscd), api.buildRequestBody())
	if err != nil {
		panic(err)
	}
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("authorization", getAlwaysValidAccessToken())
	r.Header.Add("appkey", appKey)
	r.Header.Add("appsecret", appSecret)
	r.Header.Add("tr_id", "FHKST03010100")
	r.Header.Add("custtype", "P")

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	post := &ApiInqueryDailyItemChartPriceResponse{}
	derr := json.NewDecoder(res.Body).Decode(post)
	if derr != nil {
		panic(derr)
	}
	return post
}
