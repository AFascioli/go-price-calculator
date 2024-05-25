package cmdmanager

import "fmt"

type CMDManager struct {
}

func New() *CMDManager {
	return &CMDManager{}
}

func (cmd CMDManager) ReadLines() ([]string, error) {
	fmt.Println("enter your prices, confirm each one pressing ENTER")

	var prices []string
	for {
		var price string
		fmt.Print("price: ")
		fmt.Scan(&price)
		if price == "0" {
			break
		}
		prices = append(prices, price)
	}
	return prices, nil
}

func (cmd CMDManager) WriteResult(data interface{}) error {
	fmt.Println(data)
	return nil
}
