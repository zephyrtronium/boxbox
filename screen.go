package boxbox

import "github.com/nsf/termbox-go"

type Screen struct {
	x, y, w, h int
}

func NewScreen(x, y, w, h int) *Screen {
	return &Screen{x, y, w, h}
}

func (scr *Screen) Subscreen(x, y, w, h int) *Screen {
	return &Screen{x + scr.x, y + scr.y, w, h}
}

func (scr *Screen) PutCell(x, y int, cell *termbox.Cell) {
	termbox.PutCell(x + scr.x, y + scr.y, cell)
}

func (scr *Screen) ChangeCell(x, y int, ch rune, fg, bg termbox.Attribute) {
	termbox.PutCell(x + scr.x, y + scr.y, &termbox.Cell{ch, fg, bg})
}
