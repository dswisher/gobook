// Mandelbrot emits a PNG image of the Mandelbrot fractal
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"math/rand"
	"os"
)

const (
	iterations = 255 // must be <= 255
)

var palette []color.Color

func main() {
	const (
		xmin, ymin, xmax, ymax = -2.2, -1.8, +1.1, +1.8
		width, height          = 1024, 1024
	)

	buildPalette()

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func buildPalette() {

	s1 := rand.NewSource(316)
	r1 := rand.New(s1)

	palette = make([]color.Color, iterations)

	for i := uint8(2); i < iterations; i++ {
		r := uint8(r1.Intn(255))
		g := uint8(r1.Intn(255))
		b := uint8(r1.Intn(255))
		palette[i] = color.RGBA{r, g, b, 255}
	}

	palette[1] = palette[2]
	palette[0] = palette[1]
}

func mandelbrot(z complex128) color.Color {
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return palette[n]
		}
	}

	return color.Black
}
