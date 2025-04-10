package main

import (
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type C = layout.Context
type D = layout.Dimensions

func main() {
	go func() {
		window := new(app.Window)
		if err := run(window); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func run(window *app.Window) error {
	window.Option(app.Title("LinGo"))
	theme := material.NewTheme()
	theme.Palette.Bg = color.NRGBA{46, 52, 64, 0xFF}
	theme.Palette.Fg = color.NRGBA{0xFF, 0xFF, 0xFF, 0xFF}
	textInput := widget.Editor{
		SingleLine: true,
		Filter:     "abcdefghijklmnopqrstuvwxyz",
		MaxLen:     15,
	}
	var ops op.Ops
	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)
			paint.Fill(&ops, color.NRGBA{46, 52, 64, 0xFF})
			layout.Flex{
				Axis:    layout.Vertical,
				Spacing: layout.SpaceEnd,
			}.Layout(gtx,
				layout.Rigid(
					func(gtx C) D {
						margins := layout.Inset{
							Top:    unit.Dp(25),
							Bottom: unit.Dp(25),
							Right:  unit.Dp(80),
							Left:   unit.Dp(80),
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
										return layout.UniformInset(unit.Dp(10)).Layout(gtx,
											material.Editor(theme, &textInput, "Enter Word").Layout)
									},
								)
							},
						)
					},
				),
			)
			e.Frame(gtx.Ops)
		}
	}
}
