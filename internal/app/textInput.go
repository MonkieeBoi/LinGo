package app

import (
	"image/color"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func newTextInput() *widget.Editor {
	return &widget.Editor{
		Filter:     "abcdefghijklmnopqrstuvwxyz",
		MaxLen:     15,
		SingleLine: true,
		Submit:     true,
	}
}

func layoutTextInput(d *data, u *ui) func(C) D {
	return func(gtx C) D {
		margins := layout.Inset{
			Top:    unit.Dp(25),
			Bottom: unit.Dp(25),
		}
		border := widget.Border{
			Color:        color.NRGBA{R: 204, G: 204, B: 204, A: 0xFF},
			CornerRadius: unit.Dp(3),
			Width:        unit.Dp(2),
		}
		return margins.Layout(gtx,
			func(gtx C) D {
				return border.Layout(gtx,
					func(gtx C) D {
						gtx.Constraints.Max.X = int(unit.Dp(400))
						text := d.left()
						if len(d.words)-len(d.found) == 1 {
							text += " word remaining"
						} else {
							text += " words remaining"
						}
						return layout.UniformInset(unit.Dp(10)).Layout(gtx,
							material.Editor(u.th, u.input, text).Layout)
					},
				)
			},
		)
	}
}
