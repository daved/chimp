package chimp

import (
	"bytes"
	"testing"
)

func TestChimpSurface(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "Basic plain text",
			input: "hello world",
			want:  "hello world",
		},
		{
			name:  "Empty string",
			input: "",
			want:  "",
		},
		{
			name:  "Text with newline",
			input: "line1\nline2",
			want:  "line1\nline2",
		},
		{
			name:  "Simple red bold token",
			input: "[[Red,Bold]]text[[end]]",
			want:  "\033[31m\033[1mtext\033[0m",
		},
		{
			name:  "Text with red bold token",
			input: "before [[Red,Bold]]middle[[end]] after",
			want:  "before \033[31m\033[1mmiddle\033[0m after",
		},
		{
			name:  "Multiple red bold tokens",
			input: "[[Red,Bold]]first[[end]] and [[Red,Bold]]second[[end]]",
			want:  "\033[31m\033[1mfirst\033[0m and \033[31m\033[1msecond\033[0m",
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
				t.Errorf("Expected output %q, got %q", tt.want, got)
			}
		})
	}
}
