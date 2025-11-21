package ant

import (
	"math/rand"
	"testing"
)

func TestNewAnt(t *testing.T) {
	ant := NewAnt(5, 10)

	if ant.X != 5 {
		t.Errorf("Expected X=5, got %d", ant.X)
	}
	if ant.Y != 10 {
		t.Errorf("Expected Y=10, got %d", ant.Y)
	}
	if ant.r == nil {
		t.Error("Expected random generator to be initialized")
	}
}

func TestAntMove(t *testing.T) {
	// Create ant with deterministic random source for testing
	r := rand.New(rand.NewSource(42))
	ant := &Ant{
		X: 5,
		Y: 5,
		r: r,
	}

	// Test that Move returns valid coordinates
	newX, newY := ant.Move()

	// Movement should be within -1, 0, +1 of original position
	if newX < ant.X-1 || newX > ant.X+1 {
		t.Errorf("X movement out of range: got %d, expected %d to %d", newX, ant.X-1, ant.X+1)
	}
	if newY < ant.Y-1 || newY > ant.Y+1 {
		t.Errorf("Y movement out of range: got %d, expected %d to %d", newY, ant.Y-1, ant.Y+1)
	}
}

func TestAntMoveMultipleTimes(t *testing.T) {
	ant := NewAnt(0, 0)

	// Test that ant can move multiple times
	moved := false
	for i := 0; i < 100; i++ {
		newX, newY := ant.Move()
		if newX != 0 || newY != 0 {
			moved = true
			break
		}
	}

	if !moved {
		t.Error("Ant never moved from starting position after 100 attempts")
	}
}

func TestAntMoveDeterministic(t *testing.T) {
	// Test that with same seed, we get same movement
	r1 := rand.New(rand.NewSource(123))
	ant1 := &Ant{X: 10, Y: 10, r: r1}

	r2 := rand.New(rand.NewSource(123))
	ant2 := &Ant{X: 10, Y: 10, r: r2}

	newX1, newY1 := ant1.Move()
	newX2, newY2 := ant2.Move()

	if newX1 != newX2 || newY1 != newY2 {
		t.Errorf("Expected same movement with same seed: got (%d,%d) and (%d,%d)",
			newX1, newY1, newX2, newY2)
	}
}
