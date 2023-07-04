package frame

import (
	"strconv"
	"github.com/gdamore/tcell"
)

func addstring(s tcell.Screen, style tcell.Style, x int, y int, text string) {
	for i, ch := range text {
		s.SetContent(i+x, y, ch, []rune{}, style)
	}
}

// draw a character on the screen

func (f *Frame) typeChar(ch rune) {
	f.portals[f.index].Insert(ch)

	f.drawAtCurs(ch)
	f.cursMove(true)
}

// remove a character from the screen

func (f *Frame) delChar() {
	if p := &f.portals[f.index]; p.Length() > 0 {
		p.Del()

		f.cursMove(false)
		f.drawAtCurs(' ')
	}
}

// draw a list of all of the open buffers at the bottom of the screen

func (f *Frame) drawPortalList() {
	var (
		style tcell.Style = tcell.StyleDefault
		curs int = 0

		converted string
	)

	// write a text string at the bottom of the screen

	writeLabel := func(style tcell.Style, label string) {
		addstring(f.scr, style, curs, f.height-1, label)
	}

	for i, _ := range f.portals {
		converted = strconv.Itoa(i+1) // the current index acts as a label for the editing window

		if i == f.index {
			writeLabel(style.Reverse(true), converted)
		} else {
			writeLabel(style, converted)
		}

		curs += len(converted)+1
	}
}

func (f *Frame) drawPortal() {
	for _, ch := range f.portals[f.index].Text() {
		f.drawAtCurs(ch)

		if f.x == f.width-1 && f.y == f.height-3 {
			break
		}

		f.cursMove(true)
	}
}

// draw all of the UI elements

func (f *Frame) drawFrame() {
	f.drawPortalList()
	f.drawPortal()

	f.showCurs()

	f.scr.Show()
}
