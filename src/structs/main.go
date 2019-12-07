package main

import "fmt"

type person struct {
	name string
	age int
	botKills int
	isProPlayer bool
}

func main() {
	clayman := person{ name: "Alok", age:  28}
	krazzy := person{ name: "KP", age:  24}

	fmt.Println(clayman.canPlayCS())
	fmt.Println(krazzy.canPlayCS())

	fmt.Println("------------------")
	fury := person{name:"Deva", age: 25, botKills: 99, isProPlayer: false}

	// Output -  {Deva 25 99 false}
	validateCSPlayer(fury)
	fmt.Println(fury)

	fmt.Println("------------------")
	magni := &person{name:"Suggu", age: 26, botKills: 97, isProPlayer: false}

	//below function creates a copy of player with updated value
	validateCSPlayerWithReference(magni)
	fmt.Println(magni)
}

//below function creates a copy of player with updated value
func validateCSPlayer(player person) {
	player.isProPlayer = player.botKills > 90
}

func validateCSPlayerWithReference(player *person) {
	player.isProPlayer = player.botKills > 90
}

func (player person) canPlayCS() string {
	if player.age < 25 {
		return "Player "+player.name+" cannot play CS 1.6"
	}

	return "Player "+player.name+" can play CS 1.6"
}
