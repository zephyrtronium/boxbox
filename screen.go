package boxbox

import "github.com/nsf/termbox-go"

type Screen struct {
	x, y, w, h int
}

func NewScreen(x, y, w, h int) *Screen {
	return &Screen{x, y, w, h}
}

func (scr *Screen) Subscreen(x, y, w, h int) *Screen {
	if x > scr.w || y > scr.h {
		return &Screen{x + scr.x, y + scr.y, 0, 0}
	}
	if x+w > scr.w {
		w = scr.w - x
	}
	if x < 0 {
		w += x
		x = 0
	}
	if y+h > scr.h {
		h = scr.h - y
	}
	if y < 0 {
		h += y
		y = 0
	}
	return &Screen{x + scr.x, y + scr.y, w, h}
}

func (scr *Screen) SetCell(x, y int, ch rune, fg, bg termbox.Attribute) {
	if !(x < scr.x || y < scr.y || x > scr.w || y > scr.h) {
		termbox.SetCell(x+scr.x, y+scr.y, ch, fg, bg)
	}
}

func (scr *Screen) Position() (x, y int) {
	return scr.x, scr.y
}

func (scr *Screen) Size() (w, h int) {
	return scr.w, scr.h
}

func (scr *Screen) SetCursor(x, y int) {
	if !(x < scr.x || y < scr.y || x > scr.w || y > scr.h) {
		termbox.SetCursor(x+scr.x, y+scr.y)
	} else {
		termbox.HideCursor()
	}
}
