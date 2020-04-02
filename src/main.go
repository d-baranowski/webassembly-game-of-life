package main

import (
	"fmt"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/d-baranowski/webassembly-game-of-life/canvas"
	"github.com/d-baranowski/webassembly-game-of-life/src/wasm/life"
	"image/color"
	"net/url"
	"strconv"
	"syscall/js"
	"time"
)

var done chan struct{}

var cvs *canvas.Canvas2d

var lifeState = &(life.Life{})

func main() {
	urlString := js.Global().Get("location").Get("href").String()
	u, err := url.Parse(urlString)
	sizeMod := 0.7

	if err == nil {
		query := u.Query()
		sizeModString := query.Get("sizeMod")
		f, e := strconv.ParseFloat(sizeModString, 32)

		if e == nil {
			sizeMod = f
		}
	}


	cvs, _ = canvas.NewCanvas2d(false)
	cvs.Create(int(js.Global().Get("innerWidth").Float() * sizeMod), int(js.Global().Get("innerHeight").Float() * sizeMod)) // Make Canvas 90% of window size.  For testing rendering canvas smaller than full windows

	lifeState.Initialise(cvs.Width(), cvs.Height(), life.RandomMap(cvs.Width(), cvs.Height()))


	cvs.Start(10, Render)
	<-done
}

func update(gc *draw2dimg.GraphicContext) {
	gc.Current.FillColor = color.Black
	gc.Clear()

	start := time.Now()
	lifeState.Tick()
	elapsed := time.Since(start).Milliseconds()
	println(fmt.Sprintf("Tick Duration %d ms", elapsed))

	start = time.Now()
	generatedImg := lifeState.Draw(cvs.GetBuffer())
	elapsed = time.Since(start).Milliseconds()
	println(fmt.Sprintf("Img generation img %d ms", elapsed))

	start = time.Now()
	cvs.SetImage(generatedImg)
	elapsed = time.Since(start).Milliseconds()
	println(fmt.Sprintf("Draw Duration %d ms", elapsed))
}

// Called from the 'requestAnnimationFrame' function.   It may also be called seperatly from a 'doEvery' function, if the user prefers drawing to be seperate from the annimationFrame callback
func Render(gc *draw2dimg.GraphicContext) bool {
	go update(gc)

	return true
}
