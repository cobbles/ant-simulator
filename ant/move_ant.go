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
	// if food within 5 squares move towards it
    // for i := range food {
    //     dist := abs(ant.X-food[i].X) + abs(ant.Y-food[i].Y)
    //     if dist <= 5 && dist < minDist {
    //         minDist = dist
    //         target = &food[i]
    //     }
    // }
	moveDir := rand.New(rand.New(rand.NewSource(time.Now().UnixNano())))
	if (moveDir.Intn(2) == 1) {
		dx := r.Intn(3) - 1
		ant.X = max(0, min(ant.X + dx, gridSize - 1))
	} else {
		dy := r.Intn(3) - 1
		ant.Y = max(0, min(ant.Y + dy, gridSize - 1))
	}
}

