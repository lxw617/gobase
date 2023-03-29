package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"io/ioutil"
	"os"

	"base/tools/writepic/file"

	"github.com/golang/freetype"
)

// 有点问题 --TODO
type DrawText struct {
	JPG    draw.Image
	Merged *os.File

	Title string
	X0    int
	Y0    int
	Size0 float64

	SubTitle string
	X1       int
	Y1       int
	Size1    float64
}

func DrawPoster(d *DrawText, fontName string) error {
	fontSource := "bg.jpg"
	fontSourceBytes, err := ioutil.ReadFile(fontSource)
	if err != nil {
		return err
	}

	trueTypeFont, err := freetype.ParseFont(fontSourceBytes)
	if err != nil {
		return err
	}

	fc := freetype.NewContext()
	fc.SetDPI(72)
	fc.SetFont(trueTypeFont)
	fc.SetFontSize(d.Size0)
	fc.SetClip(d.JPG.Bounds())
	fc.SetDst(d.JPG)
	fc.SetSrc(image.Black)

	pt := freetype.Pt(d.X0, d.Y0)
	_, err = fc.DrawString(d.Title, pt)
	if err != nil {
		return err
	}

	fc.SetFontSize(d.Size1)
	_, err = fc.DrawString(d.SubTitle, freetype.Pt(d.X1, d.Y1))
	if err != nil {
		return err
	}

	err = jpeg.Encode(d.Merged, d.JPG, nil)
	if err != nil {
		return err
	}

	return nil
}

type Rect struct {
	Name string
	X0   int
	Y0   int
	X1   int
	Y1   int
}
type Pt struct {
	X int
	Y int
}

func main() {
	f, err := file.MustOpen("bg.jpg", "./")
	if err != nil {
		fmt.Println(err)
	}
	rect := &Rect{
		Name: "name",
		X0:   200,
		Y0:   200,
		X1:   200,
		Y1:   200,
	}
	jpg := image.NewRGBA(image.Rect(rect.X0, rect.Y0, rect.X1, rect.Y1))
	err = DrawPoster(&DrawText{
		JPG:    jpg,
		Merged: f,

		Title: "Golang Gin 系列文章",
		X0:    80,
		Y0:    160,
		Size0: 42,

		SubTitle: "---煎鱼",
		X1:       320,
		Y1:       220,
		Size1:    36,
	}, "msyhbd.ttc")

	if err != nil {
		fmt.Println(err)
	}
}
