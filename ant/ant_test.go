package ant

import (
	"math/rand"
	"testing"

	"github.com/cobbles/ant-simulator/food"
)

func TestMoveAntWithinBounds(t *testing.T) {
	r := rand.New(rand.NewSource(1))

	ant := &Ant{X: 0, Y: 0}
	gridSize := 10
	food := []food.TFood{{X: 1, Y: 1}}
	MoveAnt(r, ant, gridSize, food)

	if ant.X < 0 || ant.X >= gridSize {
		t.Errorf("X out of bounds: got %d", ant.X)
	}
	if ant.Y < 0 || ant.Y >= gridSize {
		t.Errorf("Y out of bounds: got %d", ant.Y)
	}
}

func TestMoveAntMoves(t *testing.T) {
	r := rand.New(rand.NewSource(42))

	ant := &Ant{X: 5, Y: 5}
	before := *ant
	gridSize := 10
	food := []food.TFood{{X: 1, Y: 1}}
	MoveAnt(r, ant, gridSize, food)

	if ant.X == before.X && ant.Y == before.Y {
		t.Errorf("ant did not move: %+v", ant)
	}
}

func TestMoveAntMovesTowardsFood(t *testing.T) {
	r := rand.New(rand.NewSource(42))

	ant := &Ant{X: 5, Y: 5}
	before := *ant
	gridSize := 10
	food := []food.TFood{{X: 1, Y: 1}}
	MoveAnt(r, ant, gridSize, food)

	if ant.X == before.X && ant.Y == before.Y {
		t.Errorf("ant did not move: %+v", ant)
	}
}
