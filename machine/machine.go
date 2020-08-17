package machine

import (
	"fmt"
	"sync"
)

//this variable denotes coffee machine
var CoffeeMachine *machine
var total_orders = 0

type machine struct {
	totalIngredientsLeft	map[string]int
	beverages				map[string]*beverage
	outlets_num 			int
	outlets					[]*outlet
	ingredientsMutex 		sync.RWMutex
}

//constructor
func newMachine(totalIngredientsLeft map[string]int, beverages map[string]*beverage,outlets_num int, outlets []*outlet) *machine {
	return &machine{
		totalIngredientsLeft: totalIngredientsLeft,
		beverages:            beverages,
		outlets_num:          outlets_num,
		outlets:              outlets,
		ingredientsMutex:     sync.RWMutex{},
	}

}

//this function checks if the machine contains the ingredient in sufficient quantity
//if the machine doesn't, this function returns false and posts a Refill log
func (m *machine)isIngredientSufficient(name string, qty int) bool {
	if(m.totalIngredientsLeft[name] >= qty) {
		return true
	}
	fmt.Println("Refill "+name+"!")
	return false
}


//this function updates the ingredient qty left in the machine storage after using for this beverage.
func (m *machine)useIngredient(name string, qty int) {
		m.totalIngredientsLeft[name] -= qty
}

//this function will refill the ingredient in the machine
func (m *machine)refillIngredient(name string, qty int) {
	m.ingredientsMutex.Lock()
	defer m.ingredientsMutex.Unlock()

	if _,ok := m.totalIngredientsLeft[name]; ok{
		m.totalIngredientsLeft[name] += qty
	} else {
		fmt.Println("Wrong ingredient!")
	}
}

//this function finds an empty outlet else returns
//then checks if ingredients left in the machine are enough for the beverage
//then prepares the beverage
//this function uses ingredientsMutex to make sure that at a time only one thread makes the beverage
func (m *machine)serve(b_name string) {
	total_orders += 1
	b,ok := m.beverages[b_name]
	if !ok {
		fmt.Println("Wrong beverage selected")
		return
	}


	//////////get a free outlet/////////
	outlet_id := m.getOutlet()
	if	outlet_id == -1 {
		fmt.Println("No outlet free")
		return
	}
	fmt.Println("Serving "+b_name+" in outlet ",outlet_id)

	//////////get the ingredients of the beverage/////////
	ingredients := b.ingredients

	m.ingredientsMutex.Lock()
	defer m.ingredientsMutex.Unlock()

	//we will first check if the machine has enough ingredients
	for name := range ingredients {
		//check if machine has enough qty
		if !m.isIngredientSufficient(name, ingredients[name]) {
			fmt.Println(b_name+" cannot be prepared because "+name+" is not available")
			//freeing the outlet
			m.outlets[outlet_id].changeOutletStatus(false)
			return
		}
	}
	//all ingredients are sufficient, let's make the beverage
	for name := range ingredients {
		m.useIngredient(name,int(ingredients[name]))
	}

	fmt.Println(b_name+" is prepared")
	//freeing the outlet
	m.outlets[outlet_id].changeOutletStatus(false)
}

//this function gets the first empty outlet for serving the ordered drink
func (m *machine)getOutlet() int{
	for id,outlet := range m.outlets {
		if outlet.getOutlet() {
			return id
		}
	}
	return -1
}

