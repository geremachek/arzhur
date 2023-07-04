package frame

import "github.com/gdamore/tcell"

func (f *Frame) movePortal(dir rune) {
	if dir == 'h' && f.index > 0 {
		f.index -= 1
	} else if dir == 'l' && f.index < len(f.portals)-1 {
		f.index += 1
	}

	f.drawFrame()
}*/

// filter input to the editing buffer and general functions

func (f *Frame) filterInput(input *tcell.EventKey) {
	key := input.Key()

	// escape toggles between editing and command mode

	if key == tcell.KeyESC {
		f.toggleEditing()
	} else if f.editing { // send keystrokes to the buffer, or delete characters
		switch key {
			case tcell.KeyDEL, tcell.KeyBackspace: f.delChar()
			case tcell.KeyRune:                                  f.typeChar(input.Rune())
		}

		f.showCurs()
	} else if key == tcell.KeyRune { // move between buffers
		switch ch := input.Rune(); ch {
			case 'q': f.running = false
			case 'h':
				if f.index > 0 {
					f.index -= 1
					f.drawFrame()
				}
			case 'l':
				if f.index < len(f.portals)-1 {
					f.index += 1
					f.drawFrame()
				}
		}
	}

	f.scr.Show()
}

// draw the screen and listen for input

func (f *Frame) Start() error {
	if e := f.scr.Init(); e == nil {
		//f.width, f.height = f.scr.Size()

		// draw UI elements

		//f.drawFrame()

		// handle input

		var input tcell.Event

		for f.running {
			input = f.scr.PollEvent()

			switch k := input.(type) {
				case *tcell.EventKey: f.filterInput(k)
				case *tcell.EventResize: 
					f.scr.Clear()

					f.width, f.height = k.Size()
					f.x, f.y = 0, 0

					f.drawFrame()
			}
		}

		f.scr.Fini()
	} else {
		return e
	}

	return nil
}
