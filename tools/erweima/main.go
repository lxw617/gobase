package main

import (
	"fmt"
	"image/png"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

func main() {
	// Create the barcode
	qrCode, _ := qr.Encode("Hello World", qr.M, qr.Auto)

	// Scale the barcode to 200x200 pixels
	qrCode, _ = barcode.Scale(qrCode, 200, 200)

	// create the output file
	file, _ := os.Create("./godaily/erweima/qrcode.png")
	defer file.Close()

	// encode the barcode as png
	if err := png.Encode(file, qrCode); err != nil {
		fmt.Println(err)
	}
}
