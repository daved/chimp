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
// It returns the number of bytes successfully written to the underlying writer.
func (c *Chimp) Write(data []byte) (n int, err error) {
	written := 0
	for i := 0; i < len(data); {
		if i+1 < len(data) && data[i] == '[' && data[i+1] == '[' {
			advance, n, err := handleStyleTag(c.writer, data[i:], &c.styles, &c.lastStyles)
			if err != nil {
				return written, err
			}
			written += n
			i += advance
			if c.state == "normal" {
				c.state = "content"
			}
		} else {
			if c.state == "content" && !stylesTextsMatch(c.styles, c.lastStyles) {
				n, err := applyStyleChanges(c.writer, c.styles, &c.lastStyles)
				if err != nil {
					return written, err
				}
				written += n
			}
			n, err := c.writer.Write([]byte{data[i]})
			if err != nil {
				return written, err
			}
			written += n
			i++
		}
	}
	return written, nil
}

// handleStyleTag parses a style tag and applies changes, returning bytes advanced and written.
func handleStyleTag(w io.Writer, data []byte, styles, lastStyles *[]string) (advance, n int, err error) {
	newStyles, advance, continueParsing, err := splitStyles(data, *styles)
	if err != nil {
		return 0, 0, err
	}
	*styles = newStyles
	if !continueParsing {
		n, err := applyStyleChanges(w, *styles, lastStyles)
		if err != nil {
			return advance, 0, err
		}
		return advance, n, nil
	}
	return advance, 0, nil
}

// applyStyleChanges writes styles or resets if changed, updating lastStyles.
func applyStyleChanges(w io.Writer, styles []string, lastStyles *[]string) (n int, err error) {
	if !stylesTextsMatch(styles, *lastStyles) {
		s := joinStylesTexts(styles)
		if s != "" {
			dbg("Writing styles: %q\n", s)
			n, err = w.Write([]byte(s))
			if err != nil {
				return 0, err
			}
			*lastStyles = append([]string(nil), styles...)
			return n, nil
		} else if len(*lastStyles) > 0 {
			n, err = w.Write([]byte("\033[0m"))
			if err != nil {
				return 0, err
			}
			*lastStyles = nil
			return n, nil
		}
	}
	return 0, nil
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

// joinStylesTexts combines a slice of styles into a single ANSI string.
func joinStylesTexts(styles []string) string {
	var ansi strings.Builder
	for _, style := range styles {
		ansi.WriteString(stylesTextToSequencesText(style))
	}
	return ansi.String()
}

// stylesTextToSequencesText converts a comma-separated style string into ANSI escape sequences.
func stylesTextToSequencesText(styles string) string {
	ansi := ""
	for _, style := range strings.Split(styles, ",") {
		trimmed := strings.TrimSpace(style)
		if seq := Style(trimmed).ToSequence(); seq != SequenceUnknown {
			ansi += string(seq)
		}
	}
	return ansi
}

// stylesTextsMatch compares two style slices for equality.
func stylesTextsMatch(a, b []string) bool {
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
