package main

import "snake-ladder/domain"

func main() {
	players := []*domain.Player{
		domain.CreatePlayer("Rohit"),
		domain.CreatePlayer("Shreyas"),
		domain.CreatePlayer("Taran"),
		domain.CreatePlayer("Rupanshi"),
	}

	dice := domain.CreateDice(1)
	game := domain.CreateGame(10, dice, players)
	for !game.IsFinished() {
		game.Play()
	}
}
