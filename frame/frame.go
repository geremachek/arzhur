package frame

import (
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

func NewFrame(starters ...string) (*Frame, error) {
	if s, err := tcell.NewScreen(); err == nil {
		ps := []portal.Portal{}

		for _, s := range starters {
			ps = append(ps, portal.NewPortal(s))
		}

		return &Frame {ps, 0, 0, 0, 0, 0, s, true, true}, nil
	} else {
		return nil, err
	}
}