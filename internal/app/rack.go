package app

import (
	"image/color"

	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func tile(th *material.Theme, s rune) layout.FlexChild {
	return layout.Rigid(
		func(gtx C) D {
			margins := layout.Inset{
				Top:    unit.Dp(25),
				Bottom: unit.Dp(25),
				Right:  unit.Dp(5),
				Left:   unit.Dp(5),
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
							textH4 := material.H4(th, string(s))
							textH4.Alignment = text.Middle
							gtx.Constraints.Min.X = int(unit.Dp(33))
							return layout.UniformInset(unit.Dp(10)).Layout(gtx, textH4.Layout)
						},
					)
				},
			)
		},
	)
}

func empty(th *material.Theme) layout.FlexChild {
	h5 := material.H5(th,
		"Failed to find words\n"+
			"C-w to input words\n"+
			"C-f to add words from a file (one word per line)\n"+
			"C-p to switch back here")
	h5.Alignment = text.Middle
	return layout.Rigid(
		h5.Layout,
	)
}

func newRack(th *material.Theme, rack []rune) []layout.FlexChild {
	t := make([]layout.FlexChild, 0)
	for _, r := range rack {
		t = append(t, tile(th, r))
	}
	if len(t) == 0 {
		t = append(t, empty(th))
	}
	return t
}

func layoutRack(d *data, u *ui) func(C) D {
	return func(gtx C) D {
		return layout.Flex{
			Axis:    layout.Horizontal,
			Spacing: layout.SpaceSides,
		}.Layout(gtx,
			newRack(u.th, d.rack)...,
		)
	}

}
