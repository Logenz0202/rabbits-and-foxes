package graphics

import (
	"embed"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed assets/tiles/*.jpg assets/animals/*.png
var tileFiles embed.FS

var (
	GrassTile *ebiten.Image
	DirtTile  *ebiten.Image
	Rabbit    *ebiten.Image
)

func LoadAssets() {
	load := func(path string) *ebiten.Image {
		file, err := tileFiles.Open(path)
		if err != nil {
			log.Fatalf("cannot open asset: %v", err)
		}
		img, _, err := image.Decode(file)
		if err != nil {
			log.Fatalf("cannot decode asset: %v", err)
		}
		return ebiten.NewImageFromImage(img)
	}

	GrassTile = load("assets/tiles/grass.jpg")
	DirtTile = load("assets/tiles/dirt.jpg")
	Rabbit = load("assets/animals/rabbit.png")
}
