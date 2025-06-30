/**
 * @Author Awen
 * @Date 2024/06/01
 * @Email wengaolng@gmail.com
 **/

package go_captcha_2_0_4

import (
	"github.com/cocoyes/go-captcha/click"
	"github.com/cocoyes/go-captcha/rotate"
	"github.com/cocoyes/go-captcha/slide"
)

// Version # of captcha
const Version = "2.0.4"

// NewClickBuilder .
func NewClickBuilder(opts ...click.Option) click.Builder {
	return click.NewBuilder(opts...)
}

// NewSlideBuilder .
func NewSlideBuilder(opts ...slide.Option) slide.Builder {
	return slide.NewBuilder(opts...)
}

// NewRotateBuilder .
func NewRotateBuilder(opts ...rotate.Option) rotate.Builder {
	return rotate.NewBuilder(opts...)
}
