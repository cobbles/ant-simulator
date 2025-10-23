package app

import (
	"math/rand"
	"time"

	"github.com/cobbles/ant-simulator/ant"
)

const antTotal = 10
const gridSize = 75
const tickMs = 100

type TNest struct {
	X         int `json:"x"`
	Y         int `json:"y"`
	FoodCount int `json:"foodCount"`
}

type TFood struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type TWorld struct {
	Ants  []ant.Ant `json:"ants"`
	Nest  TNest     `json:"nest"`
	Foods []TFood   `json:"foods"`
}

var World = TWorld{}
var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

func Start() {
	World.Nest = TNest{
		X:         rand.Intn(gridSize),
		Y:         rand.Intn(gridSize),
		FoodCount: 3,
	}

	World.Foods = append(World.Foods, TFood{
		X: rand.Intn(gridSize),
		Y: rand.Intn(gridSize),
	})

	for range antTotal {
		World.Ants = append(World.Ants, ant.Ant{
			X: rand.Intn(gridSize),
			Y: rand.Intn(gridSize),
		})
	}

	go simulationLoop()
}

func simulationLoop() {
	for {
		for i := range World.Ants {
			ant.MoveAnt(rng, &World.Ants[i], gridSize)
		}
		time.Sleep(tickMs * time.Millisecond)
	}
}
