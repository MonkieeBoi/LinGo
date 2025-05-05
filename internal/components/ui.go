package components

import (
	"gioui.org/app"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/explorer"
	"github.com/MonkieeBoi/goana/internal/theme"
)

const (
	TabPlay = iota
	TabAdd
)

type ui struct {
	th       *material.Theme
	tab      int8
	input    *widget.Editor
	editor   *widget.Editor
	button   *widget.Clickable
	explorer *explorer.Explorer
}

func newUi(window *app.Window) *ui {
	return &ui{
		th:       theme.NewTheme(),
		tab:      TabPlay,
		input:    NewTextInput(),
		editor:   NewEditor(),
		button:   &widget.Clickable{},
		explorer: explorer.NewExplorer(window),
	}
}
