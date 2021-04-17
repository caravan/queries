package lexer

type (
	// Located is any value that has line and column location info
	Located interface {
		Line() int
		Column() int
	}

	// Location stores line and column location info
	Location struct {
		line   int
		column int
	}
)

// Line returns the line where this Token occurs
func (l *Location) Line() int {
	return l.line
}

// Column returns the column where this Token occurs
func (l *Location) Column() int {
	return l.column
}
