package main

import (
	movement "Game/cmd"
	"fmt"
	"os"
)

var (
	url    = os.Getenv("URL")
	token  = os.Getenv("TOKEN")
	player = os.Getenv("PLAYER")
)

func main() {
	moveEndpoint := fmt.Sprintf("%s/my/%s/action/move", url, player)

	movement.MovePlayer(moveEndpoint, token)
}
