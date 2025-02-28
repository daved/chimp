package chimp

import (
	"fmt"
	"os"
)

var dbg = func() func(string, ...any) {
	if os.Getenv("DEBUG") == "" {
		return func(string, ...any) {}
	}
	return func(format string, as ...any) {
		fmt.Printf(format, as...)
	}
}()
