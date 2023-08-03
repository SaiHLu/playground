package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Income struct {
	Source string
	Amount int
}

func main() {
	var bankBalance int
	var balance sync.Mutex

	fmt.Println("Initial Balance: ", bankBalance)

	incomes := []Income{
		{Source: "Main Job", Amount: 500},
		{Source: "Partime Job", Amount: 50},
		{Source: "Investment", Amount: 100},
	}

	wg.Add(len(incomes))

	for _, income := range incomes {
		go func(income Income) {
			defer wg.Done()

			for week := 1; week <= 52; week++ {
				balance.Lock()
				bankBalance += income.Amount
				balance.Unlock()

				fmt.Printf("On week %d, you get Income %d from %s\n", week, income.Amount, income.Source)
			}
		}(income)
	}

	wg.Wait()

	fmt.Println("Total Income: ", bankBalance)
}
