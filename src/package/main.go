package main

import (
	"fmt"
	"workspace/grpc-java/demo/src/package/model"
)

func main() {
	playerList := model.GetPlayerList()
	for _, player := range playerList {
		fmt.Println(player.Playcs())
	}

}
