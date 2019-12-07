package model

type person struct {
	name string
	age int
	botKills int
	isProPlayer bool
}
type teeneger struct {
	name string
	age int
	country string
}

//Capitalize interface method and the struct method to make it accessible from outside package

func (player person) Playcs() string {
	if player.age < 25 {
		return "Player "+player.name+" cannot play CS 1.6"
	}
	return "Player "+player.name+" can play CS 1.6"
}

func (player teeneger) Playcs() string {
	if player.age < 15 {
		return "Player "+player.name+" cannot play CS 1.6"
	}
	return "Player "+player.name+" can play CS 1.6"
}

type playcser interface {
	Playcs() string
}

//Method capitalized to make it accessible from outside package
func GetPlayerList() []playcser {

	fury := &person{name:"Deva", age: 25, botKills: 99, isProPlayer: false}
	clayman := &person{name:"Alok", age: 28, botKills: 999, isProPlayer: false}
	neo := &teeneger{name:"Politski", age: 13, }

	playerList := []playcser{fury, clayman, neo}

	return playerList;
}