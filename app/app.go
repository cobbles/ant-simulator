package app

import (
	"github.com/cobbles/ant-simulator/ant"
	"math/rand"
	"time"
)

const antTotal = 10
const gridSize = 50
const tickMs = 100

type WorldStruct struct {
	Ants []ant.Ant `json:"ants"`
}

var World = WorldStruct{}
var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

func Start() {
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
