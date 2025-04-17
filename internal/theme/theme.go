package theme

import (
	"image/color"

	"gioui.org/widget/material"
)

var nord0 = color.NRGBA{46, 52, 64, 0xFF}
var white = color.NRGBA{0xFF, 0xFF, 0xFF, 0xFF}

func NewTheme() *material.Theme {
	th := material.NewTheme()

	th.Palette.Bg = nord0
	th.Palette.Fg = white
	th.Palette.ContrastBg, th.Palette.ContrastFg = th.Palette.ContrastFg, th.Palette.ContrastBg

	return th
}
