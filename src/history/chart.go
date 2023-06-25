package history

import (
	"strings"
)

type Record struct {
	//Date,Open,High,Low,Close,Adj Close,Volume
	Date     string
	Open     string
	High     string
	Low      string
	Close    string
	AdjClose string
	Volume   string
}

func NewRecord(rawRecord string) Record {
	//2022-12-27,117.500000,119.669998,108.760002,109.099998,109.099998,208643400
	s := strings.Split(rawRecord, ",")
	return Record{
		Date:     s[0],
		Open:     s[1],
		High:     s[2],
		Low:      s[3],
		Close:    s[4],
		AdjClose: s[5],
		Volume:   s[6],
	}
}

type Sheet struct {
	record []Record
}

func NewSheet(rawSheet string) Sheet {
	records := strings.Split(rawSheet, "\n")
	sheet := Sheet{}
	for _, record := range records[1:] {
		sheet.record = append(sheet.record, NewRecord(record))
	}

	return sheet
}
