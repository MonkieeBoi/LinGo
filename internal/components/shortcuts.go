package components

import (
	"errors"

	"gioui.org/io/key"
	"gioui.org/x/explorer"
	"github.com/MonkieeBoi/goana/internal/db"
)

func pressed(gtx C, filter key.Filter) bool {
	ke, ok := gtx.Event(filter)
	switch ke := ke.(type) {
	case key.Event:
		return ok && ke.State == key.Press
	}
	return false
}

func ProcessShortcuts(gtx C, d *data, u *ui) {
	if pressed(gtx, key.Filter{Name: "F", Required: key.ModShortcut}) {
		go func() {
			io, err := u.explorer.ChooseFile()
			if err != nil {
				if errors.Is(err, explorer.ErrNotAvailable) {
					u.tab = TabAdd
				}
			} else {
				if err := db.AddWords(io); err != nil {
					u.tab = TabAdd
				}
				d.refresh()
			}
		}()
	}

	if pressed(gtx, key.Filter{Name: "W", Required: key.ModShortcut}) {
		u.tab = TabAdd
	}

	if pressed(gtx, key.Filter{Name: "P", Required: key.ModShortcut}) {
		u.tab = TabPlay
		d.refresh()
	}

	if pressed(gtx, key.Filter{Focus: u.input, Name: key.NameDeleteBackward, Required: key.ModShortcut}) {
		u.input.SetText("")
	}

	if pressed(gtx, key.Filter{Focus: u.input, Name: key.NameEscape}) {
		d.end = true
		for k := range d.words {
			d.found[k] = true
		}
	}
}
