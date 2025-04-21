package components

import (
	"image/color"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/widget"
)

func NewAppWindow(window *app.Window) error {
	d, err := newData()
	if err != nil {
		return err
	}
	u := newUi()
	window.Option(app.Title("LinGo"))
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

func drawApp(gtx C, d *data, u *ui) error {
	paint.Fill(gtx.Ops, color.NRGBA{46, 52, 64, 0xFF})
	for tev, ok := u.input.Update(gtx); ok; tev, ok = u.input.Update(gtx) {
		switch tev.(type) {
		case widget.SubmitEvent:
			word := tev.(widget.SubmitEvent).Text
			if _, ok := d.words[word]; ok {
				d.found[word] = true
			}
			if len(d.found) == len(d.words) {
				if err := d.gen(); err != nil {
					return err
				}
			}
			u.input.SetText("")
		}
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
			LayoutRack(u.th, d.rack),
		),
	)
	return nil
}
