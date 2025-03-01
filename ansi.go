package chimp

import "strings"

// Sequence represents an ANSI escape sequence.
type Sequence string

// makeSequence creates a Sequence from a string using style-to-sequence conversion.
func makeSequence(s string) Sequence {
	return styleToSequence(Style(s))
}

// ToStyle converts a Sequence to its corresponding Style.
func (s Sequence) ToStyle() Style {
	return sequenceToStyle(s)
}

// Style represents an ANSI style name that maps to an escape sequence.
type Style string

// makeStyle creates a Style from a string using sequence-to-style conversion.
func makeStyle(s string) Style {
	return sequenceToStyle(Sequence(s))
}

// ApplyStyles converts a list of style names to an ANSI sequence string.
// Unknown styles are ignored.
func ApplyStyles(styles ...string) string {
	var ansi strings.Builder
	for _, style := range styles {
		if seq := Style(style).ToSequence(); seq != SequenceUnknown {
			ansi.WriteString(string(seq))
		}
	}
	return ansi.String()
}

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

// ToSequence converts a Style to its corresponding Sequence.
func (s Style) ToSequence() Sequence {
	return styleToSequence(s)
}

// ANSI style names as exported Style constants and their corresponding Sequence constants.
const (
	// Special Cases
	StyleUnknown    Style    = "unknown"
	SequenceUnknown Sequence = "unknown"
	StyleUnset      Style    = ""
	SequenceUnset   Sequence = ""

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

// sequenceToStyle converts a Sequence to its corresponding Style.
func sequenceToStyle(s Sequence) Style {
	switch s {
	case SequenceReset:
		return StyleReset
	case SequenceBold:
		return StyleBold
	case SequenceFaint:
		return StyleFaint
	case SequenceItalic:
		return StyleItalic
	case SequenceUnderline:
		return StyleUnderline
	case SequenceBlink:
		return StyleBlink
	case SequenceRapidBlink:
		return StyleRapidBlink
	case SequenceInverse:
		return StyleInverse
	case SequenceHidden:
		return StyleHidden
	case SequenceStrikethrough:
		return StyleStrikethrough

	// Foreground Colors
	case SequenceBlack:
		return StyleBlack
	case SequenceRed:
		return StyleRed
	case SequenceGreen:
		return StyleGreen
	case SequenceYellow:
		return StyleYellow
	case SequenceBlue:
		return StyleBlue
	case SequenceMagenta:
		return StyleMagenta
	case SequenceCyan:
		return StyleCyan
	case SequenceWhite:
		return StyleWhite

	// Background Colors
	case SequenceBgBlack:
		return StyleBgBlack
	case SequenceBgRed:
		return StyleBgRed
	case SequenceBgGreen:
		return StyleBgGreen
	case SequenceBgYellow:
		return StyleBgYellow
	case SequenceBgBlue:
		return StyleBgBlue
	case SequenceBgMagenta:
		return StyleBgMagenta
	case SequenceBgCyan:
		return StyleBgCyan
	case SequenceBgWhite:
		return StyleBgWhite

	// Bright Foreground Colors
	case SequenceBrightBlack:
		return StyleBrightBlack
	case SequenceBrightRed:
		return StyleBrightRed
	case SequenceBrightGreen:
		return StyleBrightGreen
	case SequenceBrightYellow:
		return StyleBrightYellow
	case SequenceBrightBlue:
		return StyleBrightBlue
	case SequenceBrightMagenta:
		return StyleBrightMagenta
	case SequenceBrightCyan:
		return StyleBrightCyan
	case SequenceBrightWhite:
		return StyleBrightWhite

	// Bright Background Colors
	case SequenceBgBrightBlack:
		return StyleBgBrightBlack
	case SequenceBgBrightRed:
		return StyleBgBrightRed
	case SequenceBgBrightGreen:
		return StyleBgBrightGreen
	case SequenceBgBrightYellow:
		return StyleBgBrightYellow
	case SequenceBgBrightBlue:
		return StyleBgBrightBlue
	case SequenceBgBrightMagenta:
		return StyleBgBrightMagenta
	case SequenceBgBrightCyan:
		return StyleBgBrightCyan
	case SequenceBgBrightWhite:
		return StyleBgBrightWhite

	// Special Cases
	case SequenceUnknown:
		return StyleUnknown
	case SequenceUnset:
		return StyleUnset
	}
	return StyleUnknown // Default for unrecognized sequences
}

// styleToSequence converts a Style to its corresponding Sequence.
func styleToSequence(s Style) Sequence {
	switch {
	case StyleReset.Matches(string(s)):
		return SequenceReset
	case StyleBold.Matches(string(s)):
		return SequenceBold
	case StyleFaint.Matches(string(s)):
		return SequenceFaint
	case StyleItalic.Matches(string(s)):
		return SequenceItalic
	case StyleUnderline.Matches(string(s)):
		return SequenceUnderline
	case StyleBlink.Matches(string(s)):
		return SequenceBlink
	case StyleRapidBlink.Matches(string(s)):
		return SequenceRapidBlink
	case StyleInverse.Matches(string(s)):
		return SequenceInverse
	case StyleHidden.Matches(string(s)):
		return SequenceHidden
	case StyleStrikethrough.Matches(string(s)):
		return SequenceStrikethrough

	// Foreground Colors
	case StyleBlack.Matches(string(s)):
		return SequenceBlack
	case StyleRed.Matches(string(s)):
		return SequenceRed
	case StyleGreen.Matches(string(s)):
		return SequenceGreen
	case StyleYellow.Matches(string(s)):
		return SequenceYellow
	case StyleBlue.Matches(string(s)):
		return SequenceBlue
	case StyleMagenta.Matches(string(s)):
		return SequenceMagenta
	case StyleCyan.Matches(string(s)):
		return SequenceCyan
	case StyleWhite.Matches(string(s)):
		return SequenceWhite

	// Background Colors
	case StyleBgBlack.Matches(string(s)):
		return SequenceBgBlack
	case StyleBgRed.Matches(string(s)):
		return SequenceBgRed
	case StyleBgGreen.Matches(string(s)):
		return SequenceBgGreen
	case StyleBgYellow.Matches(string(s)):
		return SequenceBgYellow
	case StyleBgBlue.Matches(string(s)):
		return SequenceBgBlue
	case StyleBgMagenta.Matches(string(s)):
		return SequenceBgMagenta
	case StyleBgCyan.Matches(string(s)):
		return SequenceBgCyan
	case StyleBgWhite.Matches(string(s)):
		return SequenceBgWhite

	// Bright Foreground Colors
	case StyleBrightBlack.Matches(string(s)):
		return SequenceBrightBlack
	case StyleBrightRed.Matches(string(s)):
		return SequenceBrightRed
	case StyleBrightGreen.Matches(string(s)):
		return SequenceBrightGreen
	case StyleBrightYellow.Matches(string(s)):
		return SequenceBrightYellow
	case StyleBrightBlue.Matches(string(s)):
		return SequenceBrightBlue
	case StyleBrightMagenta.Matches(string(s)):
		return SequenceBrightMagenta
	case StyleBrightCyan.Matches(string(s)):
		return SequenceBrightCyan
	case StyleBrightWhite.Matches(string(s)):
		return SequenceBrightWhite

	// Bright Background Colors
	case StyleBgBrightBlack.Matches(string(s)):
		return SequenceBgBrightBlack
	case StyleBgBrightRed.Matches(string(s)):
		return SequenceBgBrightRed
	case StyleBgBrightGreen.Matches(string(s)):
		return SequenceBgBrightGreen
	case StyleBgBrightYellow.Matches(string(s)):
		return SequenceBgBrightYellow
	case StyleBgBrightBlue.Matches(string(s)):
		return SequenceBgBrightBlue
	case StyleBgBrightMagenta.Matches(string(s)):
		return SequenceBgBrightMagenta
	case StyleBgBrightCyan.Matches(string(s)):
		return SequenceBgBrightCyan
	case StyleBgBrightWhite.Matches(string(s)):
		return SequenceBgBrightWhite

	// Special Cases
	case StyleUnknown.Matches(string(s)):
		return SequenceUnknown
	case StyleUnset.Matches(string(s)):
		return SequenceUnset
	}
	return SequenceUnknown // Default for unrecognized styles
}
