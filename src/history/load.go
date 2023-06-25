package history

var historicRSI []int

// The raw data.
var RawSheet Sheet

// Loads historical data, convert them to structs easily managable.
func Initialize() {
	Load()

	Process()
}

func Load() {
	RawSheet = GetHistoricalData()
}

func Process() {
	processRSI()
}

// Get current RSI value.
func GetRSI(code string) {

}
