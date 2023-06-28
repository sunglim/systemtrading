package koreainvestment

// Implementatino of
//
import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type ApiInqueryPrice struct {
}

func (api ApiInqueryPrice) url(iscd string) string {
	url := productionUrl + "/uapi/domestic-stock/v1/quotations/inquire-price?fid_cond_mrkt_div_code=J"
	return fmt.Sprintf("%s&fid_input_iscd=%s", url, iscd)
}

func (api ApiInqueryPrice) buildRequestBody() *bytes.Buffer {
	body := []byte(fmt.Sprintf(`{
		"grant_type": "client_credentials",
		"appkey": "%s",
		"appsecret": "%s"
	}`, appKey, appSecret))

	return bytes.NewBuffer(body)
}

type output struct {
	StockPrsentPrice string `json:"stck_prpr"`
}

type InqueryPriceResponse struct {
	AccessToken string `json:"msg1"`
	Output      output `json:"output"`
}

func (api ApiInqueryPrice) Call(iscd string) string {
	r, err := http.NewRequest(http.MethodPost, api.url(iscd), api.buildRequestBody())
	if err != nil {
		panic(err)
	}
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("authorization", getAlwaysValidAccessToken())
	r.Header.Add("appkey", appKey)
	r.Header.Add("appsecret", appSecret)
	r.Header.Add("tr_id", "FHKST01010100")

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	post := &InqueryPriceResponse{}
	derr := json.NewDecoder(res.Body).Decode(post)
	if derr != nil {
		panic(derr)
	}
	return post.Output.StockPrsentPrice
}
