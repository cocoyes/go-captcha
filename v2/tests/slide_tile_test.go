package tests

import (
	"encoding/json"
	"fmt"
	"image"
	"log"
	"testing"

	"github.com/cocoyes/go-captcha/v2/base/option"
	"github.com/cocoyes/go-captcha/v2/slide"
)

var slideTileCapt slide.Captcha

func init() {
	builder := slide.NewBuilder(
	//slide.WithGenGraphNumber(2),
	//slide.WithEnableGraphVerticalRandom(true),
	)

	bgImage, err := loadJpg("resources/images/1.jpg")
	if err != nil {
		log.Fatalln(err)
	}

	bgImage1, err := loadJpg("resources/images/1.jpg")
	if err != nil {
		log.Fatalln(err)
	}

	graphs := getSlideTileGraphArr()

	builder.SetResources(
		slide.WithGraphImages(graphs),
		slide.WithBackgrounds([]image.Image{
			bgImage,
			bgImage1,
		}),
		//slide.WithThumbBackgrounds([]image.Image{
		//	img1,
		//}),
	)

	slideTileCapt = builder.Make()
	//slideTileCapt = builder.MakeWithRegion()
}

func getSlideTileGraphArr() []*slide.GraphImage {
	tileImage1, err := loadPng("resources/images/tile/tile-1/tile.png")
	if err != nil {
		log.Fatalln(err)
	}

	tileShadowImage1, err := loadPng("resources/images/tile/tile-1/tile-shadow.png")
	if err != nil {
		log.Fatalln(err)
	}
	tileMaskImage1, err := loadPng("resources/images/tile/tile-1/tile-mask.png")
	if err != nil {
		log.Fatalln(err)
	}

	tileImage2, err := loadPng("resources/images/tile/tile-2/tile.png")
	if err != nil {
		log.Fatalln(err)
	}

	tileShadowImage2, err := loadPng("resources/images/tile/tile-2/tile-shadow.png")
	if err != nil {
		log.Fatalln(err)
	}

	tileMaskImage2, err := loadPng("resources/images/tile/tile-2/tile-mask.png")
	if err != nil {
		log.Fatalln(err)
	}

	return []*slide.GraphImage{
		{
			OverlayImage: tileImage1,
			ShadowImage:  tileShadowImage1,
			MaskImage:    tileMaskImage1,
		},
		{
			OverlayImage: tileImage2,
			ShadowImage:  tileShadowImage2,
			MaskImage:    tileMaskImage2,
		},
	}
}

func TestSlideTileCaptcha(t *testing.T) {
	captData, err := slideTileCapt.Generate()
	if err != nil {
		log.Fatalln(err)
	}

	blockData := captData.GetData()
	if blockData == nil {
		log.Fatalln(">>>>> generate err")
	}

	block, _ := json.Marshal(blockData)
	fmt.Println(string(block))
	fmt.Println(captData.GetMasterImage().ToBase64())
	fmt.Println(captData.GetTileImage().ToBase64())

	err = captData.GetMasterImage().SaveToFile("../.cache/master.jpg", option.QualityNone)
	if err != nil {
		fmt.Println(err)
	}
	err = captData.GetTileImage().SaveToFile("../.cache/thumb.png")
	if err != nil {
		fmt.Println(err)
	}
}
