package main

import (
	"fmt"
	"sync"
)

var philosophers = []string{"Plato", "Socrates", "Aristole", "Pascal", "Locke"}
var wg sync.WaitGroup

func dinningProblem(philosopher string, leftFork, rightFork *sync.Mutex) {
	defer wg.Done()

	fmt.Println(philosopher, " is seated.")

	// lock the forks
	// each philosopher will eat 3 times (just assumption)
	for eatCount := 3; eatCount > 0; eatCount-- {
		fmt.Println(philosopher, " will grap the forks.")

		leftFork.Lock()
		fmt.Printf("\t%s pick up the left fork. \n", philosopher)

		rightFork.Lock()
		fmt.Printf("\t%s pick up the right fork. \n", philosopher)

		fmt.Printf("\t%s has both forks and start eating. ", philosopher)

		leftFork.Unlock()
		fmt.Printf("\t%s put down the fork on his left side. \n", philosopher)
		rightFork.Unlock()
		fmt.Printf("\t%s put down the fork on his right side. \n", philosopher)
	}

	fmt.Println(philosopher, " is satisfied.")
}

func main() {
	wg.Add(len(philosophers))

	leftFork := &sync.Mutex{}

	// spawn goroutine for each philosopher
	for i := 0; i < len(philosophers); i++ {
		rightFork := &sync.Mutex{}

		go dinningProblem(philosophers[i], leftFork, rightFork)

		leftFork = rightFork
	}

	wg.Wait()

	fmt.Println("The table is empty.")
}
