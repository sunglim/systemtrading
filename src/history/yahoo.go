package history

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetHistoricalData() Sheet {
	url := "https://query1.finance.yahoo.com/v7/finance/download/TSLA?period1=1671753600&period2=1687478400&interval=1d&events=history&includeAdjustedClose=true"
	buffer := bytes.NewBuffer([]byte(``))

	r, err := http.NewRequest("GET", url, buffer)
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", string(data))
	sheet := NewSheet(string(data))
	return sheet
}
