package xwidget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type TappedIcon struct {
	widget.Icon

	OnLeftClick  func(*fyne.PointEvent)
	OnRightClick func(*fyne.PointEvent)
}

func NewTappedIcon(icon fyne.Resource) *TappedIcon {
	w := &TappedIcon{}
	w.Icon.SetResource(icon)
	w.ExtendBaseWidget(w)
	return w
}

func (w *TappedIcon) Tapped(event *fyne.PointEvent) {
	if w.OnLeftClick != nil {
		w.OnLeftClick(event)
	}
}

func (w *TappedIcon) TappedSecondary(event *fyne.PointEvent) {
	if w.OnRightClick != nil {
		w.OnRightClick(event)
	}
}
