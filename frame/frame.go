package frame

import (
	"strings"
	"github.com/gdamore/tcell"
	"github.com/geremachek/arzhur/portal"
)

// "Frame" that strings all of the "portals" together

type Frame struct {
	portals []portal.Portal

	index int
	selected int

	width int
	height int

	x int
	y int

	scr tcell.Screen
	running bool
	editing bool
}

// create a new frame

func NewFrame(starters []string) (*Frame, error) {
	if s, err := tcell.NewScreen(); err == nil {
		ps := []portal.Portal{}


		// convert each line of text into a "portal"

		for _, s := range starters {
			ps = append(ps, portal.NewPortal(s))
		}

		return &Frame {ps, 0, -1, 0, 0, 0, 0, s, true, true}, nil
	} else {
		return nil, err
	}
}

// return all of the windows as a single string together

func (f Frame) returnAll() string {
	var all strings.Builder

	for _, p := range f.portals {
		all.WriteString(p.String() + "\n") // connected by newlines
	}

	return all.String()
}

// return the text of the focused window

func (f Frame) focusedText() string {
	return f.portals[f.index].String()
}

// set the text of the focused window

func (f *Frame) setFocused(text string) {
	f.portals[f.index].Set(text)
}

// return the text of marked window

func (f Frame) markedText() string {
	return f.portals[f.selected].String()
}

// set the text of the marked window

func (f *Frame) setMarked(text string) {
	f.portals[f.selected].Set(text)
}

// open a new window

func (f *Frame) newWindow(text string) {
	f.portals = append(f.portals, portal.NewPortal(text))
	f.index = len(f.portals)-1
}

// close the most recently opened window

func (f *Frame) popWindow() {
	if len := len(f.portals); len > 1 { // you cannot close a the last window
		f.portals = f.portals[:len-1]

		// deselect if the final window is marked

		if len-1 == f.selected {
			f.selected = -1
		}

		// focus the new final window if we are deleting the former

		if f.index == len-1 {
			f.index--
		}
	}
}

// "mark" the current portal

func (f *Frame) markCurrent() {
	if f.index == f.selected {
		f.selected = -1 // deselected current
	} else {
		f.selected = f.index
	}
}

// return true if there is a marked window

func (f Frame) isMarked() bool {
	return f.selected > -1
}

// execute the text of the focused window a command

func (f Frame) execFocused(piping bool) (string, error) {
	input := ""

	if piping && f.isMarked() {
		input = f.markedText()
	}

	return runExternal(f.focusedText(), input, piping)
}