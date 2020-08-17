package tests

import (
	"awesomeProject/machine"
	"fmt"
	"testing"
)

func TestCreateMachine(t *testing.T) {
	machine.CoffeeMachine = machine.CreateMachineFromFile("machineSpec-test.json")
}

func TestPlaceOrders(t *testing.T) {
	machine.CoffeeMachine = machine.CreateMachineFromFile("machineSpec-test.json")
	go machine.ProcessOrders("order-test.json")
	machine.ProcessOrders("order-test.json")
	go machine.ProcessOrders("order-test-1.json")
	machine.ProcessOrders("order-test-1.json")
	fmt.Println("ok")
}
