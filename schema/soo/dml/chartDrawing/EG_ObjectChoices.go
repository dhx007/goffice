package chartDrawing

import (
	"encoding/xml"

	"github.com/dhx007/goffice"
)

type EG_ObjectChoices struct {
	Choice *EG_ObjectChoicesChoice
}

func NewEG_ObjectChoices() *EG_ObjectChoices {
	ret := &EG_ObjectChoices{}
	return ret
}

func (m *EG_ObjectChoices) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.Choice != nil {
		m.Choice.MarshalXML(e, xml.StartElement{})
	}
	return nil
}

func (m *EG_ObjectChoices) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lEG_ObjectChoices:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
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
				goffice.Log("skipping unsupported element on EG_ObjectChoices %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lEG_ObjectChoices
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the EG_ObjectChoices and its children
func (m *EG_ObjectChoices) Validate() error {
	return m.ValidateWithPath("EG_ObjectChoices")
}

// ValidateWithPath validates the EG_ObjectChoices and its children, prefixing error messages with path
func (m *EG_ObjectChoices) ValidateWithPath(path string) error {
	if m.Choice != nil {
		if err := m.Choice.ValidateWithPath(path + "/Choice"); err != nil {
			return err
		}
	}
	return nil
}
