package statement

import (
	"github.com/caravan/kombi/parse"
	"github.com/caravan/queries/internal/query/parser/ast"
	"github.com/caravan/queries/internal/query/parser/expression"
	"github.com/caravan/queries/internal/query/parser/literal"
)

// SELECT Parser
var (
	condMet    = &struct{}{}
	condNotMet = &struct{}{}

	Alias       = literal.AS.Then(literal.Identifier).Or(literal.Identifier)
	FromClause  = RequireIf(literal.FROM, SourceSelectors)
	WhereClause = RequireIf(literal.WHERE, WhereSelector)

	SelectStatement = parse.Parser(
		func(i parse.Input) (*parse.Success, *parse.Failure) {
			res := &ast.SelectStatement{}
			sel := literal.SELECT.
				Then(ColumnSelectors.Capture(func(cs parse.Result) {
					res.ColumnSelectors = cs.(ast.ColumnSelectors)
				})).
				Then(FromClause.Capture(func(ss parse.Result) {
					if ss != nil {
						res.SourceSelectors = ss.(ast.SourceSelectors)
					}
				}).Optional()).
				Then(WhereClause.Capture(func(wc parse.Result) {
					if wc != nil {
						res.SelectCondition = wc.(*ast.SelectCondition)
					}
				}).Optional()).
				Return(res)
			return sel(i)
		},
	)

	ColumnSelectors = ColumnSelector.Delimited(literal.Comma).Combine(
		func(in ...parse.Result) parse.Result {
			out := make(ast.ColumnSelectors, len(in))
			for i, s := range in {
				out[i] = s.(*ast.ColumnSelector)
			}
			return out
		},
	)

	ColumnSelector = parse.Parser(
		func(i parse.Input) (*parse.Success, *parse.Failure) {
			res := &ast.ColumnSelector{}
			col := expression.Expression.
				Capture(func(r parse.Result) {
					res.Expression = r.(ast.Expression)
				}).
				Then(Alias.Capture(func(r parse.Result) {
					res.Name = r.(*ast.Identifier).Name
				})).
				Or(literal.Identifier.Capture(func(r parse.Result) {
					id := r.(*ast.Identifier)
					res.Expression = id
					res.Name = id.Name
				})).
				Return(res)
			return col(i)
		},
	)

	SourceSelectors = SourceSelector.Delimited(literal.Comma).Combine(
		func(in ...parse.Result) parse.Result {
			out := make(ast.SourceSelectors, len(in))
			for i, s := range in {
				out[i] = s.(*ast.SourceSelector)
			}
			return out
		},
	)

	SourceSelector = parse.Parser(
		func(i parse.Input) (*parse.Success, *parse.Failure) {
			res := &ast.SourceSelector{}
			src := literal.Identifier.
				Capture(func(r parse.Result) {
					id := r.(*ast.Identifier)
					res.Name = id.Name
					res.Source = id.Name
				}).
				Then(Alias.Capture(func(r parse.Result) {
					res.Name = r.(*ast.Identifier).Name
				}).Optional()).
				Return(res)
			return src(i)
		},
	)

	WhereSelector = literal.Identifier.Map(func(r parse.Result) parse.Result {
		return &ast.SelectCondition{
			Expression: r.(ast.Expression),
		}
	})
)
