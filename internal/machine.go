package internal

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

const (
	DefaultServiceCoinsAmount = 5
	DefaultServiceItemsAmount = 5
	NotAvailableProduct       = Error("Not available product")
)

type Machine struct {
	cashEngine    *CashEngine
	itemCatalogue *ItemCatalogue
	itemStock     map[int]int
	logger        *log.Logger
}

// Instantiates a new vending machine, including its CashEngine, ItemCatalogue and ItemStock
func NewMachine(catalogue *ItemCatalogue, engine *CashEngine, logger *log.Logger) *Machine {
	itemStock := make(map[int]int)
	for _, i := range catalogue.Items {
		itemStock[i.selector] = 0
	}

	return &Machine{
		cashEngine:    engine,
		itemCatalogue: catalogue,
		itemStock:     itemStock,
		logger:        logger,
	}

}

// ExecuteCommand interprets the commands that arrives to the VendingMachine to execute the desired actions
func (m *Machine) ExecuteCommand(command string) {
	separateRegexp := `[\w.-]+`
	r := regexp.MustCompile(separateRegexp)
	allCommands := r.FindAllString(command, -1)

	for _, s := range allCommands {
		_, err := strconv.ParseFloat(s, 32)
		if err != nil {
			action := strings.Split(s, "-")
			if strings.EqualFold(action[0], "GET") && len(action) > 1 {
				err := m.SellItem(action[1])
				if err != nil {
					m.logger.Printf("Error selling item, %s", err)
				}
			} else if strings.EqualFold(action[0], "RETURN") && len(action) > 1 && strings.EqualFold(action[0], "COIN") {
				m.ReturnCoins()
			} else if strings.EqualFold(s, "SERVICE") {
				m.DefaultService()
				m.logger.Println("SERVICE in progress")
			} else {
				m.logger.Println("Invalid command, please try again")
			}
		} else {
			err = m.InsertCoins(s)
			if err != nil {
				m.logger.Printf("%s, %s", err, s)
			}
		}

	}

}

// InsertCoins is used to load money in the machine before buying items.
func (m *Machine) InsertCoins(coins ...string) error {
	for _, c := range coins {
		err := m.cashEngine.InsertCoins(c)
		if err != nil {
			return err
		}
	}

	return nil
}

// ReturnCoins returns the inserted coins in the last service
func (m *Machine) ReturnCoins() {
	coins := m.cashEngine.DropCoins()

	s := fmt.Sprintf("-> ")
	for i, c := range coins {
		if i == len(coins)-1 {
			s += fmt.Sprintf("%s", c.value.String())
			continue
		}
		s += fmt.Sprintf("%s, ", c.value.String())
	}

	m.logger.Println(s)
}

// LoadItem inserts items in the machine before it starts or during a SERVICE action
func (m *Machine) LoadItem(item Item, amount int) {
	m.itemStock[item.selector] = amount
}

// LoadItem stores money inside the machine before it starts or during a SERVICE action
func (m *Machine) LoadMoney(coins ...string) error {
	for _, c := range coins {
		newCoin, err := NewCoin(c)
		if err != nil {
			return err
		}

		m.cashEngine.StoreCoins(newCoin)
	}

	return nil
}

// SellItem serves point of connection between Machine and CashEngine to sell items to the customers
func (m *Machine) SellItem(name string) error {
	item, err := m.itemCatalogue.GetItemByName(name)
	if err != nil {
		return err
	}

	if m.itemStock[item.selector] == 0 {
		m.ReturnCoins()
		return NotAvailableProduct
	}

	coin, err := m.cashEngine.SellItem(item.price)
	if err != nil {
		return err
	}

	m.itemStock[item.selector]--
	s := fmt.Sprintf("-> %s", strings.ToUpper(name))
	for _, c := range coin {
		s += fmt.Sprintf(", %s", c.value.String())
	}

	m.logger.Println(s)

	return nil
}

// DefaultService adds a default number of each item and coin to the vending machine
func (m *Machine) DefaultService() {
	for _, i := range m.itemCatalogue.Items {
		m.LoadItem(i, DefaultServiceItemsAmount)
	}

	for _, vc := range m.cashEngine.validCoins {
		c, err := NewCoin(vc)
		if err != nil {
			m.logger.Fatalf("could not load coins, %s", err)
		}
		m.cashEngine.LoadCoins(c, DefaultServiceCoinsAmount)
	}
}
