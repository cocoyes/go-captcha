package tests

import (
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/cocoyes/go-captcha/base/canvas"
	"github.com/cocoyes/go-captcha/base/codec"
	"github.com/cocoyes/go-captcha/base/helper"
	"github.com/cocoyes/go-captcha/base/option"
	"github.com/cocoyes/go-captcha/click"
	"golang.org/x/image/draw"
)

var shapeCapt click.Captcha

func init() {
	builder := click.NewBuilder(
		click.WithRangeLen(option.RangeVal{Min: 3, Max: 6}),
		click.WithRangeVerifyLen(option.RangeVal{Min: 2, Max: 3}),
		click.WithRangeThumbBgDistort(1),
	)

	bgImage, err := loadJpg("resources/images/1.jpg")
	if err != nil {
		log.Fatalln(err)
	}

	bgImage1, err := loadJpg("resources/images/1.jpg")
	if err != nil {
		log.Fatalln(err)
	}

	shapes := getShapeMaps()

	builder.SetResources(
		click.WithShapes(shapes),
		click.WithBackgrounds([]image.Image{
			bgImage,
			bgImage1,
		}),
		//click.WithThumbBackgrounds([]image.Image{
		//	img1,
		//}),
	)

	shapeCapt = builder.MakeWithShape()
}

func getShapeMaps() map[string]image.Image {
	shapeImage1, err := loadPng("resources/images/shape/shape1.png")
	if err != nil {
		log.Fatalln(err)
	}

	shapeImage2, err := loadPng("resources/images/shape/shape2.png")
	if err != nil {
		log.Fatalln(err)
	}

	shapeImage3, err := loadPng("resources/images/shape/shape3.png")
	if err != nil {
		log.Fatalln(err)
	}

	shapeImage4, err := loadPng("resources/images/shape/shape4.png")
	if err != nil {
		log.Fatalln(err)
	}

	shapeImage5, err := loadPng("resources/images/shape/shape5.png")
	if err != nil {
		log.Fatalln(err)
	}

	shapeImage6, err := loadPng("resources/images/shape/shape6.png")
	if err != nil {
		log.Fatalln(err)
	}

	return map[string]image.Image{
		"shape1": shapeImage1,
		"shape2": shapeImage2,
		"shape3": shapeImage3,
		"shape4": shapeImage4,
		"shape5": shapeImage5,
		"shape6": shapeImage6,
	}
}

func TestClickShapeCaptcha(t *testing.T) {
	captData, err := shapeCapt.Generate("")
	if err != nil {
		log.Fatalln(err)
	}

	dotData := captData.GetData()
	if dotData == nil {
		log.Fatalln(">>>>> generate err")
	}

	dots, _ := json.Marshal(dotData)
	fmt.Println(string(dots))
	fmt.Println(captData.GetMasterImage().ToBase64())
	fmt.Println(captData.GetThumbImage().ToBase64())

	err = captData.GetMasterImage().SaveToFile("../.cache/master.jpg", option.QualityNone)
	if err != nil {
		fmt.Println(err)
	}
	err = captData.GetThumbImage().SaveToFile("../.cache/thumb.png")
	if err != nil {
		fmt.Println(err)
	}
}

func TestScale(t *testing.T) {
	shapeBytes4, err := ioutil.ReadFile("resources/images/shape/shape4.png")
	if err != nil {
		log.Fatalln(err)
	}
	shapeImage4, err := codec.DecodeByteToPng(shapeBytes4)
	if err != nil {
		log.Fatalln(err)
	}

	var colorArr []color.RGBA
	bc, _ := helper.ParseHexColor("#ffffff")
	cArr := append(colorArr, bc)
	cvs := canvas.CreatePaletteCanvas(30, 30, cArr)
	draw.BiLinear.Scale(cvs.Get(), cvs.Bounds(), shapeImage4, shapeImage4.Bounds(), draw.Over, nil)
	pngFile, err := os.Create("../.cache/output.png")
	if err != nil {
		log.Fatalf("Error creating PNG file: %v", err)
	}
	defer pngFile.Close()

	err = png.Encode(pngFile, cvs)
}
