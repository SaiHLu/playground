package main

import "fmt"

type Position struct {
	x int
	y int
}

type Entity struct {
	id   string
	name string
	Position
}

type SpecialEntity struct {
	specialValue float64
	Entity
}

func main() {
	entity := SpecialEntity{
		specialValue: 10.0,
		Entity: Entity{
			id:   "one",
			name: "sai hlaing",
			Position: Position{
				x: 100,
				y: 200,
			},
		},
	}

	fmt.Printf("The entity is %+v\n", entity)

	fmt.Printf("Position x is %d and y is %d\n", entity.x, entity.y)

	fmt.Println("----------------------")
	fmt.Println("The color is: ", ColorBlack)
}
