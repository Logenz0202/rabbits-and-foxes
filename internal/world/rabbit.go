package world

import "math/rand"

const (
	RabbitInitialEnergy     = 10.0
	RabbitMaxEnergy         = 20.0
	RabbitEatGain           = 5.0
	RabbitReproductionCD    = 5
	RabbitEnergyLossPerTick = 0.2
)

type Rabbit struct {
	X, Y                 int
	Energy               float64
	ReproductionCooldown int
}

func NewRabbit(x, y int) *Rabbit {
	return &Rabbit{
		X:                    x,
		Y:                    y,
		Energy:               RabbitInitialEnergy,
		ReproductionCooldown: 10,
	}
}

func (r *Rabbit) IsAlive() bool {
	return r.Energy > 0
}

func (r *Rabbit) CanReproduce() bool {
	return r.ReproductionCooldown <= 0 && r.Energy > RabbitInitialEnergy
}

func (r *Rabbit) Move(w *World) {
	dx := rand.Intn(3) - 1 // -1, 0, 1
	dy := rand.Intn(3) - 1
	nx := Clamp(r.X+dx, 0, MapWidth-1)
	ny := Clamp(r.Y+dy, 0, MapHeight-1)
	r.X, r.Y = nx, ny
}

func (r *Rabbit) Eat(w *World) {
	tile := &w.Tiles[r.Y][r.X]
	if tile.Grass != nil && tile.Grass.Amount > 0.5 {
		r.Energy += RabbitEatGain
		tile.Grass.Amount -= RabbitEatGain
		if tile.Grass.Amount < 0 {
			tile.Grass.Amount = 0
		}
		if r.Energy > RabbitMaxEnergy {
			r.Energy = RabbitMaxEnergy
		}
	}
}
