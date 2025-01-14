package formula

import (
	"strconv"

	"github.com/dhx007/goffice"
	"github.com/dhx007/goffice/spreadsheet/update"
)

// Number is a nubmer expression.
type Number struct {
	v float64
}

// NewNumber constructs a new number expression.
func NewNumber(v string) Expression {
	f, err := strconv.ParseFloat(v, 64)
	if err != nil {
		goffice.Log("error parsing formula number %s: %s", v, err)
	}
	return Number{f}
}

// Eval evaluates and returns a number.
func (n Number) Eval(ctx Context, ev Evaluator) Result {
	return MakeNumberResult(n.v)
}

// Reference returns an invalid reference for Number.
func (n Number) Reference(ctx Context, ev Evaluator) Reference {
	return ReferenceInvalid
}

// String returns a string representation of Number.
func (n Number) String() string {
	return strconv.FormatFloat(n.v, 'f', -1, 64)
}

// Update returns the same object as updating sheet references does not affect Number.
func (n Number) Update(q *update.UpdateQuery) Expression {
	return n
}
