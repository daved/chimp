package chimp

import "io"

type Chimp struct {
	parser *Parse
}

func New(w io.Writer) *Chimp {
	return &Chimp{
		parser: NewParse(w),
	}
}

func (c *Chimp) Write(p []byte) (n int, err error) {
	return c.parser.Write(p)
}
