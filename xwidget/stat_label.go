package xwidget

import (
	"fmt"
	"fyne.io/fyne/v2/widget"
)

type StatLabel struct {
	widget.Label

	prefix string
}

func NewStatLabel(prefix string) *StatLabel {
	w := &StatLabel{
		prefix: prefix,
	}
	w.ExtendBaseWidget(w)
	return w
}

func (w *StatLabel) SetNumber(num int) {
	w.SetText(fmt.Sprint(w.prefix, ": ", num))
}

func (w *StatLabel) SetNumberOf(num, of int) {
	w.SetText(fmt.Sprint(w.prefix, ": ", num, "/", of))
}
