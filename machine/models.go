package machine

//to read input jsons
type Spec struct {
	Machine Machine `json:"machine"`
}

type Machine struct {
	Outlets            Outlets                   `json:"outlets"`
	TotalItemsQuantity map[string]int            `json:"total_items_quantity"`
	Beverages          map[string]map[string]int `json:"beverages"`
}

type Outlets struct {
	Count_n 			int	`json:"count_n"`
}

type Orders struct {
	Orders 	[]Order `json:"orders"`
}

type Order struct {
	Name 	string	`json:"name"`
}

type Refill struct {
	Item	Item	`json:"item"`
}

type Item struct {
	Name	string `json:"name"`
	Qty		int	`json:"qty"`
}



