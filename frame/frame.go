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

		return &Frame {ps, 0, 0, 0, 0, 0, s, true, true}, nil
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