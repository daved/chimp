package chimp

import (
	"fmt"
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
				newStyles, advance, continueParsing, err := splitStyles(data[i:], c.styles)
				if err != nil {
					return i, err
				}
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
				newStyles, advance, continueParsing, err := splitStyles(data[i:], c.styles)
				if err != nil {
					return i, err
				}
				c.styles = newStyles
				i += advance
				if !continueParsing && len(c.lastStyles) > 0 && len(c.styles) == 0 {
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
					if s := joinStyles(c.styles); s != "" {
						c.lastStyles = append([]string(nil), c.styles...)
					}
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

// splitStyles updates the styles stack based on parsed input.
func splitStyles(data []byte, styles []string) (newStyles []string, advance int, continueParsing bool, err error) {
	parsed, advance, continueParsing, err := parseStyle(data)
	if err != nil {
		return nil, 0, true, err
	}
	if !continueParsing {
		if parsed == "end" {
			if len(styles) > 0 {
				newStyles = styles[:len(styles)-1]
			} else {
				newStyles = styles
			}
			dbg("After [[end]], styles: %v\n", newStyles)
		} else {
			newStyles = append(styles, parsed)
			dbg("Processing style: %q\n", parsed)
		}
	}
	return newStyles, advance, continueParsing, nil
}

// writeStyles applies the current stack of ANSI styles to the writer.
func (c *Chimp) writeStyles() (n int, err error) {
	s := joinStyles(c.styles)
	if s == "" {
		return 0, nil
	}
	dbg("Writing styles: %q\n", s)
	return c.writer.Write([]byte(s))
}

// parseStyle extracts a style or end marker from input data.
func parseStyle(data []byte) (style string, advance int, continueParsing bool, err error) {
	if len(data) < 2 || data[0] != '[' || data[1] != '[' {
		return "", 0, true, nil
	}

	var buffer strings.Builder
	i := 2 // Skip [[
	for ; i < len(data); i++ {
		if i+1 < len(data) && data[i] == ']' && data[i+1] == ']' {
			return buffer.String(), i + 2, false, nil // Include ]]
		}
		if i+5 <= len(data) && string(data[i:i+5]) == "end]]" {
			return "end", i + 5, false, nil // Include [[end]]
		}
		if i == len(data)-1 {
			return "", 0, true, fmt.Errorf("unclosed style tag")
		}
		buffer.WriteByte(data[i])
	}
	return "", 0, true, nil // Need more data
}

// joinStyles combines a slice of styles into a single ANSI string.
func joinStyles(styles []string) string {
	var ansi strings.Builder
	for _, style := range styles {
		ansi.WriteString(getANSIStyles(style))
	}
	return ansi.String()
}

// getANSIStyles converts a comma-separated style string into ANSI escape sequences.
func getANSIStyles(styles string) string {
	ansi := ""
	for _, style := range strings.Split(styles, ",") {
		trimmed := strings.TrimSpace(style)
		if seq := Style(trimmed).ToSequence(); seq != SequenceUnknown {
			ansi += string(seq)
		}
	}
	return ansi
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
