package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"io/ioutil"
	"os"

	"base/tools/writepic/file"
	"base/tools/writepic/qrcode"

	"github.com/boombuler/barcode/qr"
	"github.com/golang/freetype"
)

type ArticlePoster struct {
	PosterName string
	Qr         *qrcode.QrCode
}

func NewArticlePoster(posterName string, qr *qrcode.QrCode) *ArticlePoster {
	return &ArticlePoster{
		PosterName: posterName,
		Qr:         qr,
	}
}

func (a *ArticlePoster) CheckMergedImage(path string) bool {
	return !file.CheckNotExist(path + a.PosterName)
}

func (a *ArticlePoster) OpenMergedImage(path string) (*os.File, error) {
	f, err := file.MustOpen(a.PosterName, path)
	if err != nil {
		return nil, err
	}

	return f, nil
}

type ArticlePosterBg struct {
	Name string
	*ArticlePoster
	*Rect
	*Pt
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

func NewArticlePosterBg(name string, ap *ArticlePoster, rect *Rect, pt *Pt) *ArticlePosterBg {
	return &ArticlePosterBg{
		Name:          name,
		ArticlePoster: ap,
		Rect:          rect,
		Pt:            pt,
	}
}

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

func (a *ArticlePosterBg) DrawPoster(d *DrawText, fontName string) error {
	fontSource := "./"
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

func (a *ArticlePosterBg) Generate() (string, string, error) {
	// fullPath := qrcode.GetQrCodeFullPath()
	// fileName, path, err := a.Qr.Encode(fullPath)
	// if err != nil {
	// 	return "", "", err
	// }
	path := "/base/tools/writepic/qrcode.png"
	fileName := "qrcode"
	if !a.CheckMergedImage(path) {
		mergedF, err := a.OpenMergedImage(path)
		if err != nil {
			return "", "", err
		}
		defer mergedF.Close()

		bgF, err := file.MustOpen(a.Name, path)
		if err != nil {
			return "", "", err
		}
		defer bgF.Close()

		qrF, err := file.MustOpen(fileName, path)
		if err != nil {
			return "", "", err
		}
		defer qrF.Close()

		bgImage, err := jpeg.Decode(bgF)
		if err != nil {
			return "", "", err
		}
		qrImage, err := jpeg.Decode(qrF)
		if err != nil {
			return "", "", err
		}

		jpg := image.NewRGBA(image.Rect(a.Rect.X0, a.Rect.Y0, a.Rect.X1, a.Rect.Y1))

		draw.Draw(jpg, jpg.Bounds(), bgImage, bgImage.Bounds().Min, draw.Over)
		draw.Draw(jpg, jpg.Bounds(), qrImage, qrImage.Bounds().Min.Sub(image.Pt(a.Pt.X, a.Pt.Y)), draw.Over)

		err = a.DrawPoster(&DrawText{
			JPG:    jpg,
			Merged: mergedF,

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
			return "", "", err
		}
	}

	return fileName, path, nil
}

const (
	QRCODE_URL = "https://github.com/EDDYCJY/blog#gin%E7%B3%BB%E5%88%97%E7%9B%AE%E5%BD%95"
)

func main() {
	qr := qrcode.NewQrCode(QRCODE_URL, 300, 300, qr.M, qr.Auto)
	posterName := "poster-" + qrcode.GetQrCodeFileName(qr.URL) + qr.GetQrCodeExt()
	articlePoster := NewArticlePoster(posterName, qr)
	a := NewArticlePosterBg(
		"2036042932-5b3cedeb3bb16_fix732.jfif",
		articlePoster,
		&Rect{
			X0: 0,
			Y0: 0,
			X1: 550,
			Y1: 700,
		},
		&Pt{
			X: 125,
			Y: 298,
		},
	)
	_, filePath, err := a.Generate()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(filePath)
}

/*
1、freetype.NewContext：创建一个新的 Context，会对其设置一些默认值
func NewContext() *Context {
    return &Context{
        r:        raster.NewRasterizer(0, 0),
        fontSize: 12,
        dpi:      72,
        scale:    12 << 6,
    }
}
2、fc.SetDPI：设置屏幕每英寸的分辨率
3、fc.SetFont：设置用于绘制文本的字体
4、fc.SetFontSize：以磅为单位设置字体大小
5、fc.SetClip：设置剪裁矩形以进行绘制
6、fc.SetDst：设置目标图像
7、fc.SetSrc：设置绘制操作的源图像，通常为 image.Uniform
var (
        // Black is an opaque black uniform image.
        Black = NewUniform(color.Black)
        // White is an opaque white uniform image.
        White = NewUniform(color.White)
        // Transparent is a fully transparent uniform image.
        Transparent = NewUniform(color.Transparent)
        // Opaque is a fully opaque uniform image.
        Opaque = NewUniform(color.Opaque)
)
8、fc.DrawString：根据 Pt 的坐标值绘制给定的文本内容
*/
