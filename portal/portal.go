package portal

// portal struct

type Portal struct {
	text []rune
}

// create a new portal text box

func NewPortal(text string) Portal {
	return Portal {[]rune(text)}
}

// return the text as a string

func (p Portal) String() string {
	return string(p.text)
}

// return the length of the text slice

func (p Portal) Length() int {
	return len(p.text)
}

// append a character to the text buffer

func (p *Portal) Insert(ch rune) {
	p.text = append(p.text, ch)
}

// Delete the final character within the text box.

func (p *Portal) Del() {
	if len := p.Length()-1; len > -1 {
		p.text = p.text[:len]
	}
}