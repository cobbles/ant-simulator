package ant

import (
	"math/rand"
	"time"
)

type Ant struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func MoveAnt(r *rand.Rand, ant *Ant, gridSize int) {
	moveDir := rand.New(rand.New(rand.NewSource(time.Now().UnixNano())))
	if (moveDir.Intn(2) == 1) {
		dx := r.Intn(3) - 1
		ant.X = max(0, min(ant.X + dx, gridSize - 1))
	} else {
		dy := r.Intn(3) - 1
		ant.Y = max(0, min(ant.Y + dy, gridSize - 1))
	}
}

