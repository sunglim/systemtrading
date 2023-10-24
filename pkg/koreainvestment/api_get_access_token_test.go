package koreainvestment

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"
)

type fakeService func(*http.Request) (*http.Response, error)

func (s fakeService) RoundTrip(req *http.Request) (*http.Response, error) {
	return s(req)
}

var testResponse GetAccessTokenResponse = GetAccessTokenResponse{
	AccessToken:             "testAccessToken",
	AccessTokenTokenExpired: "testAccessTokenExpired",
}

var testResponseString = getTestResponseString(testResponse)

func getTestResponseString(response GetAccessTokenResponse) string {
	jsonBytes, _ := json.Marshal(testResponse)
	return string(jsonBytes)
}

func TestApiCallBasic(t *testing.T) {
	fakeClient := &http.Client{Transport: fakeService(func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
			Body: io.NopCloser(strings.NewReader(testResponseString)),
		}, nil
	})}
	api := ApiGetAccessToken{Credential{}, fakeClient}
	response := api.Call()
	if *response != testResponse {
		t.Error("api call response failed")
	}
}
