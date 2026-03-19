package domain

type Jump struct {
	Start int
	End   int
}

func CreateJump(start, end int) *Jump {
	return &Jump{
		Start: start,
		End:   end,
	}
}
