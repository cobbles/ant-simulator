package ant

import "math/rand"

// RandomGenerator defines the interface for random number generation
type RandomGenerator interface {
	Intn(n int) int
}

// RealRandomGenerator wraps the standard rand.Rand for production use
type RealRandomGenerator struct {
	r *rand.Rand
}

// NewRealRandomGenerator creates a new RealRandomGenerator with a random seed
func NewRealRandomGenerator() *RealRandomGenerator {
	return &RealRandomGenerator{
		r: rand.New(rand.NewSource(rand.Int63())),
	}
}

// NewRealRandomGeneratorWithSeed creates a new RealRandomGenerator with a specific seed
func NewRealRandomGeneratorWithSeed(seed int64) *RealRandomGenerator {
	return &RealRandomGenerator{
		r: rand.New(rand.NewSource(seed)),
	}
}

// Intn returns a random integer in [0, n)
func (r *RealRandomGenerator) Intn(n int) int {
	return r.r.Intn(n)
}
