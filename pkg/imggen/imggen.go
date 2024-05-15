package imggen

import (
	"fmt"
	"image"
	"image/color"
	"math"
	"math/rand"

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
	printPieChart(fb, 350, 100, 50, b)

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

func printPieChart(fb Display, centerX int, centerY int, radius int, b *plan.Board) {
	totalFeatures := countTotalFeatures(b)
	if totalFeatures == 0 {
		return
	}

	startAngle := 0.0
	usedColors := make(map[color.RGBA]bool)

	for _, lane := range b.Lanes {
		laneFeatureCount := len(lane.Features)
		angle := 2 * math.Pi * float64(laneFeatureCount) / float64(totalFeatures)
		color := randomUniqueColor(usedColors)
		drawPieSlice(fb, centerX, centerY, radius, startAngle, startAngle+angle, color)
		startAngle += angle
	}
}

func drawPieSlice(fb Display, centerX int, centerY int, radius int, startAngle float64, endAngle float64, clr color.RGBA) {
	for angle := startAngle; angle < endAngle; angle += 0.01 {
		for r := 0; r < radius; r++ {
			x := centerX + int(float64(r)*math.Cos(angle))
			y := centerY + int(float64(r)*math.Sin(angle))
			fb.SetPixel(int16(x), int16(y), clr)
		}
	}
}

func randomUniqueColor(usedColors map[color.RGBA]bool) color.RGBA {
	colors := []color.RGBA{
		colornames.Red,
		colornames.Green,
		colornames.Blue,
		colornames.Yellow,
		colornames.Orange,
		colornames.Purple,
		colornames.Cyan,
		colornames.Magenta,
	}
	for {
		color := colors[rand.Intn(len(colors))]
		if !usedColors[color] {
			usedColors[color] = true
			return color
		}
	}
}
