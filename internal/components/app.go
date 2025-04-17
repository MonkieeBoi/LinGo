package components

import (
	"fmt"
	"image/color"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/widget"
	"github.com/MonkieeBoi/LinGo/internal/theme"
)

func NewAppWindow(window *app.Window) error {
	window.Option(app.Title("LinGo"))
	th := theme.NewTheme()
	textInput := widget.Editor{
		Filter:     "abcdefghijklmnopqrstuvwxyz",
		MaxLen:     15,
		SingleLine: true,
		Submit:     true,
	}
	var ops op.Ops
	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)
			paint.Fill(&ops, color.NRGBA{46, 52, 64, 0xFF})
			for tev, ok := textInput.Update(gtx); ok; tev, ok = textInput.Update(gtx) {
				switch tev.(type) {
				case widget.SubmitEvent:
					word := tev.(widget.SubmitEvent).Text
					textInput.SetText("")
					fmt.Println(word)
				}
			}
			layout.Flex{
				Axis:      layout.Vertical,
				Spacing:   layout.SpaceEnd,
				Alignment: layout.Middle,
			}.Layout(gtx,
				layout.Rigid(
					LayoutTextInput(&textInput, th),
				),
				layout.Rigid(
					LayoutRack(th, []rune{'E', 'I', 'N', 'R', 'S', 'T'}),
				),
			)
			e.Frame(gtx.Ops)
		}
	}
}
