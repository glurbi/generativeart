package util

import (
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func DisplayImage(goimg *image.Image) {
	img = ebiten.NewImageFromImage(*goimg)
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Render an image")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

var img *ebiten.Image

func init() {
}

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(img, nil)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}
