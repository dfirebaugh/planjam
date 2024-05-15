package imggen

import (
	"fmt"
	"image"
	"image/color"

	"github.com/dfirebaugh/planjam/pkg/plan"
	"github.com/dfirebaugh/tortuga/pkg/imagefb"
	"github.com/dfirebaugh/tortuga/pkg/math/geom"
	"golang.org/x/image/colornames"

	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/proggy"
)

type Display interface {
	// Size returns the current size of the display.
	Size() (x, y int16)

	// SetPizel modifies the internal buffer.
	SetPixel(x, y int16, c color.RGBA)

	// Display sends the buffer (if any) to the screen.
	Display() error
}

const (
	imageWidth    = 400
	imageHeight   = 200
	letterYOffset = 11
	linePadding   = 2
)

func Gen(boardLabel string, b *plan.Board) *image.RGBA {
	fb := imagefb.New(imageWidth, imageHeight)

	printAt(fb, fmt.Sprintf("Plan Jam Report -- Board: %s", boardLabel), 5, letterYOffset, colornames.Greenyellow)

	// printAt(fb, fmt.Sprintf("%d", 12), 0, 0+letterYOffset+linePadding, colornames.Greenyellow)
	// printAt(fb, fmt.Sprintf("%d", 33), 0, (2*letterYOffset)+linePadding, colornames.Greenyellow)

	printBarChart(fb, 0, 30)

	img := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))
	img.Pix = fb.GetFrame()

	return img
}

func printAt(d Display, s string, x int, y int, c color.RGBA) {
	tinyfont.WriteLine(d, &proggy.TinySZ8pt7b, int16(x), int16(y), s, c)
}

func percentageBar(fb Display, percent float64, x float64, y float64, color color.RGBA) {
	background := geom.MakeRect(x, y, 200, 10)
	fill := geom.MakeRect(x, y, 200*percent, 10)

	background.Filled(fb, colornames.Grey)
	fill.Filled(fb, colornames.Green)
}

func printBarChart(fb Display, x int, y int) {
	printAt(fb, fmt.Sprintf("%d", 12), x+15, y+18, colornames.Greenyellow)
	percentageBar(fb, .25, float64(x)+30, float64(y)+10, colornames.Green)
	printAt(fb, fmt.Sprintf("%d", 15), x+15, y+33, colornames.Greenyellow)
	percentageBar(fb, .69, float64(x)+30, float64(y)+25, colornames.Green)
	printAt(fb, fmt.Sprintf("%d", 24), x+15, y+48, colornames.Greenyellow)
	percentageBar(fb, 1, float64(x)+30, float64(y)+40, colornames.Green)
}
