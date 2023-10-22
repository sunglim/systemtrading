package koreainvestment

import (
	"testing"
)

func TestCachedToken(t *testing.T) {
	credential := Credential{
		AppKey:    "dummy",
		AppSecret: "dummy",
	}
	token := NewToken(credential)

	// First run calls an API internally.
	token.Get()
}
