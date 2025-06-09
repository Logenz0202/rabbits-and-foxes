package world

import (
	"math/rand"
)

const (
	MaxGrassAmount = 1.0
	GrowthRate     = 0.01
	SpreadChance   = 0.01
)

type Grass struct {
	Amount float64 // 0 to MaxGrassAmount
}

func NewGrass() *Grass {
	return &Grass{Amount: 0.2 + rand.Float64()*0.3}
}

func (g *Grass) Update(w *World, x, y int) {
	if g.Amount < MaxGrassAmount {
		g.Amount += GrowthRate
		if g.Amount > MaxGrassAmount {
			g.Amount = MaxGrassAmount
		}
	}

	if rand.Float64() < SpreadChance {
		dx := rand.Intn(3) - 1
		dy := rand.Intn(3) - 1
		if dx != 0 || dy != 0 {
			nx, ny := x+dx, y+dy
			if nx >= 0 && ny >= 0 && nx < MapWidth && ny < MapHeight {
				neighbor := &w.Tiles[ny][nx]
				if neighbor.Grass == nil {
					neighbor.Grass = NewGrass()
				}
			}
		}
	}
}
