package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/colornames"
	pixelFont "golang.org/x/image/font"
	"io/ioutil"
	"math"
	"os"
)

func loadTTF(size float64) pixelFont.Face {
	file, err := os.Open("Pacifico.ttf")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	font, err := truetype.Parse(bytes)
	if err != nil {
		panic(err)
	}

	return truetype.NewFace(font, &truetype.Options{
		Size:              size,
		GlyphCacheEntries: 1,
	})
}

func main() {
	pixelgl.Run(run)
}

func run() {
	font := loadTTF(80)

	win := createWindow()

	imd := imdraw.New(nil)
	atlas := text.NewAtlas(font, text.ASCII)

	for !win.Closed() {
		win.Clear(colornames.Aliceblue)
		drawArrow(imd)
		drawWatchFace(imd)
		drawText(win, font, atlas)
		imd.Draw(win)
		win.Update()
	}

}

func drawText(win *pixelgl.Window, font pixelFont.Face, atlas *text.Atlas) {
	radians := float64(0) * (math.Pi / 180)
	x1 := 280 + 220*math.Sin(radians)
	y1 := 300 + 220*math.Cos(radians)
	txt := text.New(pixel.V(x1, y1), atlas)
	txt.Color = colornames.Black
	txt.WriteString("XII")
	txt.Draw(win, pixel.IM.Scaled(txt.Orig, 0.5))

	radians = float64(90) * (math.Pi / 180)
	x1 = 300 + 220*math.Sin(radians)
	y1 = 290 + 220*math.Cos(radians)
	txt = text.New(pixel.V(x1, y1), atlas)
	txt.Color = colornames.Black
	txt.WriteString("III")
	txt.Draw(win, pixel.IM.Scaled(txt.Orig, 0.5))

	radians = float64(180) * (math.Pi / 180)
	x1 = 290 + 220*math.Sin(radians)
	y1 = 280 + 220*math.Cos(radians)
	txt = text.New(pixel.V(x1, y1), atlas)
	txt.Color = colornames.Black
	txt.WriteString("VI")
	txt.Draw(win, pixel.IM.Scaled(txt.Orig, 0.5))

	radians = float64(270) * (math.Pi / 180)
	x1 = 280 + 220*math.Sin(radians)
	y1 = 290 + 220*math.Cos(radians)
	txt = text.New(pixel.V(x1, y1), atlas)
	txt.Color = colornames.Black
	txt.WriteString("IX")
	txt.Draw(win, pixel.IM.Scaled(txt.Orig, 0.5))

}

func drawWatchFace(imd *imdraw.IMDraw) {

	imd.Color = colornames.Blue
	imd.EndShape = imdraw.RoundEndShape
	imd.Push(pixel.V(300, 300))
	imd.Circle(200, 10)
	for i := 0; i < 360; i = i + 90 {
		radians := float64(i) * (math.Pi / 180)
		x1 := 300 + 200*math.Cos(radians)
		y1 := 300 + 200*math.Sin(radians)
		x2 := 300 + 170*math.Cos(radians)
		y2 := 300 + 170*math.Sin(radians)
		imd.Color = colornames.Blue
		imd.Push(pixel.V(x2, y2), pixel.V(x1, y1))
		imd.Line(7)
	}
	for i := 0; i < 360; i = i + 6 {
		radians := float64(i) * (math.Pi / 180)
		x1 := 300 + 200*math.Cos(radians)
		y1 := 300 + 200*math.Sin(radians)
		x2 := 300 + 185*math.Cos(radians)
		y2 := 300 + 185*math.Sin(radians)
		imd.Color = colornames.Blue
		imd.Push(pixel.V(x2, y2), pixel.V(x1, y1))
		imd.Line(3)
	}
}

func drawArrow(imd *imdraw.IMDraw) {

}

func createWindow() *pixelgl.Window {
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 600, 600),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	return win
}
