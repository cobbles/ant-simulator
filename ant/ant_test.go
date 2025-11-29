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
	tests := []struct {
		name       string
		startX     int
		startY     int
		attempts   int
		shouldMove bool
	}{
		{
			name:       "ant should move from origin",
			startX:     0,
			startY:     0,
			attempts:   100,
			shouldMove: true,
		},
		{
			name:       "ant should move from center",
			startX:     5,
			startY:     5,
			attempts:   50,
			shouldMove: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ant := NewAnt(tt.startX, tt.startY)

			var moved bool
			for i := 0; i < tt.attempts; i++ {
				newX, newY := ant.Move()
				if newX != tt.startX || newY != tt.startY {
					moved = true
					break
				}
			}

			if tt.shouldMove && !moved {
				t.Errorf("Expected ant to move after %d attempts from (%d,%d), but it didn't",
					tt.attempts, tt.startX, tt.startY)
			}
		})
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
