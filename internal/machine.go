package internal

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

const (
	NotAvailableProduct = Error("Not available product")
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

func (m *Machine) ExecuteCommand(command string) {
	separateRegexp := `[\w.-]+`
	r := regexp.MustCompile(separateRegexp)
	allCommands := r.FindAllString(command, -1)

	for _, s := range allCommands {
		_, err := strconv.ParseFloat(s, 32)
		if err != nil {
			action := strings.Split(s, "-")
			if action[0] == "GET" {
				err := m.SellItem(action[1])
				if err != nil {
					m.logger.Printf("Error selling item, %s", err)
				}
			} else if action[0] == "RETURN" {
				m.ReturnCoins()
			} else {
				m.logger.Printf("Action: %s", s)
			}
		} else {
			err = m.InsertCoins(s)
			if err != nil {
				m.logger.Printf("%s, %s", err, s)
			}
		}

	}

}

func (m *Machine) InsertCoins(coins ...string) error {
	for _, c := range coins {
		err := m.cashEngine.InsertCoins(c)
		if err != nil {
			return err
		}
	}

	return nil
}

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

func (m *Machine) LoadItem(item Item, amount int) {
	m.itemStock[item.selector] = amount
}

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
