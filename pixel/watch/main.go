package main

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/colornames"
	pixelFont "golang.org/x/image/font"
	"image"
	_ "image/png"
	"io/ioutil"
	"math"
	"os"
	"time"
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

	p, err := loadPicture("gopher.png")
	if err != nil {
		panic(err)
	}
	sprite := pixel.NewSprite(p, p.Bounds())

	for !win.Closed() {
		win.Clear(colornames.Aliceblue)
		drawArrow(imd)
		drawWatchFace(imd)
		drawText(win, font, atlas)
		imd.Draw(win)

		mat := pixel.IM
		mat = drawGopher()
		sprite.Draw(win, mat)

		win.Update()
	}

}

func drawText(win *pixelgl.Window, font pixelFont.Face, atlas *text.Atlas) {
	for i:=0 ; i <= 11 ; i++ {
		radians := float64(i * 30) * (math.Pi / 180)
		x1 := 290 + 230*math.Sin(radians)
		y1 := 290 + 230*math.Cos(radians)
		txt := text.New(pixel.V(x1, y1), atlas)
		txt.Color = colornames.Black
		if i == 0 {
			txt.WriteString("12")
		} else {
			txt.WriteString(fmt.Sprintf("%d", i))
		}
		txt.Draw(win, pixel.IM.Scaled(txt.Orig, 0.5))
	}

}

func drawWatchFace(imd *imdraw.IMDraw) {
	//imd.Color = colornames.Blue
	//imd.EndShape = imdraw.RoundEndShape
	//imd.Push(pixel.V(300, 300))
	//imd.Circle(200, 10)
	for i := 0; i <= 11 ; i++ {
		radians := float64(i*30) * (math.Pi / 180)
		x1 := 300 + 200*math.Cos(radians)
		y1 := 300 + 200*math.Sin(radians)
		x2 := 300 + 170*math.Cos(radians)
		y2 := 300 + 170*math.Sin(radians)
		imd.Color = colornames.Black
		imd.Push(pixel.V(x2, y2), pixel.V(x1, y1))
		imd.Line(7)
	}
	for i := 0; i < 360; i = i + 6 {
		radians := float64(i) * (math.Pi / 180)
		x1 := 300 + 200*math.Cos(radians)
		y1 := 300 + 200*math.Sin(radians)
		x2 := 300 + 185*math.Cos(radians)
		y2 := 300 + 185*math.Sin(radians)
		imd.Color = colornames.Black
		imd.Push(pixel.V(x2, y2), pixel.V(x1, y1))
		imd.Line(3)
	}
}

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func drawGopher () pixel.Matrix {
	t := time.Now()
	s := float64(t.Second())
	radians := s * 6 * (math.Pi / 180)
	x1 := 300 + 150*math.Sin(radians)
	y1 := 300 + 150*math.Cos(radians)
	mat := pixel.IM
	mat = mat.Moved(pixel.V(x1, y1))
	mat = mat.ScaledXY(pixel.V(x1, y1), pixel.V(0.1, 0.1))
	return mat
}

func drawArrow(imd *imdraw.IMDraw) {
	imd.Clear()
	t := time.Now()
	s := float64(t.Second())
	m := float64(t.Minute())
	h := float64(t.Hour())
	if h >= 12 {
		h = h - 12
	}
	radians := s * 6 * (math.Pi / 180)
	x1 := 300 + 180*math.Sin(radians)
	y1 := 300 + 180*math.Cos(radians)
	//imd.Color = colornames.Black
	//imd.Push(pixel.V(300, 300), pixel.V(x1, y1))
	//imd.Line(3)

	radians = (m * 6 + s * 0.1) * (math.Pi / 180)
	x1 = 300 + 160*math.Sin(radians)
	y1 = 300 + 160*math.Cos(radians)
	imd.Color = colornames.Black
	imd.Push(pixel.V(300, 300), pixel.V(x1, y1))
	imd.Line(6)

	radians = (h * 30 + m * 0.5) * (math.Pi / 180)
	x1 = 300 + 130*math.Sin(radians)
	y1 = 300 + 130*math.Cos(radians)
	imd.Color = colornames.Black
	imd.Push(pixel.V(300, 300), pixel.V(x1, y1))
	imd.Line(10)


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
