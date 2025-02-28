package chimp

import "strings"

// getANSIStyles suports standard SGR codes for colors and text attributes that
// don't require terminal response.
func getANSIStyles(styles string) string {
	ansi := ""
	for _, style := range strings.Split(styles, ",") {
		switch strings.TrimSpace(style) {
		// Reset
		case "Reset":
			ansi += "\033[0m"

		// Text Attributes
		case "Bold":
			ansi += "\033[1m"
		case "Faint":
			ansi += "\033[2m"
		case "Italic":
			ansi += "\033[3m"
		case "Underline":
			ansi += "\033[4m"
		case "Blink":
			ansi += "\033[5m" // Slow blink (less common, but supported)
		case "RapidBlink":
			ansi += "\033[6m" // Rapid blink (less common)
		case "Inverse":
			ansi += "\033[7m"
		case "Hidden":
			ansi += "\033[8m"
		case "Strikethrough":
			ansi += "\033[9m"

		// Foreground Colors (Standard)
		case "Black":
			ansi += "\033[30m"
		case "Red":
			ansi += "\033[31m"
		case "Green":
			ansi += "\033[32m"
		case "Yellow":
			ansi += "\033[33m"
		case "Blue":
			ansi += "\033[34m"
		case "Magenta":
			ansi += "\033[35m"
		case "Cyan":
			ansi += "\033[36m"
		case "White":
			ansi += "\033[37m"

		// Background Colors (Standard)
		case "BgBlack":
			ansi += "\033[40m"
		case "BgRed":
			ansi += "\033[41m"
		case "BgGreen":
			ansi += "\033[42m"
		case "BgYellow":
			ansi += "\033[43m"
		case "BgBlue":
			ansi += "\033[44m"
		case "BgMagenta":
			ansi += "\033[45m"
		case "BgCyan":
			ansi += "\033[46m"
		case "BgWhite":
			ansi += "\033[47m"

		// Bright Foreground Colors (Widely supported)
		case "BrightBlack":
			ansi += "\033[90m"
		case "BrightRed":
			ansi += "\033[91m"
		case "BrightGreen":
			ansi += "\033[92m"
		case "BrightYellow":
			ansi += "\033[93m"
		case "BrightBlue":
			ansi += "\033[94m"
		case "BrightMagenta":
			ansi += "\033[95m"
		case "BrightCyan":
			ansi += "\033[96m"
		case "BrightWhite":
			ansi += "\033[97m"

		// Bright Background Colors (Widely supported)
		case "BgBrightBlack":
			ansi += "\033[100m"
		case "BgBrightRed":
			ansi += "\033[101m"
		case "BgBrightGreen":
			ansi += "\033[102m"
		case "BgBrightYellow":
			ansi += "\033[103m"
		case "BgBrightBlue":
			ansi += "\033[104m"
		case "BgBrightMagenta":
			ansi += "\033[105m"
		case "BgBrightCyan":
			ansi += "\033[106m"
		case "BgBrightWhite":
			ansi += "\033[107m"
		}
	}
	return ansi
}
