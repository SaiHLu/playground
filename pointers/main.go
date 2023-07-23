package main

import "fmt"

type Player struct {
	HP int
}

func (p *Player) takeDamage(amount int) int {
	p.HP -= amount

	return p.HP
}

func main() {
	player := &Player{
		HP: 100,
	}

	fmt.Printf("Init Player HP: %d\n", player.HP)
	fmt.Printf("After Player Take Damage: %d\n", player.takeDamage(20))
	fmt.Printf("Finally Player Information: %+v\n", player)

}
