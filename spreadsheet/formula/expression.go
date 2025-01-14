package formula

import "github.com/dhx007/goffice/spreadsheet/update"

type Expression interface {
	Eval(ctx Context, ev Evaluator) Result
	Reference(ctx Context, ev Evaluator) Reference
	String() string
	Update(updateQuery *update.UpdateQuery) Expression
}
