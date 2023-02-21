package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	measurements := map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
	return measurements
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	bill := make(map[string]int)

	return bill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	unitValue, unitExists := units[unit]
	if unitExists {
		bill[item] += unitValue
		return true
	} else {
		return false
	}
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	itemValue, itemExists := bill[item]
	unitValue, unitExists := units[unit]

	if !itemExists || !unitExists {
		return false
	}
	if itemValue-unitValue < 0 {
		return false
	} else {
		bill[item] -= unitValue
		if bill[item] == 0 {
			delete(bill, item)
		}
		return true
	}
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	itemInTheBill, itemExists := bill[item]

	if !itemExists {
		return 0, false
	} else {
		return itemInTheBill, true
	}
}
