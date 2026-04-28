package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen":3,
		"half_of_a_dozen":6,
		"dozen":12,
		"small_gross":120,
		"gross":144,
		"great_gross":1728,
	}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	score, unitExists := units[unit]
	if !unitExists {
		return false
	}
	billValue, billExists := bill[item]
	if !billExists {
		bill[item] = score
		return true
	}
	bill[item] = billValue + score
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	score, unitExists := units[unit]
	if !unitExists {
		return false
	}
	billValue, billExists := bill[item]
	if !billExists {
		return false
	}
	newValue := billValue - score
	switch {
	case newValue == 0:
		delete(bill, item)
		return true
	case newValue < 0:
		return false
	default:
		bill[item] = newValue
		return true
	}
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	billValue, billExists := bill[item]
	if !billExists {
        return 0, false
	}
    return billValue, true
}
