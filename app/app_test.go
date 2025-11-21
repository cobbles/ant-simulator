package app

import (
	"testing"

	"github.com/cobbles/ant-simulator/ant"
	"github.com/cobbles/ant-simulator/food"
)

func TestMoveAntBounceOffLeftWall(t *testing.T) {
	// Create ant at position (0, 10) - at left edge
	testAnt := ant.NewAnt(0, 10)

	// Try to move ant many times - eventually it should try to move left
	// and bounce back instead of going to negative coordinates
	for i := 0; i < 100; i++ {
		MoveAnt(&testAnt)

		// Ant should never have negative X coordinate
		if testAnt.X < 0 {
			t.Errorf("Ant should never have negative X coordinate, got X=%d", testAnt.X)
			break
		}

		// If ant moved away from edge, reset it back to edge for continued testing
		if testAnt.X > 0 {
			testAnt.X = 0
		}
	}

	// After many attempts, ant should still be at valid position
	if testAnt.X < 0 || testAnt.X >= gridSize {
		t.Errorf("Ant should be within grid bounds [0, %d], got X=%d", gridSize-1, testAnt.X)
	}
}

func TestMoveAntBounceOffRightWall(t *testing.T) {
	// Create ant at position (gridSize-1, 10) - at right edge
	testAnt := ant.NewAnt(gridSize-1, 10)

	// Try to move ant many times - eventually it should try to move right
	// and bounce back instead of going beyond gridSize
	for i := 0; i < 100; i++ {
		MoveAnt(&testAnt)

		// Ant should never go beyond gridSize-1
		if testAnt.X >= gridSize {
			t.Errorf("Ant should never go beyond gridSize-1=%d, got X=%d", gridSize-1, testAnt.X)
			break
		}

		// If ant moved away from edge, reset it back to edge for continued testing
		if testAnt.X < gridSize-1 {
			testAnt.X = gridSize - 1
		}
	}

	// After many attempts, ant should still be at valid position
	if testAnt.X < 0 || testAnt.X >= gridSize {
		t.Errorf("Ant should be within grid bounds [0, %d], got X=%d", gridSize-1, testAnt.X)
	}
}

func TestMoveAntBounceOffTopWall(t *testing.T) {
	// Create ant at position (10, 0) - at top edge
	testAnt := ant.NewAnt(10, 0)

	// Try to move ant many times - eventually it should try to move up
	// and bounce back instead of going to negative coordinates
	for i := 0; i < 100; i++ {
		MoveAnt(&testAnt)

		// Ant should never have negative Y coordinate
		if testAnt.Y < 0 {
			t.Errorf("Ant should never have negative Y coordinate, got Y=%d", testAnt.Y)
			break
		}

		// If ant moved away from edge, reset it back to edge for continued testing
		if testAnt.Y > 0 {
			testAnt.Y = 0
		}
	}

	// After many attempts, ant should still be at valid position
	if testAnt.Y < 0 || testAnt.Y >= gridSize {
		t.Errorf("Ant should be within grid bounds [0, %d], got Y=%d", gridSize-1, testAnt.Y)
	}
}

func TestMoveAntBounceOffBottomWall(t *testing.T) {
	// Create ant at position (10, gridSize-1) - at bottom edge
	testAnt := ant.NewAnt(10, gridSize-1)

	// Try to move ant many times - eventually it should try to move down
	// and bounce back instead of going beyond gridSize
	for i := 0; i < 100; i++ {
		MoveAnt(&testAnt)

		// Ant should never go beyond gridSize-1
		if testAnt.Y >= gridSize {
			t.Errorf("Ant should never go beyond gridSize-1=%d, got Y=%d", gridSize-1, testAnt.Y)
			break
		}

		// If ant moved away from edge, reset it back to edge for continued testing
		if testAnt.Y < gridSize-1 {
			testAnt.Y = gridSize - 1
		}
	}

	// After many attempts, ant should still be at valid position
	if testAnt.Y < 0 || testAnt.Y >= gridSize {
		t.Errorf("Ant should be within grid bounds [0, %d], got Y=%d", gridSize-1, testAnt.Y)
	}
}

func TestMoveAntToFoodBounceOffLeftWall(t *testing.T) {
	// Create ant at position (0, 10) and food to the left (outside grid)
	testAnt := ant.NewAnt(0, 10)
	testFood := food.TFood{X: -5, Y: 10}

	MoveAntToFood(&testAnt, &testFood)

	// Should bounce back to X=1, not go to -1
	if testAnt.X != 1 {
		t.Errorf("Expected ant to bounce to X=1, got X=%d", testAnt.X)
	}
}

func TestMoveAntToFoodBounceOffRightWall(t *testing.T) {
	// Create ant at position (gridSize-1, 10) and food to the right (outside grid)
	testAnt := ant.NewAnt(gridSize-1, 10)
	testFood := food.TFood{X: gridSize + 5, Y: 10}

	MoveAntToFood(&testAnt, &testFood)

	// Should bounce back to X=gridSize-2, not go to gridSize
	if testAnt.X != gridSize-2 {
		t.Errorf("Expected ant to bounce to X=%d, got X=%d", gridSize-2, testAnt.X)
	}
}

func TestMoveAntToFoodBounceOffTopWall(t *testing.T) {
	// Create ant at position (10, 0) and food above (outside grid)
	testAnt := ant.NewAnt(10, 0)
	testFood := food.TFood{X: 10, Y: -5}

	MoveAntToFood(&testAnt, &testFood)

	// Should bounce back to Y=1, not go to -1
	if testAnt.Y != 1 {
		t.Errorf("Expected ant to bounce to Y=1, got Y=%d", testAnt.Y)
	}
}

func TestMoveAntToFoodBounceOffBottomWall(t *testing.T) {
	// Create ant at position (10, gridSize-1) and food below (outside grid)
	testAnt := ant.NewAnt(10, gridSize-1)
	testFood := food.TFood{X: 10, Y: gridSize + 5}

	MoveAntToFood(&testAnt, &testFood)

	// Should bounce back to Y=gridSize-2, not go to gridSize
	if testAnt.Y != gridSize-2 {
		t.Errorf("Expected ant to bounce to Y=%d, got Y=%d", gridSize-2, testAnt.Y)
	}
}
