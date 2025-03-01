package chimp

import (
	"bytes"
	"errors"
	"io"
	"testing"
)

// Existing TestParseWrite remains unchanged...

// TestWriteEdgeCases tests Surface scope edge cases for Chimp.Write.
func TestWriteEdgeCases(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantN   int
		wantErr bool
	}{
		{
			name:    "Incomplete tag at end",
			input:   "[[Red",
			want:    "",
			wantN:   0,
			wantErr: true,
		},
		{
			name:    "Rapid style switches",
			input:   "[[Red]][[end]][[Bold]]text[[end]]",
			want:    "\033[31m\033[0m\033[1mtext\033[0m",
			wantN:   len("\033[31m\033[0m\033[1mtext\033[0m"), // 4 + 2 + 4 + 4 + 2 = 16
			wantErr: false,
		},
		{
			name:    "Nested with incomplete inner",
			input:   "[[Red]][[Bold",
			want:    "\033[31m",
			wantN:   len("\033[31m"), // 4
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			c := New(&buf)
			n, err := c.Write([]byte(tt.input))
			if (err != nil) != tt.wantErr {
				t.Errorf("Write() error = %v, wantErr %v", err, tt.wantErr)
			}
			if n != tt.wantN {
				t.Errorf("Write() wrote %d bytes, want %d", n, tt.wantN)
			}
			got := buf.String()
			if got != tt.want {
				t.Errorf("Write(%q) wrote %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

// TestWriterFailure tests Surface scope with a failing writer.
func TestWriterFailure(t *testing.T) {
	c := New(failingWriter{})
	_, err := c.Write([]byte("[[Red]]text"))
	if err == nil {
		t.Errorf("Write() with failing writer should return an error")
	}
	if !errors.Is(err, io.ErrClosedPipe) {
		t.Errorf("Write() error = %v, want %v", err, io.ErrClosedPipe)
	}
}

// failingWriter always fails with io.ErrClosedPipe.
type failingWriter struct{}

func (failingWriter) Write(p []byte) (n int, err error) {
	return 0, io.ErrClosedPipe
}

// TestSplitStyles tests Unit scope for splitStyles.
func TestSplitStyles(t *testing.T) {
	tests := []struct {
		name         string
		data         []byte
		styles       []string
		wantStyles   []string
		wantAdvance  int
		wantContinue bool
		wantErr      bool
	}{
		{
			name:         "New style",
			data:         []byte("[[Bold]]text"),
			styles:       []string{},
			wantStyles:   []string{"Bold"},
			wantAdvance:  8,
			wantContinue: false,
			wantErr:      false,
		},
		{
			name:         "End style",
			data:         []byte("[[end]]more"),
			styles:       []string{"Red"},
			wantStyles:   []string{},
			wantAdvance:  7,
			wantContinue: false,
			wantErr:      false,
		},
		{
			name:         "Incomplete tag",
			data:         []byte("[[Red"),
			styles:       []string{},
			wantStyles:   nil,
			wantAdvance:  0,
			wantContinue: true,
			wantErr:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			newStyles, advance, continueParsing, err := splitStyles(tt.data, tt.styles)
			if (err != nil) != tt.wantErr {
				t.Errorf("splitStyles() error = %v, wantErr %v", err, tt.wantErr)
			}
			if advance != tt.wantAdvance {
				t.Errorf("splitStyles() advance = %d, want %d", advance, tt.wantAdvance)
			}
			if continueParsing != tt.wantContinue {
				t.Errorf("splitStyles() continueParsing = %v, want %v", continueParsing, tt.wantContinue)
			}
			if !stylesTextsMatch(newStyles, tt.wantStyles) {
				t.Errorf("splitStyles() styles = %v, want %v", newStyles, tt.wantStyles)
			}
		})
	}
}
