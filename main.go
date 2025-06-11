package main

import (
	"math/rand"
	"rabbits-and-foxes/internal/game"
	"rabbits-and-foxes/internal/graphics"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	graphics.LoadAssets()
	g := game.NewGame()

	ebiten.SetWindowTitle("Rabbits and Foxes – Ecosystem")
	ebiten.SetWindowSize(956, 800)
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
