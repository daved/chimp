package chimp

import "testing"

func TestApplyStyles(t *testing.T) {
	tests := []struct {
		name   string
		styles []string
		want   string
	}{
		{
			name:   "Red and Bold",
			styles: []string{"Red", "Bold"},
			want:   "\033[31m\033[1m",
		},
		{
			name:   "Unknown style",
			styles: []string{"Foo"},
			want:   "",
		},
		{
			name:   "Empty",
			styles: []string{},
			want:   "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ApplyStyles(tt.styles...)
			if got != tt.want {
				t.Errorf("ApplyStyles(%v) = %q, want %q", tt.styles, got, tt.want)
			}
		})
	}
}
