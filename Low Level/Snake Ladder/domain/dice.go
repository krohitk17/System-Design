package domain

import "math/rand"

type Dice struct {
	Count int
	Rolls int
}

func CreateDice(n int) *Dice {
	return &Dice{
		Count: n,
		Rolls: 0,
	}
}

func (this *Dice) Roll() int {
	return rand.Intn(6*this.Count-this.Count) + this.Count
}
