package chimp

import (
	"bytes"
	"testing"
)

func TestParseWrite(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "Simple red bold",
			input: "[[Red,Bold]]text[[end]]",
			want:  "\033[31m\033[1mtext\033[0m",
		},
		{
			name:  "Red bold with surrounding text",
			input: "before [[Red,Bold]]middle[[end]] after",
			want:  "before \033[31m\033[1mmiddle\033[0m after",
		},
		{
			name:  "Multiple red bold tags",
			input: "[[Red,Bold]]first[[end]] and [[Red,Bold]]second[[end]]",
			want:  "\033[31m\033[1mfirst\033[0m and \033[31m\033[1msecond\033[0m",
		},
		{
			name:  "Plain text no tags",
			input: "just plain text",
			want:  "just plain text",
		},
		{
			name:  "Nested red then bold",
			input: "[[Red]][[Bold]]text[[end]][[end]]",
			want:  "\033[31m\033[1mtext\033[0m",
		},
		{
			name:  "Progressive styling with words",
			input: "plain [[Red]]redonly [[Bold]]redandbold[[end]] redonlyagain[[end]] plainagain",
			want:  "plain \033[31mredonly \033[31m\033[1mredandbold\033[31m redonlyagain\033[0m plainagain",
		},
		{
			name:  "Unknown style ignored",
			input: "[[Foo]]text[[end]]",
			want:  "text",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			c := New(&buf)
			n, err := c.Write([]byte(tt.input))
			if err != nil {
				t.Fatalf("Write failed: %v", err)
			}
			if n != len(tt.input) {
				t.Errorf("Expected %d bytes written, got %d", len(tt.input), n)
			}
			got := buf.String()
			if got != tt.want {
				t.Errorf("Chimp.Write(%q) wrote %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}
