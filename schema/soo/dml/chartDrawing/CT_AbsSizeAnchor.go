package chartDrawing

import (
	"encoding/xml"

	"github.com/dhx007/goffice"
	"github.com/dhx007/goffice/schema/soo/dml"
)

type CT_AbsSizeAnchor struct {
	From   *CT_Marker
	Ext    *dml.CT_PositiveSize2D
	Choice *EG_ObjectChoicesChoice
}

func NewCT_AbsSizeAnchor() *CT_AbsSizeAnchor {
	ret := &CT_AbsSizeAnchor{}
	ret.From = NewCT_Marker()
	ret.Ext = dml.NewCT_PositiveSize2D()
	return ret
}

func (m *CT_AbsSizeAnchor) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	sefrom := xml.StartElement{Name: xml.Name{Local: "from"}}
	e.EncodeElement(m.From, sefrom)
	seext := xml.StartElement{Name: xml.Name{Local: "ext"}}
	e.EncodeElement(m.Ext, seext)
	if m.Choice != nil {
		m.Choice.MarshalXML(e, xml.StartElement{})
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_AbsSizeAnchor) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.From = NewCT_Marker()
	m.Ext = dml.NewCT_PositiveSize2D()
lCT_AbsSizeAnchor:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", Local: "from"}:
				if err := d.DecodeElement(m.From, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", Local: "ext"}:
				if err := d.DecodeElement(m.Ext, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", Local: "sp"}:
				m.Choice = NewEG_ObjectChoicesChoice()
				if err := d.DecodeElement(&m.Choice.Sp, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", Local: "grpSp"}:
				m.Choice = NewEG_ObjectChoicesChoice()
				if err := d.DecodeElement(&m.Choice.GrpSp, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", Local: "graphicFrame"}:
				m.Choice = NewEG_ObjectChoicesChoice()
				if err := d.DecodeElement(&m.Choice.GraphicFrame, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", Local: "cxnSp"}:
				m.Choice = NewEG_ObjectChoicesChoice()
				if err := d.DecodeElement(&m.Choice.CxnSp, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", Local: "pic"}:
				m.Choice = NewEG_ObjectChoicesChoice()
				if err := d.DecodeElement(&m.Choice.Pic, &el); err != nil {
					return err
				}
			default:
				goffice.Log("skipping unsupported element on CT_AbsSizeAnchor %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_AbsSizeAnchor
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_AbsSizeAnchor and its children
func (m *CT_AbsSizeAnchor) Validate() error {
	return m.ValidateWithPath("CT_AbsSizeAnchor")
}

// ValidateWithPath validates the CT_AbsSizeAnchor and its children, prefixing error messages with path
func (m *CT_AbsSizeAnchor) ValidateWithPath(path string) error {
	if err := m.From.ValidateWithPath(path + "/From"); err != nil {
		return err
	}
	if err := m.Ext.ValidateWithPath(path + "/Ext"); err != nil {
		return err
	}
	if m.Choice != nil {
		if err := m.Choice.ValidateWithPath(path + "/Choice"); err != nil {
			return err
		}
	}
	return nil
}
