package frame

import "github.com/gdamore/tcell"

// key presses result in a selection of portals, or not

type portalSelection int

const (
	None portalSelection = iota
	All
	Current
)

func (f *Frame) movePortal(dir int) {
	f.index += dir // change the index
	len := len(f.portals)-1

	// "loop back" if we are at the first or final window

	if f.index < 0 {
		f.index = len
	} else if f.index > len {
		f.index = 0
	}

	// update the screen

	f.scr.Clear()
	f.cursReset()

	f.drawPortalList()
	f.drawPortal()
}

// filter input to the editing buffer and general functions

func (f *Frame) filterInput(input *tcell.EventKey) portalSelection {
	key := input.Key()

	// escape toggles between editing and command mode

	if key == tcell.KeyESC {
		f.toggleEditing()
	} else if f.editing { // main editing mode
		switch key {
			case tcell.KeyDEL, tcell.KeyBackspace: f.delChar() // delete a character
			case tcell.KeyRune:                                  f.typeChar(input.Rune()) // insert a character
			case tcell.KeyEnter:                                  return Current // return the current line to stdout
			case tcell.KeyTAB:                                    f.typeChar('\t') // handle tabs
		}

		f.showCurs()
	} else if key == tcell.KeyRune { // move between buffers or quit
		switch ch := input.Rune(); ch {
			case 'q': f.running = false
			case 'h': f.movePortal(-1)
			case 'l': f.movePortal(1)
			case 'a': return All
		}
	}

	f.scr.Show()

	return None // by default, no portal(s) are selected
}

// draw the screen and listen for input

func (f *Frame) Start() (string, error) {
	if e := f.scr.Init(); e == nil {
		defer f.scr.Fini() // close the screen when we exit

		// handle input

		var input tcell.Event

		for f.running {
			input = f.scr.PollEvent()

			switch k := input.(type) { // the user either presses a key or the screen changes size
				case *tcell.EventKey:
					switch f.filterInput(k) { // return text, or not
						case None: continue
						case All: return f.returnAll(), nil
						case Current: return f.portals[f.index].String(), nil
					}
				case *tcell.EventResize: // this runs at start up as well.
					f.scr.Clear()

					f.width, f.height = k.Size()
					f.cursReset()

					f.drawFrame()
			}
		}
	} else {
		return "", e
	}

	return "", nil
}
