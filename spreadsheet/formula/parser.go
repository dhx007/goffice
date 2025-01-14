package formula

import (
	"io"
	"strings"

	"github.com/dhx007/goffice"
)

//go:generate goyacc -l -o grammar.go  grammar.y
type plex struct {
	nodes  chan *node
	result Expression
}

func (f *plex) Lex(lval *yySymType) int {
	//yyDebug = 3
	yyErrorVerbose = true
	n := <-f.nodes
	if n != nil {
		lval.node = n
		return int(lval.node.token)
	}
	return 0
}

func (f *plex) Error(s string) {
	goffice.Log("parse error: %s", s)
}

func Parse(r io.Reader) Expression {
	p := &plex{LexReader(r), nil}
	yyParse(p)
	return p.result
}

func ParseString(s string) Expression {
	if s == "" {
		return NewEmptyExpr()
	}
	return Parse(strings.NewReader(s))
}
