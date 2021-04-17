package lexer

import "sync"

type (
	// Sequence represents a Lexer sequence
	Sequence interface {
		Split() (*Token, Sequence, bool)
	}

	resolver func() (*Token, Sequence, bool)

	sequence struct {
		sync.Once
		resolver resolver

		ok     bool
		result *Token
		rest   Sequence
	}
)

func newSequence(r resolver) Sequence {
	return &sequence{
		resolver: r,
		result:   nil,
		rest:     nil,
	}
}

func (l *sequence) Split() (*Token, Sequence, bool) {
	l.Do(func() {
		l.result, l.rest, l.ok = l.resolver()
		l.resolver = nil
	})
	return l.result, l.rest, l.ok
}

// Filter returns a filtered Lexer Sequence
func Filter(s Sequence, fn func(*Token) bool) Sequence {
	r := s
	var resolver resolver
	var f *Token
	var ok bool

	resolver = func() (*Token, Sequence, bool) {
		for f, r, ok = r.Split(); ok; f, r, ok = r.Split() {
			if fn(f) {
				return f, newSequence(resolver), true
			}
		}
		return MakeToken(EndOfFile, nil), nil, false
	}

	return newSequence(resolver)
}
