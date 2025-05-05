package components

import (
	"image/color"
	"strings"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/MonkieeBoi/goana/internal/db"
)

type (
	C = layout.Context
	D = layout.Dimensions
)

func NewAppWindow(window *app.Window) error {
	u := newUi(window)
	d := &data{}
	if has, _ := db.HasWords(); has {
		d.gen()
	}
	window.Option(app.Title("Goana"))
	var ops op.Ops
	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)
			if err := drawApp(gtx, d, u); err != nil {
				return err
			}
			e.Frame(gtx.Ops)
		}
	}
}

func processSubmit(gtx C, d *data, u *ui) error {
	for tev, ok := u.input.Update(gtx); ok; tev, ok = u.input.Update(gtx) {
		switch tev.(type) {
		case widget.SubmitEvent:
			word := tev.(widget.SubmitEvent).Text
			if _, ok := d.words[word]; ok {
				d.found[word] = true
			}
			if len(d.found) == len(d.words) {
				d.gen()
			}
			u.input.SetText("")
		}
	}
	return nil
}

func processButton(gtx C, u *ui) {
	if _, ok := u.button.Update(gtx); !ok {
		return
	}
	db.AddWords(strings.NewReader(u.editor.Text()))
	u.editor.SetText("")
}

func processEvents(gtx C, d *data, u *ui) error {
	ProcessShortcuts(gtx, d, u)
	processButton(gtx, u)
	if err := processSubmit(gtx, d, u); err != nil {
		return err
	}
	return nil
}

func drawPlay(gtx C, d *data, u *ui) {
	if has, _ := db.HasWords() ; len(d.words) == 0 && has {
		d.gen()
	}
	layout.Flex{
		Axis:      layout.Vertical,
		Spacing:   layout.SpaceEnd,
		Alignment: layout.Middle,
	}.Layout(gtx,
		layout.Rigid(
			LayoutTextInput(d, u),
		),
		layout.Rigid(
			LayoutRack(d, u),
		),
		layout.Rigid(
			LayoutFound(d, u),
		),
	)
}

func drawAdd(gtx C, u *ui) {
	layout.Flex{
		Axis:      layout.Vertical,
		Spacing:   layout.SpaceEnd,
		Alignment: layout.Middle,
	}.Layout(gtx,
		layout.Rigid(
			func(gtx C) D {
				margins := layout.Inset{
					Top:    unit.Dp(25),
					Bottom: unit.Dp(25),
				}
				border := widget.Border{
					Color:        u.th.ContrastFg,
					CornerRadius: unit.Dp(3),
					Width:        unit.Dp(2),
				}
				return margins.Layout(gtx,
					func(gtx C) D {
						return border.Layout(gtx,
							func(gtx C) D {
								gtx.Constraints.Max.Y *= 3
								gtx.Constraints.Max.Y /= 4
								gtx.Constraints.Min.Y = gtx.Constraints.Max.Y
								gtx.Constraints.Max.X -= 50
								return layout.UniformInset(unit.Dp(10)).Layout(gtx,
									material.Editor(u.th, u.editor, "Enter one word per line").Layout)
							},
						)
					},
				)
			},
		),
		layout.Rigid(
			func(gtx C) D {
				gtx.Constraints.Max.X = 100
				return material.Button(u.th, u.button, "Add").Layout(gtx)
			},
		),
	)
}

func drawApp(gtx C, d *data, u *ui) error {
	paint.Fill(gtx.Ops, color.NRGBA{46, 52, 64, 0xFF})
	if err := processEvents(gtx, d, u); err != nil {
		return err
	}
	switch u.tab {
	case TabPlay:
		drawPlay(gtx, d, u)
	case TabAdd:
		drawAdd(gtx, u)
	}
	return nil
}
