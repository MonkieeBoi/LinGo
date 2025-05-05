package components

import (
	"strconv"
	"strings"

	"github.com/MonkieeBoi/goana/internal/db"
)

type data struct {
	rack  []rune
	words map[string]bool
	found map[string]bool
	end   bool
}

func (d *data) gen() {
	d.words = make(map[string]bool)
	d.found = make(map[string]bool)
	alpha, err := db.GenAlpha()
	if err != nil {
		return
	}

	words, err := db.GetWords(alpha)
	if err != nil {
		return
	}

	for _, word := range words {
		d.words[word] = true
	}
	d.rack = []rune(strings.ToUpper(alpha))
	d.end = false
}

func (d *data) refresh() {
	words, err := db.GetWords(string(d.rack))
	if err != nil {
		return
	}

	for _, word := range words {
		d.words[word] = true
	}
}

func (d *data) left() string {
	return strconv.Itoa(len(d.words) - len(d.found))
}
