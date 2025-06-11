package world

import "math/rand"

const (
	FoxInitialEnergy     = 15.0
	FoxMaxEnergy         = 25.0
	FoxEatGain           = 10.0
	FoxReproductionCD    = 3
	FoxEnergyLossPerTick = 0.3
)

type Fox struct {
	X, Y                 int
	Energy               float64
	ReproductionCooldown int
	Direction            int // -1 for left, 1 for right
}

func NewFox(x, y int) *Fox {
	return &Fox{
		X:                    x,
		Y:                    y,
		Energy:               FoxInitialEnergy,
		ReproductionCooldown: FoxReproductionCD,
		Direction:            -1,
	}
}

func (f *Fox) IsAlive() bool {
	return f.Energy > 0
}

func (f *Fox) CanReproduce() bool {
	return f.ReproductionCooldown <= 0 && f.Energy > FoxInitialEnergy
}

func (f *Fox) Move(w *World) {
	dx := rand.Intn(3) - 1 // -1, 0, 1
	dy := rand.Intn(3) - 1
	nx := Clamp(f.X+dx, 0, MapWidth-1)
	ny := Clamp(f.Y+dy, 0, MapHeight-1)

	if dx > 0 {
		f.Direction = 1
	} else if dx < 0 {
		f.Direction = -1
	}

	if rabbit := w.IsOccupiedByRabbit(nx, ny); rabbit != nil {
		f.Energy += FoxEatGain
		if f.Energy > FoxMaxEnergy {
			f.Energy = FoxMaxEnergy
		}
		rabbit.Energy = 0
	}

	if other := w.IsOccupiedByFox(nx, ny); other == nil {
		f.X, f.Y = nx, ny
	}
}
