package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"sync"
)

//(256*256*256/2)/512
const LENGTH = 4096

type RgbOrder func(i, j, k int) (r, g, b int)

type PngBuilder struct {
	i, j, k  int
	order    RgbOrder
	filename string
}

func NewPngBuild(filename string, order RgbOrder) *PngBuilder {
	return &PngBuilder{0, 0, 0, order, filename}
}

func (pb *PngBuilder) NextColor() color.Color {
	if pb.k == 255 {
		if pb.j == 255 {
			if pb.i == 255 {
				pb.i = 0
				pb.j = 0
				pb.k = 0
			} else {
				pb.i++
				pb.j = 0
				pb.k = 0
			}
		} else {
			pb.j++
			pb.k = 0
		}
	} else {
		pb.k++
	}
	return pb.getColor(pb.order(pb.i, pb.j, pb.k))
}

func (pb *PngBuilder) getColor(i, j, k int) color.Color {
	return &color.RGBA{R: uint8(i), G: uint8(j), B: uint8(k), A: uint8(255)}
}

func (pb *PngBuilder) GenImage(wg *sync.WaitGroup) {
	defer wg.Done()
	img := image.NewNRGBA(image.Rect(0, 0, LENGTH, LENGTH))

	for i := 0; i < LENGTH; i++ {
		for j := 0; j < LENGTH; j++ {
			img.Set(i, j, pb.NextColor())
		}
	}
	f, err := os.Create(pb.filename)
	if err != nil {
		println("generate picture failure:", pb.filename, "error:", err.Error())
		os.Exit(1)
	}
	err = png.Encode(f, img)
	if err != nil {
		println("generate picture failure:", pb.filename, "error:", err.Error())
		os.Exit(1)
	}
	println("generate picture success:", pb.filename)
}

var rgb = func(i, j, k int) (int, int, int) { return i, j, k }
var rbg = func(i, j, k int) (int, int, int) { return i, k, j }
var grb = func(i, j, k int) (int, int, int) { return j, i, k }
var gbr = func(i, j, k int) (int, int, int) { return j, k, i }
var bgr = func(i, j, k int) (int, int, int) { return k, j, i }
var brg = func(i, j, k int) (int, int, int) { return k, i, j }

func main() {
	var wg sync.WaitGroup
	wg.Add(6)
	go NewPngBuild("rgb.png", rgb).GenImage(&wg)
	go NewPngBuild("rbg.png", rbg).GenImage(&wg)
	go NewPngBuild("grb.png", grb).GenImage(&wg)
	go NewPngBuild("gbr.png", gbr).GenImage(&wg)
	go NewPngBuild("bgr.png", bgr).GenImage(&wg)
	go NewPngBuild("brg.png", brg).GenImage(&wg)
	wg.Wait()
}
