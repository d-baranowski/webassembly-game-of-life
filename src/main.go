package main

import (
	"fmt"
	"github.com/markfarnan/go-canvas/src/wasm/life"
	"time"

	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/markfarnan/go-canvas/canvas"
)

var done chan struct{}

var cvs *canvas.Canvas2d

var renderDelay = 20 * time.Millisecond

var lifeState = life.Initialise()

func main() {
	FrameRate := time.Second / renderDelay
	println("Hello Browser FPS:", fmt.Sprintf("%v\n", FrameRate))

	cvs, _ = canvas.NewCanvas2d(false)
	/*	cvs.Create(int(js.Global().Get("innerWidth").Float()*0.9), int(js.Global().Get("innerHeight").Float()*0.9)) // Make Canvas 90% of window size.  For testing rendering canvas smaller than full windows
	 */

	cvs.Create(400, 400)

	cvs.Start(60, Render)
	<-done
}

// Called from the 'requestAnnimationFrame' function.   It may also be called seperatly from a 'doEvery' function, if the user prefers drawing to be seperate from the annimationFrame callback
func Render(gc *draw2dimg.GraphicContext) bool {
	gc.Clear()

	start := time.Now()
	lifeState = lifeState.Tick()
	elapsed := time.Since(start).Milliseconds()
	println(fmt.Sprintf("Tick Duration %d ms", elapsed))

	start = time.Now()
	gc.DrawImage(lifeState.Draw())
	elapsed = time.Since(start).Milliseconds()
	println(fmt.Sprintf("Draw Duration %d ms", elapsed))

	return true
}
