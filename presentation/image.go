package presentation

import (
	"github.com/dhx007/goffice/drawing"
	"github.com/dhx007/goffice/schema/soo/dml"
	"github.com/dhx007/goffice/schema/soo/pml"
)

// Image is an image within a slide.
type Image struct {
	x *pml.CT_Picture
}

// Properties returns the properties of the TextBox.
func (i Image) Properties() drawing.ShapeProperties {
	if i.x.SpPr == nil {
		i.x.SpPr = dml.NewCT_ShapeProperties()
	}
	return drawing.MakeShapeProperties(i.x.SpPr)
}
