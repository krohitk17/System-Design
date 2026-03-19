package domain

import (
	"errors"
	"fmt"
)

type Game struct {
	Size int
	Dice *Dice

	Players   []*Player
	Remaining int
	Position  []int
	Turn      int

	Snakes  map[int]int
	Ladders map[int]int
}

func CreateGame(size int, dice *Dice, players []*Player) *Game {
	for _, p := range players {
		p.AddToGame()
	}
	return &Game{
		Size:      size * size,
		Dice:      dice,
		Players:   players,
		Remaining: len(players),
		Position:  make([]int, len(players)),
	}
}

func (this *Game) validateJumpInput(start, end int) error {
	if start <= 0 {
		return errors.New("Start must be greater than 0")
	}
	if start > this.Size {
		return errors.New("Start must be less than board size")
	}

	if end <= 0 {
		return errors.New("End must be greater than 0")
	}
	if end > this.Size {
		return errors.New("End must be less than board size")
	}
	return nil
}

func (this *Game) AddSnake(start, end int) error {
	if err := this.validateJumpInput(start, end); err != nil {
		return err
	}
	if end >= start {
		return errors.New("Snake must be descending")
	}
	if _, ok := this.Snakes[start]; ok {
		return errors.New(fmt.Sprintf("Snake already created at %d", start))
	}
	if _, ok := this.Ladders[start]; ok {
		return errors.New(fmt.Sprintf("Ladder found at %d", start))
	}

	this.Snakes[start] = end
	return nil
}

func (this *Game) AddLadder(start, end int) error {
	if err := this.validateJumpInput(start, end); err != nil {
		return err
	}
	if end <= start {
		return errors.New("Ladder must be ascending")
	}
	if _, ok := this.Ladders[start]; ok {
		return errors.New(fmt.Sprintf("Ladder already created at %d", start))
	}
	if _, ok := this.Snakes[start]; ok {
		return errors.New(fmt.Sprintf("Snake found at %d", start))
	}
	this.Ladders[start] = end
	return nil
}

func (this *Game) advanceTurn() {
	this.Turn = (this.Turn + 1) % this.Remaining
}

func (this *Game) ReduceRemaining() {
	this.Position = append(this.Position[:this.Turn], this.Position[this.Turn+1:]...)
	this.Players = append(this.Players[:this.Turn], this.Players[this.Turn+1:]...)
	this.Remaining--
	fmt.Println(fmt.Sprintf("%d Players remaining", this.Remaining))
}

func (this *Game) Play() {
	if this.IsFinished() {
		fmt.Println("This game has ended. Please create a new game.")
		return
	}

	player := this.Players[this.Turn]
	current := this.Position[this.Turn]

	fmt.Println("Current player: " + player.Name)
	fmt.Println(fmt.Sprintf("Player %s is at %d", player.Name, current))

	roll := this.Dice.Roll()
	fmt.Println(fmt.Sprintf("Player %s rolled %d", player.Name, roll))

	if roll+current == this.Size-1 {
		fmt.Println("Player " + player.Name + " reached the finish!")
		player.WonGame()
		this.ReduceRemaining()
	} else if roll+current > this.Size-1 {
		fmt.Println(fmt.Sprintf("Oops! You need to roll %d to finish", this.Size-current-1))
		fmt.Println("Turn skipped")
	} else {
		next := current + roll

		if end, ok := this.Snakes[next]; ok {
			fmt.Println("Oops, you stepped on a snake!")
			next = end
		} else if end, ok := this.Ladders[next]; ok {
			fmt.Println("You found a ladder!")
			next = end
		}

		fmt.Println(fmt.Sprintf("Player %s moved to %d", player.Name, next))
		this.Position[this.Turn] = next
	}

	if this.IsFinished() {
		fmt.Println("Game over!")
	} else {
		this.advanceTurn()
	}
}

func (this *Game) IsFinished() bool {
	return this.Remaining == 0
}
