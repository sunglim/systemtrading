package order

type CodeAndQuantity struct {
	Code     string `yaml:"code"`
	Quantity int    `yaml:"quantity"`
}

type CodeAndQuantityAndPrice struct {
	Code     string `yaml:"code"`
	Price    int    `yaml:"price"`
	Quantity int    `yaml:"quantity"`
}
