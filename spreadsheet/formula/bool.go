package formula

import (
	"strconv"

	"github.com/dhx007/goffice"
	"github.com/dhx007/goffice/spreadsheet/update"
)

// Bool is a boolean expression.
type Bool struct {
	b bool
}

// NewBool constructs a new boolean expression.
func NewBool(v string) Expression {
	b, err := strconv.ParseBool(v)
	if err != nil {
		goffice.Log("error parsing formula bool %s: %s", v, err)
	}
	return Bool{b}
}

// Eval evaluates and returns a boolean.
func (b Bool) Eval(ctx Context, ev Evaluator) Result {
	return MakeBoolResult(b.b)
}

// Reference returns an invalid reference for Bool.
func (b Bool) Reference(ctx Context, ev Evaluator) Reference {
	return ReferenceInvalid
}

// String returns a string representation for Bool.
func (b Bool) String() string {
	if b.b {
		return "TRUE"
	} else {
		return "FALSE"
	}
}

// Update returns the same object as updating sheet references does not affect Bool.
func (b Bool) Update(q *update.UpdateQuery) Expression {
	return b
}
