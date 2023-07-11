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
	} else if f.editing { // send keystrokes to the buffer, or delete characters
		switch key {
			case tcell.KeyDEL, tcell.KeyBackspace: f.delChar()
			case tcell.KeyRune:                                  f.typeChar(input.Rune())
			case tcell.KeyEnter:                                  return Current
			case tcell.KeyTAB:                                    f.typeChar('\t')
		}

		f.showCurs()
	} else if key == tcell.KeyRune { // move between buffers
		switch ch := input.Rune(); ch {
			case 'q': f.running = false
			case 'h': f.movePortal(-1)
			case 'l': f.movePortal(1)
			case 'a': return All
		}
	}

	f.scr.Show()

	return None
}

// draw the screen and listen for input

func (f *Frame) Start() (string, error) {
	if e := f.scr.Init(); e == nil {
		defer f.scr.Fini() // close the screen when we exit

		// handle input

		var input tcell.Event

		for f.running {
			input = f.scr.PollEvent()

			switch k := input.(type) {
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
