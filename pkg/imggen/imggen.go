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
	Size() (x, y int16)
	SetPixel(x, y int16, c color.RGBA)
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

	printBarChart(fb, 0, 30, b)

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
	fill.Filled(fb, color)
}

func printBarChart(fb Display, x int, y int, b *plan.Board) {
	totalFeatures := countTotalFeatures(b)

	for i, lane := range b.Lanes {
		yOffset := y + i*30

		laneFeatureCount := len(lane.Features)
		percentOfTotal := float64(laneFeatureCount) / float64(totalFeatures)

		printAt(fb, fmt.Sprintf("%s: %d features", lane.Label, laneFeatureCount), x+10, yOffset+18, colornames.Greenyellow)
		percentageBar(fb, percentOfTotal, float64(x)+20, float64(yOffset)+25, colornames.Green)
	}
}

func countTotalFeatures(b *plan.Board) int {
	total := 0
	for _, lane := range b.Lanes {
		total += len(lane.Features)
	}
	return total
}
