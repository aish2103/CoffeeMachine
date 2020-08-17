package machine

type beverage struct {
	name	string
	ingredients	map[string]int
}

func newBeverage(name string,ingredients map[string]int) *beverage {
	return &beverage{
		name:name,
		ingredients:ingredients,
	}
}
