package main

import (
	"bufio"
	"fmt"
	"github.com/adolsalamanca/vending-machine/internal"
	"log"
	"os"
)

func initializeMachine() *internal.Machine {
	water := internal.NewItem(internal.Water, internal.WaterPrice, 1)
	juice := internal.NewItem(internal.Juice, internal.JuicePrice, 2)
	soda := internal.NewItem(internal.Soda, internal.SodaPrice, 3)

	catalogue := internal.NewItemCatalogue()
	catalogue.AddItem(water)
	catalogue.AddItem(juice)
	catalogue.AddItem(soda)

	engine := internal.NewCashEngine(internal.FiveCent, internal.TenCent, internal.TwentyFiveCent, internal.OneUnit)

	logger := log.Logger{}
	logger.SetOutput(os.Stdout)

	m := internal.NewMachine(catalogue, engine, &logger)

	for i := 0; i < 10; i++ {
		err := m.LoadMoney(internal.FiveCent, internal.TenCent, internal.TwentyFiveCent, internal.OneUnit)
		if err != nil {
			log.Fatalf("could not load money, %s", err)
		}
	}

	m.LoadItem(water, 5)
	m.LoadItem(juice, 5)
	m.LoadItem(soda, 5)

	return m
}

func main() {
	m := initializeMachine()
	fmt.Println("Vending machine initialized")

	for {
		buf := bufio.NewReader(os.Stdin)
		command, err := buf.ReadBytes('\n')
		if err != nil {
			fmt.Println(err)
		} else {
			m.ExecuteCommand(fmt.Sprintf("%s", command))
		}
	}
}
