package scanner

import (
	"bufio"
	"context"
	"io"
	"strings"
)

type Scanner struct {
	scanner *bufio.Scanner
}

func New(r io.Reader) *Scanner {
	return &Scanner{
		scanner: bufio.NewScanner(r),
	}
}

type Reading struct {
	Label, Value string
}

func (s *Scanner) Run(ctx context.Context, c chan<- Reading) {
	defer close(c)
	for s.scanner.Scan() {
		select {
		case <-ctx.Done():
			return
		default:
			t := s.scanner.Text()
			l, v, ok := strings.Cut(t, "\t")
			if ok {
				c <- Reading{
					Label: l,
					Value: v,
				}
			}
		}
	}
}
