package dml

import (
	"encoding/xml"

	"github.com/dhx007/goffice"
)

type EG_TextUnderlineFill struct {
	UFillTx *CT_TextUnderlineFillFollowText
	UFill   *CT_TextUnderlineFillGroupWrapper
}

func NewEG_TextUnderlineFill() *EG_TextUnderlineFill {
	ret := &EG_TextUnderlineFill{}
	return ret
}

func (m *EG_TextUnderlineFill) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.UFillTx != nil {
		seuFillTx := xml.StartElement{Name: xml.Name{Local: "a:uFillTx"}}
		e.EncodeElement(m.UFillTx, seuFillTx)
	}
	if m.UFill != nil {
		seuFill := xml.StartElement{Name: xml.Name{Local: "a:uFill"}}
		e.EncodeElement(m.UFill, seuFill)
	}
	return nil
}

func (m *EG_TextUnderlineFill) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lEG_TextUnderlineFill:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "uFillTx"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "uFillTx"}:
				m.UFillTx = NewCT_TextUnderlineFillFollowText()
				if err := d.DecodeElement(m.UFillTx, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "uFill"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "uFill"}:
				m.UFill = NewCT_TextUnderlineFillGroupWrapper()
				if err := d.DecodeElement(m.UFill, &el); err != nil {
					return err
				}
			default:
				goffice.Log("skipping unsupported element on EG_TextUnderlineFill %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lEG_TextUnderlineFill
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the EG_TextUnderlineFill and its children
func (m *EG_TextUnderlineFill) Validate() error {
	return m.ValidateWithPath("EG_TextUnderlineFill")
}

// ValidateWithPath validates the EG_TextUnderlineFill and its children, prefixing error messages with path
func (m *EG_TextUnderlineFill) ValidateWithPath(path string) error {
	if m.UFillTx != nil {
		if err := m.UFillTx.ValidateWithPath(path + "/UFillTx"); err != nil {
			return err
		}
	}
	if m.UFill != nil {
		if err := m.UFill.ValidateWithPath(path + "/UFill"); err != nil {
			return err
		}
	}
	return nil
}
