package app

import "gioui.org/widget"

func newEditor() *widget.Editor {
	return &widget.Editor{
		Filter: "abcdefghijklmnopqrstuvwxyz\n",
	}
}
