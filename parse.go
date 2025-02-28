package chimp

import (
	"io"
	"strings"
)

type Parse struct {
	writer     io.Writer
	styles     []string
	lastStyles []string
	state      string
}

func NewParse(w io.Writer) *Parse {
	return &Parse{
		writer: w,
		state:  "normal",
	}
}

func (p *Parse) Write(data []byte) (n int, err error) {
	for i := 0; i < len(data); {
		switch p.state {
		case "normal":
			if i+1 < len(data) && data[i] == '[' && data[i+1] == '[' {
				newStyles, advance, continueParsing := splitStyles(data[i:], p.styles)
				p.styles = newStyles
				i += advance
				if !continueParsing {
					p.state = "content"
				}
			} else {
				_, err = p.writer.Write([]byte{data[i]})
				if err != nil {
					return i, err
				}
				i++
			}
		case "content":
			if i+1 < len(data) && data[i] == '[' && data[i+1] == '[' {
				newStyles, advance, continueParsing := splitStyles(data[i:], p.styles)
				p.styles = newStyles
				i += advance
				if !continueParsing && len(p.styles) == 0 && len(p.lastStyles) > 0 {
					_, err = p.writer.Write([]byte("\033[0m"))
					if err != nil {
						return i, err
					}
					p.lastStyles = nil
				}
			} else {
				if !equalStyles(p.styles, p.lastStyles) && len(p.styles) > 0 {
					_, err = p.writeStyles()
					if err != nil {
						return i, err
					}
					p.lastStyles = append([]string(nil), p.styles...)
				}
				_, err = p.writer.Write([]byte{data[i]})
				if err != nil {
					return i, err
				}
				i++
			}
		}
	}
	return len(data), nil
}

func splitStyles(data []byte, styles []string) (newStyles []string, advance int, continueParsing bool) {
	if len(data) < 2 || data[0] != '[' || data[1] != '[' {
		return styles, 0, true
	}

	var buffer strings.Builder
	i := 2 // Skip [[
	for ; i < len(data); i++ {
		if i+1 < len(data) && data[i] == ']' && data[i+1] == ']' {
			newStyles = append(styles, buffer.String())
			dbg("Processing style: %q\n", buffer.String())
			return newStyles, i + 2, false // Include ]]
		}
		// Check for [[end]] safely
		if i+5 <= len(data) && string(data[i:i+5]) == "end]]" {
			if len(styles) > 0 {
				newStyles = styles[:len(styles)-1]
			} else {
				newStyles = styles
			}
			dbg("After [[end]], styles: %v\n", newStyles)
			return newStyles, i + 5, false // Include [[end]]
		}
		buffer.WriteByte(data[i])
	}
	return styles, 0, true // Need more data
}

func (p *Parse) writeStyles() (n int, err error) {
	var ansi strings.Builder
	for _, style := range p.styles {
		ansi.WriteString(getANSIStyles(style))
	}
	s := ansi.String()
	dbg("Writing styles: %q\n", s)
	return p.writer.Write([]byte(s))
}

func equalStyles(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
