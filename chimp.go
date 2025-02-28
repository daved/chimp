package chimp

import (
	"io"
	"strings"
)

// Chimp processes text incrementally, applying ANSI styles with nesting.
type Chimp struct {
	writer     io.Writer
	styles     []string
	lastStyles []string
	state      string
}

// New creates a new Chimp with the given writer.
func New(w io.Writer) *Chimp {
	return &Chimp{
		writer: w,
		state:  "normal",
	}
}

// Write processes input bytes, updating state and writing styled output.
func (c *Chimp) Write(data []byte) (n int, err error) {
	for i := 0; i < len(data); {
		switch c.state {
		case "normal":
			if i+1 < len(data) && data[i] == '[' && data[i+1] == '[' {
				newStyles, advance, continueParsing := splitStyles(data[i:], c.styles)
				c.styles = newStyles
				i += advance
				if !continueParsing {
					c.state = "content"
				}
			} else {
				_, err = c.writer.Write([]byte{data[i]})
				if err != nil {
					return i, err
				}
				i++
			}
		case "content":
			if i+1 < len(data) && data[i] == '[' && data[i+1] == '[' {
				newStyles, advance, continueParsing := splitStyles(data[i:], c.styles)
				c.styles = newStyles
				i += advance
				if !continueParsing && len(c.styles) == 0 && len(c.lastStyles) > 0 {
					_, err = c.writer.Write([]byte("\033[0m"))
					if err != nil {
						return i, err
					}
					c.lastStyles = nil
				}
			} else {
				if !equalStyles(c.styles, c.lastStyles) && len(c.styles) > 0 {
					_, err = c.writeStyles()
					if err != nil {
						return i, err
					}
					c.lastStyles = append([]string(nil), c.styles...)
				}
				_, err = c.writer.Write([]byte{data[i]})
				if err != nil {
					return i, err
				}
				i++
			}
		}
	}
	return len(data), nil
}

// splitStyles processes input to update the styles stack.
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

// writeStyles applies the current stack of ANSI styles to the writer.
func (c *Chimp) writeStyles() (n int, err error) {
	var ansi strings.Builder
	for _, style := range c.styles {
		ansi.WriteString(getANSIStyles(style))
	}
	s := ansi.String()
	dbg("Writing styles: %q\n", s)
	return c.writer.Write([]byte(s))
}

// equalStyles compares two style slices for equality.
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
