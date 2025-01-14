package presentation

import (
	"github.com/dhx007/goffice/drawing"
	"github.com/dhx007/goffice/schema/soo/dml"
	"github.com/dhx007/goffice/schema/soo/pml"
)

// TextBox is a text box within a slide.
type TextBox struct {
	x *pml.CT_Shape
}

// AddParagraph adds a paragraph to the text box
func (t TextBox) AddParagraph() drawing.Paragraph {
	p := dml.NewCT_TextParagraph()
	t.x.TxBody.P = append(t.x.TxBody.P, p)
	return drawing.MakeParagraph(p)
}

// Properties returns the properties of the TextBox.
func (t TextBox) Properties() drawing.ShapeProperties {
	if t.x.SpPr == nil {
		t.x.SpPr = dml.NewCT_ShapeProperties()
	}
	return drawing.MakeShapeProperties(t.x.SpPr)
}

// SetTextAnchor controls the text anchoring
func (t TextBox) SetTextAnchor(a dml.ST_TextAnchoringType) {
	t.x.TxBody.BodyPr = dml.NewCT_TextBodyProperties()
	t.x.TxBody.BodyPr.AnchorAttr = a
}
