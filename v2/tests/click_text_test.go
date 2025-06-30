package tests

import (
	"encoding/json"
	"fmt"
	"image"
	"log"
	"testing"

	"github.com/cocoyes/go-captcha/v2/base/option"
	"github.com/cocoyes/go-captcha/v2/click"
	"github.com/golang/freetype/truetype"
)

var textCapt click.Captcha

func init() {
	builder := click.NewBuilder(
		click.WithRangeLen(option.RangeVal{Min: 2, Max: 4}),
		click.WithRangeVerifyLen(option.RangeVal{Min: 2, Max: 2}),
		click.WithDisabledRangeVerifyLen(true),
		click.WithIsThumbNonDeformAbility(false),
	)

	fontN, err := loadFont("resources/fonts/fzshengsksjw_cu.ttf")
	if err != nil {
		log.Fatalln(err)
	}

	bgImage, err := loadJpg("resources/images/1.jpg")
	if err != nil {
		log.Fatalln(err)
	}

	builder.SetResources(
		click.WithChars(map[string][]string{
			"cn": {"鼓", "鼎", "默", "黔", "黑", "黎", "黍", "黄", "麻", "麸", "麦", "鹿"},
			"en": {"A1", "B2", "C3", "D4", "E5", "F6", "G7", "H8", "I9", "J0"},
		}),
		click.WithFonts([]*truetype.Font{

			fontN,
		}),
		click.WithBackgrounds([]image.Image{
			bgImage,
		}),
		//click.WithThumbBackgrounds([]image.Image{
		//	thumbImage,
		//}),
	)

	textCapt = builder.Make()
}

func TestClickTextCaptcha(t *testing.T) {
	captData, err := textCapt.Generate("en")
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

	err = captData.GetMasterImage().SaveToFile("D:\\AI\\master.jpg", option.QualityNone)
	if err != nil {
		fmt.Println(err)
	}
	err = captData.GetThumbImage().SaveToFile("D:\\AI\\thumb.png")
	if err != nil {
		fmt.Println(err)
	}
}
