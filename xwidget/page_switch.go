package xwidget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"strconv"
)

type PageSwitch struct {
	widget.BaseWidget

	currentPage int
	pages       int
	items       []fyne.CanvasObject

	OnPageChanged func(int)
}

func NewPageSwitch() *PageSwitch {
	w := &PageSwitch{
		currentPage: 1,
		pages:       1,
	}
	w.ExtendBaseWidget(w)
	return w
}

func (w *PageSwitch) SetPages(pages int) {
	w.pages = pages
	if w.currentPage > w.pages {
		w.setPage(w.pages)
	} else {
		w.update()
	}
}

func (w *PageSwitch) setPage(page int) {
	w.currentPage = page
	if w.OnPageChanged != nil {
		w.OnPageChanged(page)
	}
	w.update()
}

func (w *PageSwitch) update() {
	w.items = nil

	for page := 1; page <= w.pages; page++ {
		page := page
		btn := widget.NewButton(strconv.Itoa(page), func() {
			w.setPage(page)
		})

		if page == w.currentPage {
			btn.Importance = widget.HighImportance
			btn.OnTapped = nil
		}

		if page == 1 {
			w.items = append(w.items, btn)
			continue
		}

		if 1 <= w.currentPage && w.currentPage <= 3 {
			if 1 <= page && page <= 5 {
				w.items = append(w.items, btn)
			}
			continue
		}

		window := w.pages - w.currentPage
		if window <= 1 {
			if w.currentPage-page+window <= 4 {
				w.items = append(w.items, btn)
			}
			continue
		}

		if w.currentPage-page > 2 || page-w.currentPage > 2 {
			continue
		} else {
			w.items = append(w.items, btn)
		}

	}

	if w.pages != w.currentPage {
		w.items = append(w.items,
			&canvas.Rectangle{
				FillColor: color.White,
			},
			widget.NewButtonWithIcon("", theme.NavigateNextIcon(), func() {
				w.setPage(w.currentPage + 1)
			}),
		)
	}

	w.Refresh()
}

func (w *PageSwitch) CreateRenderer() fyne.WidgetRenderer {
	w.ExtendBaseWidget(w)
	return &PageToolbarRenderer{
		objects: []fyne.CanvasObject{container.NewHBox()},
		items:   &w.items,
	}
}

type PageToolbarRenderer struct {
	objects []fyne.CanvasObject
	items   *[]fyne.CanvasObject
}

func (r *PageToolbarRenderer) Destroy() {
}

func (r *PageToolbarRenderer) Layout(size fyne.Size) {
	for _, o := range r.objects {
		o.Resize(size)
	}
}

func (r *PageToolbarRenderer) MinSize() fyne.Size {
	minSize := fyne.NewSize(0, 0)
	for _, child := range r.objects {
		minSize = minSize.Max(child.MinSize())
	}
	return minSize
}

func (r *PageToolbarRenderer) Objects() []fyne.CanvasObject {
	return r.objects
}

func (r *PageToolbarRenderer) Refresh() {
	c := r.objects[0].(*fyne.Container)
	c.RemoveAll()
	for _, o := range *r.items {
		c.Add(o)
	}
	c.Refresh()
}
