package machine

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

//This function will run the testcases mentioned in json files
func InitializeAndRunTests() {
	//create machine from machine.json
	CoffeeMachine = CreateMachineFromFile("machine.json")
	//process orders from order.json
	ProcessOrders("order.json")
	//refill ingredient mentioned in refill.json
	ProcessRefill("refill.json")

	ProcessOrders("order.json")
	fmt.Println("total_orders: ", total_orders)
}

//This function parses machine input json file and returns a *machine which can be used to place orders
func CreateMachineFromFile(machineSpecsFile string) *machine{

	var spec Spec
	specs, _ := ioutil.ReadFile(machineSpecsFile)
	json.Unmarshal(specs, &spec)

	machine_obj := spec.Machine
	outlets_obj := machine_obj.Outlets
	beverage_obj := machine_obj.Beverages
	total_items_obj := machine_obj.TotalItemsQuantity

	var machine_spec map[string]interface{}

	specs, _ = ioutil.ReadFile(machineSpecsFile)
	json.Unmarshal(specs, &machine_spec)

	coffeeMachine := &machine{
		totalIngredientsLeft: total_items_obj,
		outlets_num:          outlets_obj.Count_n,
		ingredientsMutex:     sync.RWMutex{},
		beverages: make(map[string]*beverage,len(beverage_obj)),
		outlets: make([]*outlet,outlets_obj.Count_n),
	}

	//initialize outlets
	for id := range coffeeMachine.outlets {
		coffeeMachine.outlets[id] = newOutlet(id)
	}
	//store beverages related data in machine struct
	for nameOfBeverage := range beverage_obj {
		ingredients := beverage_obj[nameOfBeverage]
		coffeeMachine.beverages[nameOfBeverage] = newBeverage(nameOfBeverage,ingredients)
	}

	return coffeeMachine
}

//this function reads the order.json file to place orders for beverages in the machine
func ProcessOrders(orderFile string) {
	var orders Orders
	jsonFile, _ := os.Open(orderFile)
	order_file,_ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(order_file,&orders)
	for _,order := range orders.Orders {
		CoffeeMachine.serve(order.Name)
		//time.Sleep(10000)
	}
}

func ProcessRefill(refillFile string) {
	var refill Refill
	jsonFile, _ := os.Open(refillFile)
	refill_data,_ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(refill_data,&refill)
	CoffeeMachine.refillIngredient(refill.Item.Name,refill.Item.Qty)
	fmt.Println(refill.Item.Name+" refilled with quantity: ",refill.Item.Qty)
}
