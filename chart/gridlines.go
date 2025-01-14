package chart

import (
	"github.com/dhx007/goffice/drawing"
	"github.com/dhx007/goffice/schema/soo/dml"
	crt "github.com/dhx007/goffice/schema/soo/dml/chart"
)

type GridLines struct {
	x *crt.CT_ChartLines
}

// X returns the inner wrapped XML type.
func (g GridLines) X() *crt.CT_ChartLines {
	return g.x
}

func (g GridLines) Properties() drawing.ShapeProperties {
	if g.x.SpPr == nil {
		g.x.SpPr = dml.NewCT_ShapeProperties()
	}
	return drawing.MakeShapeProperties(g.x.SpPr)
}
