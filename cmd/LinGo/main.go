package main

import (
	"fmt"
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/text"
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

func tile(th *material.Theme, s string) layout.FlexChild {
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
							textH4 := material.H4(th, s)
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

func run(window *app.Window) error {
	window.Option(app.Title("LinGo"))
	th := material.NewTheme()
	th.Palette.Bg = color.NRGBA{46, 52, 64, 0xFF}
	th.Palette.Fg = color.NRGBA{0xFF, 0xFF, 0xFF, 0xFF}
	th.Palette.ContrastBg, th.Palette.ContrastFg = th.Palette.ContrastFg, th.Palette.ContrastBg
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
					func(gtx C) D {
						margins := layout.Inset{
							Top:    unit.Dp(25),
							Bottom: unit.Dp(25),
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
										gtx.Constraints.Max.X = int(unit.Dp(400))
										return layout.UniformInset(unit.Dp(10)).Layout(gtx,
											material.Editor(th, &textInput, "Enter Word").Layout)
									},
								)
							},
						)
					},
				),
				layout.Rigid(
					func(gtx C) D {
						return layout.Flex{
							Axis:    layout.Horizontal,
							Spacing: layout.SpaceSides,
						}.Layout(gtx,
							tile(th, "A"),
							tile(th, "E"),
							tile(th, "I"),
							tile(th, "N"),
							tile(th, "R"),
							tile(th, "S"),
							tile(th, "T"),
						)
					},
				),
			)
			e.Frame(gtx.Ops)
		}
	}
}
