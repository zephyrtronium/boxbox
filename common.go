package boxbox

import (
	"github.com/nsf/termbox-go"
)

func Blit(x, y, width int, cells []termbox.Cell) {
	tw, th := termbox.Size()
	height := len(cells) / width
	rw, rh := width, height
	offset := 0
	if x > tw || y > th {
		return
	}
	if x+width > tw {
		rw = tw - x
	}
	if x < 0 {
		rw += x
		offset = -x
		if offset > width {
			return
		}
	}
	if y+height > th {
		rh = th - y
	}
	if y < 0 {
		rh += y
		if rh <= 0 {
			return
		}
		y = 0
	}
	bb := termbox.CellBuffer()
	for i := 0; i < rh; i++ {
		copy(bb[(y+i)*tw+x+offset:], cells[i*width+offset:i*width+offset+rw])
	}
}

func Fill(x, y, width, height int, cell termbox.Cell) {
	tw, th := termbox.Size()
	if x > tw || y > th {
		return
	}
	if x < 0 {
		width += x
		x = 0
	}
	if x+width > tw {
		width = tw - x
		if width < 0 {
			return
		}
	}
	if y < 0 {
		height += y
		y = 0
	}
	if y+height > th {
		height = th - y
		if height < 0 {
			return
		}
	}
	cells := make([]termbox.Cell, width)
	for i := range cells {
		cells[i] = cell
	}
	bb := termbox.CellBuffer()
	for i := y*tw + x; i < (y+height)*tw+x; i += tw {
		copy(bb[i:], cells)
	}
}

func Highlight(x, y, width, height int, fg, bg *termbox.Attribute) {
	tw, th := termbox.Size()
	if x > tw || y > th || (fg == nil && bg == nil) {
		return
	}
	if x+width > tw {
		width = tw - x
		if width < 0 {
			return
		}
	}
	if y < 0 {
		height += y
		y = 0
	}
	if y+height > th {
		height = th - y
		if height < 0 {
			return
		}
	}
	bb := termbox.CellBuffer()
	if fg != nil {
		if bg != nil {
			for i := y*tw + x; i < (y+height)*tw+x; i += tw {
				for j := x; j < x+width; j++ {
					bb[i+j].Fg = *fg
					bb[i+j].Bg = *bg
				}
			}
		} else {
			for i := y*tw + x; i < (y+height)*tw+x; i += tw {
				for j := x; j < x+width; j++ {
					bb[i+j].Fg = *fg
				}
			}
		}
	} else {
		for i := y*tw + x; i < (y+height)*tw+x; i += tw {
			for j := x; j < x+width; j++ {
				bb[i+j].Bg = *bg
			}
		}
	}
}
