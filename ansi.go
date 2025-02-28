package chimp

import "strings"

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

// ANSI style names as exported Style constants.
const (
	StyleReset         Style = "Reset"
	StyleBold          Style = "Bold"
	StyleFaint         Style = "Faint"
	StyleItalic        Style = "Italic"
	StyleUnderline     Style = "Underline"
	StyleBlink         Style = "Blink"
	StyleRapidBlink    Style = "RapidBlink"
	StyleInverse       Style = "Inverse"
	StyleHidden        Style = "Hidden"
	StyleStrikethrough Style = "Strikethrough"

	// Foreground Colors (Standard)
	StyleBlack   Style = "Black"
	StyleRed     Style = "Red"
	StyleGreen   Style = "Green"
	StyleYellow  Style = "Yellow"
	StyleBlue    Style = "Blue"
	StyleMagenta Style = "Magenta"
	StyleCyan    Style = "Cyan"
	StyleWhite   Style = "White"

	// Background Colors (Standard)
	StyleBgBlack   Style = "BgBlack"
	StyleBgRed     Style = "BgRed"
	StyleBgGreen   Style = "BgGreen"
	StyleBgYellow  Style = "BgYellow"
	StyleBgBlue    Style = "BgBlue"
	StyleBgMagenta Style = "BgMagenta"
	StyleBgCyan    Style = "BgCyan"
	StyleBgWhite   Style = "BgWhite"

	// Bright Foreground Colors
	StyleBrightBlack   Style = "BrightBlack"
	StyleBrightRed     Style = "BrightRed"
	StyleBrightGreen   Style = "BrightGreen"
	StyleBrightYellow  Style = "BrightYellow"
	StyleBrightBlue    Style = "BrightBlue"
	StyleBrightMagenta Style = "BrightMagenta"
	StyleBrightCyan    Style = "BrightCyan"
	StyleBrightWhite   Style = "BrightWhite"

	// Bright Background Colors
	StyleBgBrightBlack   Style = "BgBrightBlack"
	StyleBgBrightRed     Style = "BgBrightRed"
	StyleBgBrightGreen   Style = "BgBrightGreen"
	StyleBgBrightYellow  Style = "BgBrightYellow"
	StyleBgBrightBlue    Style = "BgBrightBlue"
	StyleBgBrightMagenta Style = "BgBrightMagenta"
	StyleBgBrightCyan    Style = "BgBrightCyan"
	StyleBgBrightWhite   Style = "BgBrightWhite"
)

// getANSIStyles maps style names to ANSI escape sequences using Style constants.
func getANSIStyles(styles string) string {
	ansi := ""
	for _, style := range strings.Split(styles, ",") {
		trimmed := strings.TrimSpace(style)
		switch {
		case StyleReset.Matches(trimmed):
			ansi += "\033[0m"
		case StyleBold.Matches(trimmed):
			ansi += "\033[1m"
		case StyleFaint.Matches(trimmed):
			ansi += "\033[2m"
		case StyleItalic.Matches(trimmed):
			ansi += "\033[3m"
		case StyleUnderline.Matches(trimmed):
			ansi += "\033[4m"
		case StyleBlink.Matches(trimmed):
			ansi += "\033[5m"
		case StyleRapidBlink.Matches(trimmed):
			ansi += "\033[6m"
		case StyleInverse.Matches(trimmed):
			ansi += "\033[7m"
		case StyleHidden.Matches(trimmed):
			ansi += "\033[8m"
		case StyleStrikethrough.Matches(trimmed):
			ansi += "\033[9m"

		// Foreground Colors
		case StyleBlack.Matches(trimmed):
			ansi += "\033[30m"
		case StyleRed.Matches(trimmed):
			ansi += "\033[31m"
		case StyleGreen.Matches(trimmed):
			ansi += "\033[32m"
		case StyleYellow.Matches(trimmed):
			ansi += "\033[33m"
		case StyleBlue.Matches(trimmed):
			ansi += "\033[34m"
		case StyleMagenta.Matches(trimmed):
			ansi += "\033[35m"
		case StyleCyan.Matches(trimmed):
			ansi += "\033[36m"
		case StyleWhite.Matches(trimmed):
			ansi += "\033[37m"

		// Background Colors
		case StyleBgBlack.Matches(trimmed):
			ansi += "\033[40m"
		case StyleBgRed.Matches(trimmed):
			ansi += "\033[41m"
		case StyleBgGreen.Matches(trimmed):
			ansi += "\033[42m"
		case StyleBgYellow.Matches(trimmed):
			ansi += "\033[43m"
		case StyleBgBlue.Matches(trimmed):
			ansi += "\033[44m"
		case StyleBgMagenta.Matches(trimmed):
			ansi += "\033[45m"
		case StyleBgCyan.Matches(trimmed):
			ansi += "\033[46m"
		case StyleBgWhite.Matches(trimmed):
			ansi += "\033[47m"

		// Bright Foreground Colors
		case StyleBrightBlack.Matches(trimmed):
			ansi += "\033[90m"
		case StyleBrightRed.Matches(trimmed):
			ansi += "\033[91m"
		case StyleBrightGreen.Matches(trimmed):
			ansi += "\033[92m"
		case StyleBrightYellow.Matches(trimmed):
			ansi += "\033[93m"
		case StyleBrightBlue.Matches(trimmed):
			ansi += "\033[94m"
		case StyleBrightMagenta.Matches(trimmed):
			ansi += "\033[95m"
		case StyleBrightCyan.Matches(trimmed):
			ansi += "\033[96m"
		case StyleBrightWhite.Matches(trimmed):
			ansi += "\033[97m"

		// Bright Background Colors
		case StyleBgBrightBlack.Matches(trimmed):
			ansi += "\033[100m"
		case StyleBgBrightRed.Matches(trimmed):
			ansi += "\033[101m"
		case StyleBgBrightGreen.Matches(trimmed):
			ansi += "\033[102m"
		case StyleBgBrightYellow.Matches(trimmed):
			ansi += "\033[103m"
		case StyleBgBrightBlue.Matches(trimmed):
			ansi += "\033[104m"
		case StyleBgBrightMagenta.Matches(trimmed):
			ansi += "\033[105m"
		case StyleBgBrightCyan.Matches(trimmed):
			ansi += "\033[106m"
		case StyleBgBrightWhite.Matches(trimmed):
			ansi += "\033[107m"
		}
	}
	return ansi
}
