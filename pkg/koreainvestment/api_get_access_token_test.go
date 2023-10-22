package koreainvestment

import (
	"io"
	"net/http"
	"strings"
	"testing"
)

type fakeService func(*http.Request) (*http.Response, error)

func (s fakeService) RoundTrip(req *http.Request) (*http.Response, error) {
	return s(req)
}

func TestApiCallBasic(t *testing.T) {
	fakeClient := &http.Client{Transport: fakeService(func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
			Body: io.NopCloser(strings.NewReader(`{"access_token": "abcd", "access_token_token_expired": "foo"}`)),
		}, nil
	})}
	api := ApiGetAccessToken{Credential{}, fakeClient}
	api.Call()
}
