package frame

import "github.com/gdamore/tcell"

// "smart" movement of the cursor

func (f *Frame) cursMove(forward bool) {
	if forward {
		if f.x == f.width-1 { // alllow line breaks
			f.x = 0
			f.y += 1
		} else {
			f.x += 1
		}
	} else {
		if f.x == 0 && f.y != 0 { // return line breaks
			f.x = f.width-1
			f.y -= 1
		} else {
			f.x -= 1
		}
	}
}

// reset the cursor position

func (f *Frame) cursReset() {
	f.x, f.y = 0, 0
}

// show the cursor at the current position

func (f *Frame) showCurs() {
	f.scr.ShowCursor(f.x, f.y)
}

// toggle the cursor and input mode

func (f *Frame) toggleEditing() {
	if f.editing {
		f.editing = false
		f.scr.HideCursor()
	} else {
		f.editing = true
		f.showCurs()
	}
}

// draw a character at the current cursor position

func (f *Frame) drawAtCurs(ch rune) {
	f.scr.SetContent(f.x, f.y, ch, []rune{}, tcell.StyleDefault)
}
