package main

import (
	"fmt"
	"github.com/fogleman/gg"
)

func main() {
	const W = 1000
	const H = 1000
	dc := gg.NewContext(W, H)
	dc.SetRGB(0, 0, 0)
	im, err := gg.LoadImage("gopher.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	dc.DrawImage(im, 500, 500)

	dc.DrawRectangle(float64(W - 200), float64(H/2), float64(20), float64(200))
	dc.SetRGB255(100,65,23)
	dc.FillPreserve()
	dc.Stroke()

	dc.DrawCircle(float64(W - 200), float64(H/2), float64(50))
	dc.SetRGB255(50,205,50)
	dc.FillPreserve()
	dc.Stroke()

	//dc.Clear()
	dc.Stroke()
	if err := dc.LoadFontFace("AmaticSC-Regular.ttf", 96); err != nil {
		panic(err)
	}
	dc.DrawStringAnchored("I'm a gopher!!!", W/2 - 200, W/2, 0.5, 0.5)
	dc.SavePNG("out.png")

}
