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

	f.drawFrame()
}

// create a new window with the output of a command

func (f *Frame) winFromCommand(p bool) {
	if out, err := f.execFocused(p); err == nil {
		f.newWindow(out)
		f.drawFrame()
	} // if there is an error, don't create a window
}

// replace the selection or focused window with the output of a command

func (f *Frame) replaceWithOutput(p bool) {
	if out, err := f.execFocused(p); err == nil {
		// set the marked or focused text

		if f.isMarked() {
			f.setMarked(out)
			f.index = f.selected
		} else {
			f.setFocused(out)
		}

		f.x = 0
		f.drawFrame()
	}
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
			case tcell.KeyRune:                    f.typeChar(input.Rune()) // insert a character
			case tcell.KeyEnter:                   return Current // return the current line to stdout
			case tcell.KeyTAB:                     f.typeChar('\t') // handle tabs
		}

		f.showCurs()
	} else if key == tcell.KeyRune { // parse commands
		switch ch := input.Rune(); ch {
			case 'q': f.running = false
			case 'h': f.movePortal(-1)
			case 'l': f.movePortal(1)
			case 'a': return All
			case 'n': // create a new, empty window
				f.newWindow("")
				f.drawFrame()
			case 'd': // delete the last window
				f.popWindow()
				f.drawFrame()
			case 'm': // mark/unmark the current window
				f.markCurrent()
				f.drawPortalList()
			case 'c': // clear the focused text
				f.setFocused("")
				f.drawFrame()
			case '!': f.winFromCommand(false) // open a new window for command output
			case '<': f.replaceWithOutput(false) // replace the current window's/mark's text with output
			case '>': f.winFromCommand(true) // pipe the marked text and create a new window
			case '|': f.replaceWithOutput(true) // replace the marked text with piped text
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
						case Current: return f.focusedText(), nil
					}
				case *tcell.EventResize: // this runs at start up as well.
					f.width, f.height = k.Size()

					f.scr.Clear()

					f.drawFrame()

					// only show the cursor in editing mode

					if f.editing {
						f.showCurs()
					}

					f.scr.Show()
			}
		}
	} else {
		return "", e
	}

	return "", nil
}
