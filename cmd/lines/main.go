package main

import (
	_ "image/png"

	"github.com/fogleman/gg"
	util "github.com/glurbi/generativeart/internal"
)

func main() {
	dc := gg.NewContext(1000, 1000)
	dc.DrawCircle(500, 500, 400)
	dc.SetRGB(0, 1, 0)
	dc.Fill()
	img := dc.Image()
	util.DisplayImage(&img)
}
