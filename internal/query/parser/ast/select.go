package ast

type (
	// SelectStatement represents a SQL SELECT Statement
	SelectStatement struct {
		ColumnSelectors
		SourceSelectors
		*SelectCondition
	}

	// ColumnSelector represents a SQL column selector
	ColumnSelector struct {
		Expression
		Name
	}

	// SourceSelector represents a SQL source selector (FROM)
	SourceSelector struct {
		Source Name
		Name
	}

	// SelectCondition represents a SQL select condition (WHERE)
	SelectCondition struct {
		Expression // will be a BooleanExpression
	}

	// ColumnSelectors is a set of ColumnSelector
	ColumnSelectors []*ColumnSelector

	// SourceSelectors is a set of SourceSelector
	SourceSelectors []*SourceSelector
)

// Statement marks SelectStatement as a Statement
func (*SelectStatement) Statement() {}
