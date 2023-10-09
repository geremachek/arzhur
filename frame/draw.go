package frame

import (
	"strconv"
	"github.com/gdamore/tcell"
)

// draw a line of text to the screen

func addString(s tcell.Screen, style tcell.Style, x int, y int, text string) {
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

// insert a string into the current window

func (f *Frame) typeString(text string) {
	for _, ch := range text {
		f.typeChar(ch)
	}
}

// remove a character from the screen

func (f *Frame) delChar() {
	if p := &f.portals[f.index]; p.Length() > 0 { // only delete when there is something to!
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

		label string
	)

	// write a text string at the bottom of the screen and move the cursor

	writeLabel := func(style tcell.Style, label string) {
		addString(f.scr, style, curs, f.height-1, label)
		curs += len(label)
	}

	for i, _ := range f.portals {
		if i == f.selected { // selected windows aren't numbered but marked with an "*"
			label = "*"
		} else {
			label = strconv.Itoa(i+1) // the current index acts as a label for the editing window
		}

		writeLabel(style, "[")

		if i == f.index {
			writeLabel(style.Reverse(true), label) // highlight the current window
		} else {
			writeLabel(style, label)
		}

		writeLabel(style, "] ")
	}
}

// draw the window text

func (f *Frame) drawPortal() {
	for _, ch := range f.portals[f.index].String() {
		f.drawAtCurs(ch)

		if f.x == f.width-1 && f.y == f.height-3 { // fit the text to the screen
			break
		}

		f.cursMove(true)
	}
}

// draw all of the UI elements

func (f *Frame) drawFrame() {
	f.scr.Clear()
	f.cursReset()

	f.drawPortalList()
	f.drawPortal()
}