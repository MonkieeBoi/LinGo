package theme

import (
	"image/color"

	"gioui.org/widget/material"
)

var nord0 = color.NRGBA{46, 52, 64, 0xFF}
var nord1 = color.NRGBA{59, 66, 82, 0xFF}
var nord4 = color.NRGBA{216, 222, 233, 0xFF}
var white = color.NRGBA{0xFF, 0xFF, 0xFF, 0xFF}

func NewTheme() *material.Theme {
	th := material.NewTheme()

	th.Palette.Bg = nord0
	th.Palette.Fg = white
	th.Palette.ContrastBg = nord1
	th.Palette.ContrastFg = nord4

	return th
}
