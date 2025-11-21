package ant

import (
	"math/rand"
	"time"
)

type Ant struct {
	X int `json:"x"`
	Y int `json:"y"`
	r *rand.Rand
}

func NewAnt(x, y int) Ant {
	return Ant{
		X: x,
		Y: y,
		r: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (a *Ant) Move() (int, int) {
	dx := a.r.Intn(3) - 1 // -1, 0, or 1
	dy := a.r.Intn(3) - 1

	return a.X + dx, a.Y + dy
}
