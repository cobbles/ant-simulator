package app

import (
	"math"
	"math/rand"
	"time"

	"github.com/cobbles/ant-simulator/ant"
	"github.com/cobbles/ant-simulator/food"
)

const antTotal = 5
const gridSize = 75
const tickMs = 100

type TNest struct {
	X         int `json:"x"`
	Y         int `json:"y"`
	FoodCount int `json:"foodCount"`
}

type TWorld struct {
	Ants  []ant.Ant    `json:"ants"`
	Nest  TNest        `json:"nest"`
	Foods []food.TFood `json:"foods"`
}

var World = TWorld{}

func Start() {
	World.Nest = TNest{
		X:         rand.Intn(gridSize),
		Y:         rand.Intn(gridSize),
		FoodCount: 3,
	}

	World.Foods = append(World.Foods, food.TFood{
		X: rand.Intn(gridSize),
		Y: rand.Intn(gridSize),
	})

	for range antTotal {
		World.Ants = append(
			World.Ants,
			ant.NewAnt(
				rand.Intn(gridSize),
				rand.Intn(gridSize),
			),
		)
	}

	go simulationLoop()
}

func simulationLoop() {
	food := &World.Foods[0]
	for {
		for i := range World.Ants {
			ant := &World.Ants[i]
			if nearFood(ant, food) {
				MoveAntToFood(ant, food)
			} else {
				MoveAnt(ant)
			}
		}
		time.Sleep(tickMs * time.Millisecond)
	}
}

func nearFood(ant *ant.Ant, food *food.TFood) bool {
	return int(math.Abs(float64(ant.X-food.X))+math.Abs(float64(ant.Y-food.Y))) <= 10
}

func MoveAntToFood(ant *ant.Ant, food *food.TFood) {
	nextX := ant.X
	nextY := ant.Y

	if ant.X < food.X {
		nextX++
	} else if ant.X > food.X {
		nextX--
	}

	if ant.Y < food.Y {
		nextY++
	} else if ant.Y > food.Y {
		nextY--
	}

	ant.X, ant.Y = applyBounce(nextX, nextY)
}

func applyBounce(nextX, nextY int) (int, int) {
	// Bounce off walls - if next position is outside bounds, move in opposite direction
	if nextX < 0 {
		nextX = -nextX // Bounce right
	} else if nextX >= gridSize {
		nextX = 2*gridSize - 2 - nextX // Bounce left
	}

	if nextY < 0 {
		nextY = -nextY // Bounce down
	} else if nextY >= gridSize {
		nextY = 2*gridSize - 2 - nextY // Bounce up
	}

	return nextX, nextY
}

func MoveAnt(ant *ant.Ant) {
	nextX, nextY := ant.Move()
	ant.X, ant.Y = applyBounce(nextX, nextY)
}
