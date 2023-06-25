package order

// Condition
// is a condtion to be applied to Order

type Condition struct {
}

func (c Condition) Filter( /*Single stock info*/ ) bool {
	// if average is minus

	//single.average <= current
	// return false
	return true
}
