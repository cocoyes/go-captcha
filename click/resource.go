/**
 * @Author Awen
 * @Date 2024/06/01
 * @Email wengaolng@gmail.com
 **/

package click

import (
	"errors"
	"image"

	"github.com/cocoyes/go-captcha/base/helper"
	"github.com/cocoyes/go-captcha/base/logger"
	"github.com/golang/freetype/truetype"
)

// Resources defines the resources for the CAPTCHA
type Resources struct {
	chars                map[string][]string
	shapeMaps            map[string]image.Image
	shapes               []string
	rangFonts            []*truetype.Font
	rangBackgrounds      []image.Image
	rangThumbBackgrounds []image.Image
}

// NewResources .
func NewResources() *Resources {
	return &Resources{}
}

type Resource func(*Resources)

var (
	ChineseCharLenErr = errors.New("the chinese char length must be equal to 1")
	CharLenErr        = errors.New("the char length must be less than or equal to 2")
)

// WithChars is to set characters
func WithChars(chars map[string][]string) Resource {
	return func(resources *Resources) {
		if len(chars) > 0 {
			for groupName, charList := range chars {
				for _, char := range charList {
					if helper.IsChineseChar(char) {
						if helper.LenChineseChar(char) > 1 {
							logger.Logx.Warnf("WithChars(): group=%s char=%s, err=%v", groupName, char, ChineseCharLenErr)
							return
						}
					} else {
						if helper.LenChineseChar(char) > 2 {
							logger.Logx.Warnf("WithChars(): group=%s char=%s, err=%v", groupName, char, CharLenErr)
							return
						}
					}
				}
			}
		}

		resources.chars = chars
	}
}

// WithShapes is to set shape
func WithShapes(shapeMaps map[string]image.Image) Resource {
	return func(resources *Resources) {
		resources.shapeMaps = shapeMaps
		var shapes = make([]string, 0, len(shapeMaps))
		for name, _ := range shapeMaps {
			shapes = append(shapes, name)
		}
		resources.shapes = shapes
	}
}

// WithFonts is to set font
func WithFonts(fonts []*truetype.Font) Resource {
	return func(resources *Resources) {
		resources.rangFonts = fonts
	}
}

// WithBackgrounds is to set background image
func WithBackgrounds(images []image.Image) Resource {
	return func(resources *Resources) {
		resources.rangBackgrounds = images
	}
}

// WithThumbBackgrounds is to set thumbnail background image
func WithThumbBackgrounds(images []image.Image) Resource {
	return func(resources *Resources) {
		resources.rangThumbBackgrounds = images
	}
}
