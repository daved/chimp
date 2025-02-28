package chimp

import "strings"

// Sequence represents an ANSI escape sequence.
type Sequence string

// Style represents an ANSI style name that maps to an escape sequence.
type Style string

// Matches checks if the input matches the Style exactly or case-insensitively with an offset of 32.
func (s Style) Matches(input string) bool {
	// Exact match
	if string(s) == input {
		return true
	}
	// Case-insensitive match: first char exact, rest offset by 32 (e.g., "Bold" matches "BOLD")
	if len(s) != len(input) || s[0] != input[0] {
		return false
	}
	for i := 1; i < len(s); i++ {
		if s[i] != input[i] && s[i]+32 != input[i] && s[i]-32 != input[i] {
			return false
		}
	}
	return true
}

// ANSI style names as exported Style constants and their corresponding Sequence constants.
const (
	// Reset
	StyleReset    Style    = "Reset"
	SequenceReset Sequence = "\033[0m"

	// Text Attributes
	StyleBold             Style    = "Bold"
	SequenceBold          Sequence = "\033[1m"
	StyleFaint            Style    = "Faint"
	SequenceFaint         Sequence = "\033[2m"
	StyleItalic           Style    = "Italic"
	SequenceItalic        Sequence = "\033[3m"
	StyleUnderline        Style    = "Underline"
	SequenceUnderline     Sequence = "\033[4m"
	StyleBlink            Style    = "Blink"
	SequenceBlink         Sequence = "\033[5m"
	StyleRapidBlink       Style    = "RapidBlink"
	SequenceRapidBlink    Sequence = "\033[6m"
	StyleInverse          Style    = "Inverse"
	SequenceInverse       Sequence = "\033[7m"
	StyleHidden           Style    = "Hidden"
	SequenceHidden        Sequence = "\033[8m"
	StyleStrikethrough    Style    = "Strikethrough"
	SequenceStrikethrough Sequence = "\033[9m"

	// Foreground Colors (Standard)
	StyleBlack      Style    = "Black"
	SequenceBlack   Sequence = "\033[30m"
	StyleRed        Style    = "Red"
	SequenceRed     Sequence = "\033[31m"
	StyleGreen      Style    = "Green"
	SequenceGreen   Sequence = "\033[32m"
	StyleYellow     Style    = "Yellow"
	SequenceYellow  Sequence = "\033[33m"
	StyleBlue       Style    = "Blue"
	SequenceBlue    Sequence = "\033[34m"
	StyleMagenta    Style    = "Magenta"
	SequenceMagenta Sequence = "\033[35m"
	StyleCyan       Style    = "Cyan"
	SequenceCyan    Sequence = "\033[36m"
	StyleWhite      Style    = "White"
	SequenceWhite   Sequence = "\033[37m"

	// Background Colors (Standard)
	StyleBgBlack      Style    = "BgBlack"
	SequenceBgBlack   Sequence = "\033[40m"
	StyleBgRed        Style    = "BgRed"
	SequenceBgRed     Sequence = "\033[41m"
	StyleBgGreen      Style    = "BgGreen"
	SequenceBgGreen   Sequence = "\033[42m"
	StyleBgYellow     Style    = "BgYellow"
	SequenceBgYellow  Sequence = "\033[43m"
	StyleBgBlue       Style    = "BgBlue"
	SequenceBgBlue    Sequence = "\033[44m"
	StyleBgMagenta    Style    = "BgMagenta"
	SequenceBgMagenta Sequence = "\033[45m"
	StyleBgCyan       Style    = "BgCyan"
	SequenceBgCyan    Sequence = "\033[46m"
	StyleBgWhite      Style    = "BgWhite"
	SequenceBgWhite   Sequence = "\033[47m"

	// Bright Foreground Colors
	StyleBrightBlack      Style    = "BrightBlack"
	SequenceBrightBlack   Sequence = "\033[90m"
	StyleBrightRed        Style    = "BrightRed"
	SequenceBrightRed     Sequence = "\033[91m"
	StyleBrightGreen      Style    = "BrightGreen"
	SequenceBrightGreen   Sequence = "\033[92m"
	StyleBrightYellow     Style    = "BrightYellow"
	SequenceBrightYellow  Sequence = "\033[93m"
	StyleBrightBlue       Style    = "BrightBlue"
	SequenceBrightBlue    Sequence = "\033[94m"
	StyleBrightMagenta    Style    = "BrightMagenta"
	SequenceBrightMagenta Sequence = "\033[95m"
	StyleBrightCyan       Style    = "BrightCyan"
	SequenceBrightCyan    Sequence = "\033[96m"
	StyleBrightWhite      Style    = "BrightWhite"
	SequenceBrightWhite   Sequence = "\033[97m"

	// Bright Background Colors
	StyleBgBrightBlack      Style    = "BgBrightBlack"
	SequenceBgBrightBlack   Sequence = "\033[100m"
	StyleBgBrightRed        Style    = "BgBrightRed"
	SequenceBgBrightRed     Sequence = "\033[101m"
	StyleBgBrightGreen      Style    = "BgBrightGreen"
	SequenceBgBrightGreen   Sequence = "\033[102m"
	StyleBgBrightYellow     Style    = "BgBrightYellow"
	SequenceBgBrightYellow  Sequence = "\033[103m"
	StyleBgBrightBlue       Style    = "BgBrightBlue"
	SequenceBgBrightBlue    Sequence = "\033[104m"
	StyleBgBrightMagenta    Style    = "BgBrightMagenta"
	SequenceBgBrightMagenta Sequence = "\033[105m"
	StyleBgBrightCyan       Style    = "BgBrightCyan"
	SequenceBgBrightCyan    Sequence = "\033[106m"
	StyleBgBrightWhite      Style    = "BgBrightWhite"
	SequenceBgBrightWhite   Sequence = "\033[107m"
)

// getANSIStyles maps style names to ANSI escape sequences using Style constants.
func getANSIStyles(styles string) string {
	ansi := ""
	for _, style := range strings.Split(styles, ",") {
		trimmed := strings.TrimSpace(style)
		switch {
		case StyleReset.Matches(trimmed):
			ansi += string(SequenceReset)
		case StyleBold.Matches(trimmed):
			ansi += string(SequenceBold)
		case StyleFaint.Matches(trimmed):
			ansi += string(SequenceFaint)
		case StyleItalic.Matches(trimmed):
			ansi += string(SequenceItalic)
		case StyleUnderline.Matches(trimmed):
			ansi += string(SequenceUnderline)
		case StyleBlink.Matches(trimmed):
			ansi += string(SequenceBlink)
		case StyleRapidBlink.Matches(trimmed):
			ansi += string(SequenceRapidBlink)
		case StyleInverse.Matches(trimmed):
			ansi += string(SequenceInverse)
		case StyleHidden.Matches(trimmed):
			ansi += string(SequenceHidden)
		case StyleStrikethrough.Matches(trimmed):
			ansi += string(SequenceStrikethrough)

		// Foreground Colors
		case StyleBlack.Matches(trimmed):
			ansi += string(SequenceBlack)
		case StyleRed.Matches(trimmed):
			ansi += string(SequenceRed)
		case StyleGreen.Matches(trimmed):
			ansi += string(SequenceGreen)
		case StyleYellow.Matches(trimmed):
			ansi += string(SequenceYellow)
		case StyleBlue.Matches(trimmed):
			ansi += string(SequenceBlue)
		case StyleMagenta.Matches(trimmed):
			ansi += string(SequenceMagenta)
		case StyleCyan.Matches(trimmed):
			ansi += string(SequenceCyan)
		case StyleWhite.Matches(trimmed):
			ansi += string(SequenceWhite)

		// Background Colors
		case StyleBgBlack.Matches(trimmed):
			ansi += string(SequenceBgBlack)
		case StyleBgRed.Matches(trimmed):
			ansi += string(SequenceBgRed)
		case StyleBgGreen.Matches(trimmed):
			ansi += string(SequenceBgGreen)
		case StyleBgYellow.Matches(trimmed):
			ansi += string(SequenceBgYellow)
		case StyleBgBlue.Matches(trimmed):
			ansi += string(SequenceBgBlue)
		case StyleBgMagenta.Matches(trimmed):
			ansi += string(SequenceBgMagenta)
		case StyleBgCyan.Matches(trimmed):
			ansi += string(SequenceBgCyan)
		case StyleBgWhite.Matches(trimmed):
			ansi += string(SequenceBgWhite)

		// Bright Foreground Colors
		case StyleBrightBlack.Matches(trimmed):
			ansi += string(SequenceBrightBlack)
		case StyleBrightRed.Matches(trimmed):
			ansi += string(SequenceBrightRed)
		case StyleBrightGreen.Matches(trimmed):
			ansi += string(SequenceBrightGreen)
		case StyleBrightYellow.Matches(trimmed):
			ansi += string(SequenceBrightYellow)
		case StyleBrightBlue.Matches(trimmed):
			ansi += string(SequenceBrightBlue)
		case StyleBrightMagenta.Matches(trimmed):
			ansi += string(SequenceBrightMagenta)
		case StyleBrightCyan.Matches(trimmed):
			ansi += string(SequenceBrightCyan)
		case StyleBrightWhite.Matches(trimmed):
			ansi += string(SequenceBrightWhite)

		// Bright Background Colors
		case StyleBgBrightBlack.Matches(trimmed):
			ansi += string(SequenceBgBrightBlack)
		case StyleBgBrightRed.Matches(trimmed):
			ansi += string(SequenceBgBrightRed)
		case StyleBgBrightGreen.Matches(trimmed):
			ansi += string(SequenceBgBrightGreen)
		case StyleBgBrightYellow.Matches(trimmed):
			ansi += string(SequenceBgBrightYellow)
		case StyleBgBrightBlue.Matches(trimmed):
			ansi += string(SequenceBgBrightBlue)
		case StyleBgBrightMagenta.Matches(trimmed):
			ansi += string(SequenceBgBrightMagenta)
		case StyleBgBrightCyan.Matches(trimmed):
			ansi += string(SequenceBgBrightCyan)
		case StyleBgBrightWhite.Matches(trimmed):
			ansi += string(SequenceBgBrightWhite)
		}
	}
	return ansi
}
