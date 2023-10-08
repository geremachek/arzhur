package frame

import "github.com/gdamore/tcell"

// "smart" movement of the cursor

func (f *Frame) cursMove(forward bool) {
	if forward {
		if f.x == f.width-1 { // alllow line breaks
			f.x = 0
			f.y++
		} else {
			f.x++
		}
	} else {
		if f.x == 0 && f.y != 0 { // return line breaks
			f.x = f.width-1
			f.y--
		} else {
			f.x--
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
	// show the cursor only when we are in editing mode

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
	// display special characters

	switch ch {
		case '\n': ch = 'N'
		case '\t': ch = 'T'
	}

	f.scr.SetContent(f.x, f.y, ch, []rune{}, tcell.StyleDefault)
}
