package components

import (
	"image/color"
	"strings"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/widget"
	"github.com/MonkieeBoi/LinGo/internal/db"
	"github.com/MonkieeBoi/LinGo/internal/theme"
)

type data struct {
	rack  []rune
	words map[string]bool
	found map[string]bool
}

func newData() (data, error) {
	d := data{
		words: make(map[string]bool),
		found: make(map[string]bool),
	}
	alpha, err := db.GenAlpha()
	if err != nil {
		return d, err
	}

	words, err := db.GetWords(alpha)
	if err != nil {
		return d, err
	}

	for _, word := range words {
		d.words[word] = true
	}
	d.rack = []rune(strings.ToUpper(alpha))
	return d, nil
}

func NewAppWindow(window *app.Window) error {
	d, err := newData()
	if err != nil {
		return err
	}

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
					if _, ok := d.words[word]; ok {
						d.found[word] = true
					}
					if len(d.found) == len(d.words) {
						d, err = newData()
						if err != nil {
							return err
						}
					}
					textInput.SetText("")
				}
			}
			layout.Flex{
				Axis:      layout.Vertical,
				Spacing:   layout.SpaceEnd,
				Alignment: layout.Middle,
			}.Layout(gtx,
				layout.Rigid(
					LayoutTextInput(&textInput, th, d),
				),
				layout.Rigid(
					LayoutRack(th, d.rack),
				),
			)
			e.Frame(gtx.Ops)
		}
	}
}
