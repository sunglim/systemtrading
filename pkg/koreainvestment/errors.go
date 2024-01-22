package koreainvestment

type ErrorCode = string

const (
	// 주문수량을 확인 하여 주십시요.
	OrderQuantityErrorCode = "APBK0986"

	// 장주문시간이 아닙니다
	OrderClosingTimeErrorCode = "APBK0919"
)
