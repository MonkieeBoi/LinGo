package components

import (
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/widget/material"
)

func layoutWord(th *material.Theme, word string) layout.FlexChild {
	h4 := material.H5(th, word)
	h4.Alignment = text.Middle
	return layout.Rigid(h4.Layout)
}

func layoutWords(th *material.Theme, found map[string]bool) []layout.FlexChild {
	words := make([]layout.FlexChild, 0)
	for word := range found {
		words = append(words, layoutWord(th, word))
	}
	return words
}

func LayoutFound(d *data, u *ui) func(C) D {
	return func(gtx C) D {
		return layout.Flex{
			Axis:    layout.Vertical,
			Spacing: layout.SpaceEnd,
		}.Layout(gtx,
			layoutWords(u.th, d.found)...,
		)
	}
}
