package order

type Equity = string

// Order types
const (
	PriceLimit  string = "pricelimit"
	MarketPrice string = "marketprice"
)

type Order struct {
	Equity
	KrwPrice  int
	Amount    int
	OrderType string
}

func NewMakretPriceOrder(equaty Equity, amount int) *Order {
	return &Order{Equity: equaty, Amount: amount, OrderType: MarketPrice}
}

func NewPriceLimitOrder(equaty Equity, amount int, krwPrice int) *Order {
	return &Order{Equity: equaty, KrwPrice: krwPrice, Amount: amount, OrderType: PriceLimit}
}
