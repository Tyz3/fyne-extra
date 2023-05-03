package xwidget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"time"
)

type ToggleState int

const (
	Off ToggleState = iota
	Inactive
	On
)

var (
	offColor      = color.RGBA{R: 204, A: 255}
	inactiveColor = color.RGBA{R: 230, G: 230, B: 230, A: 255}
	onColor       = color.RGBA{G: 204, B: 102, A: 255}
)

type ToggleButton struct {
	widget.BaseWidget

	state ToggleState

	text       *canvas.Text
	background *canvas.Rectangle
	toggle     *canvas.Rectangle
	tappedArea *TappedLabel

	offset float32

	OnStateChanged func(state ToggleState)
}

func NewToggleButton(text string, c color.Color) *ToggleButton {
	w := &ToggleButton{
		state: Inactive,
	}
	w.text = canvas.NewText(text, color.White)
	w.text.Alignment = fyne.TextAlignCenter
	w.text.TextStyle = fyne.TextStyle{Bold: true}
	w.background = &canvas.Rectangle{}
	w.toggle = &canvas.Rectangle{FillColor: c}
	w.background.Resize(fyne.NewSize(w.text.MinSize().Width+60, w.text.MinSize().Height))
	w.toggle.Resize(fyne.NewSize(w.text.MinSize().Width+20, w.text.MinSize().Height))

	w.offset = (w.background.Size().Width - w.toggle.Size().Width) / 2
	// Default state
	w.background.FillColor = inactiveColor
	w.text.Move(fyne.NewPos(w.background.Size().Width/2, 0))
	w.toggle.Move(fyne.NewPos(w.offset, 0))

	w.tappedArea = NewTappedLabel(w.background.Size())
	w.tappedArea.Resize(w.background.Size())
	w.tappedArea.OnLeftClick = func(event *fyne.PointEvent) {
		switch w.state {
		case Off:
			w.ToggleInactiveState()
		case Inactive:
			w.ToggleOnState()
		case On:
			w.ToggleOffState()
		}
	}
	w.tappedArea.OnRightClick = func(event *fyne.PointEvent) {
		switch w.state {
		case Off:
			w.ToggleOnState()
		case Inactive:
			w.ToggleOffState()
		case On:
			w.ToggleInactiveState()
		}
	}

	w.ExtendBaseWidget(w)
	return w
}

func (w *ToggleButton) ToggleOffState() {
	w.state = Off
	a1 := canvas.NewPositionAnimation(
		w.toggle.Position(),
		fyne.NewPos(0, 0),
		250*time.Millisecond,
		func(position fyne.Position) {
			w.toggle.Move(position)
			w.toggle.Refresh()
		},
	)
	a2 := canvas.NewPositionAnimation(
		w.text.Position(),
		fyne.NewPos(w.background.Size().Width/2-w.offset, 0),
		250*time.Millisecond,
		func(position fyne.Position) {
			w.text.Move(position)
			w.text.Refresh()
		},
	)
	a3 := canvas.NewColorRGBAAnimation(
		w.background.FillColor,
		offColor,
		250*time.Millisecond,
		func(c color.Color) {
			w.background.FillColor = c
			w.background.Refresh()
		},
	)
	a1.Start()
	a2.Start()
	a3.Start()
}

func (w *ToggleButton) ToggleInactiveState() {
	w.state = Inactive
	a1 := canvas.NewPositionAnimation(
		w.toggle.Position(),
		fyne.NewPos(w.offset, 0),
		250*time.Millisecond,
		func(position fyne.Position) {
			w.toggle.Move(position)
			w.toggle.Refresh()
		},
	)
	a2 := canvas.NewPositionAnimation(
		w.text.Position(),
		fyne.NewPos(w.background.Size().Width/2, 0),
		250*time.Millisecond,
		func(position fyne.Position) {
			w.text.Move(position)
			w.text.Refresh()
		},
	)
	a3 := canvas.NewColorRGBAAnimation(
		w.background.FillColor,
		inactiveColor,
		250*time.Millisecond,
		func(c color.Color) {
			w.background.FillColor = c
			w.background.Refresh()
		},
	)
	a1.Start()
	a2.Start()
	a3.Start()
}

func (w *ToggleButton) ToggleOnState() {
	w.state = On
	a1 := canvas.NewPositionAnimation(
		w.toggle.Position(),
		fyne.NewPos(2*w.offset, 0),
		250*time.Millisecond,
		func(position fyne.Position) {
			w.toggle.Move(position)
			w.toggle.Refresh()
		},
	)
	a2 := canvas.NewPositionAnimation(
		w.text.Position(),
		fyne.NewPos(w.background.Size().Width/2+w.offset, 0),
		250*time.Millisecond,
		func(position fyne.Position) {
			w.text.Move(position)
			w.text.Refresh()
		},
	)
	a3 := canvas.NewColorRGBAAnimation(
		w.background.FillColor,
		onColor,
		250*time.Millisecond,
		func(c color.Color) {
			w.background.FillColor = c
			w.background.Refresh()
		},
	)
	a1.Start()
	a2.Start()
	a3.Start()

}

func (w *ToggleButton) CreateRenderer() fyne.WidgetRenderer {
	w.ExtendBaseWidget(w)
	return &ToggleButtonRenderer{objects: []fyne.CanvasObject{
		container.NewWithoutLayout(w.background, w.toggle, w.text, w.tappedArea),
	}}
}

type ToggleButtonRenderer struct {
	objects []fyne.CanvasObject
}

func (r *ToggleButtonRenderer) Destroy() {
}

func (r *ToggleButtonRenderer) Layout(size fyne.Size) {
	for _, o := range r.objects {
		o.Resize(size)
	}
}

func (r *ToggleButtonRenderer) MinSize() fyne.Size {
	return r.objects[0].MinSize()
}

func (r *ToggleButtonRenderer) Objects() []fyne.CanvasObject {
	return r.objects
}

func (r *ToggleButtonRenderer) Refresh() {
	for _, o := range r.objects {
		o.Refresh()
	}
}
