package components

import (
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/MonkieeBoi/LinGo/internal/theme"
)

type ui struct {
	th    *material.Theme
	input *widget.Editor
}

func newUi() *ui {
	return &ui{
		th: theme.NewTheme(),
		input: &widget.Editor{
			Filter:     "abcdefghijklmnopqrstuvwxyz",
			MaxLen:     15,
			SingleLine: true,
			Submit:     true,
		},
	}
}

