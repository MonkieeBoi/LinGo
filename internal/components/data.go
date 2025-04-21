package components

import (
	"github.com/MonkieeBoi/LinGo/internal/db"
	"strings"
	"strconv"
)

type data struct {
	rack  []rune
	words map[string]bool
	found map[string]bool
}

func newData() (*data, error) {
	d := data{}
	return &d, d.gen()
}

func (d *data) gen() error {
	d.words = make(map[string]bool)
	d.found = make(map[string]bool)
	alpha, err := db.GenAlpha()
	if err != nil {
		return err
	}

	words, err := db.GetWords(alpha)
	if err != nil {
		return err
	}

	for _, word := range words {
		d.words[word] = true
	}
	d.rack = []rune(strings.ToUpper(alpha))
	return nil
}

func (d *data) left() string {
	return strconv.Itoa(len(d.words) - len(d.found))
}
