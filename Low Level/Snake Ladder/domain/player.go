package domain

type Player struct {
	Name  string
	Wins  int
	Games int
}

func CreatePlayer(name string) *Player {
	return &Player{
		Name:  name,
		Wins:  0,
		Games: 0,
	}
}

func (this *Player) AddToGame() {
	this.Games++
}

func (this *Player) WonGame() {
	this.Wins++
}
