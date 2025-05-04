package components

import "gioui.org/widget"

func NewEditor() *widget.Editor {
	return &widget.Editor{
		Filter: "abcdefghijklmnopqrstuvwxyz\n",
	}
}
