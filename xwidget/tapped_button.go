package xwidget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type TappedButton struct {
	widget.Button

	OnLeftClick  func(*fyne.PointEvent)
	OnRightClick func(*fyne.PointEvent)
}

func NewTappedButton(text string, icon fyne.Resource) *TappedButton {
	w := &TappedButton{}
	w.Icon = icon
	w.Text = text
	w.OnTapped = nil
	w.ExtendBaseWidget(w)
	return w
}

func (w *TappedButton) Tapped(event *fyne.PointEvent) {
	if w.OnLeftClick != nil {
		w.OnLeftClick(event)
	}
}

func (w *TappedButton) TappedSecondary(event *fyne.PointEvent) {
	if w.OnRightClick != nil {
		w.OnRightClick(event)
	}
}
