package diagram

import (
	"encoding/xml"
	"fmt"

	"github.com/dhx007/goffice"
)

type CT_Constraints struct {
	Constr []*CT_Constraint
}

func NewCT_Constraints() *CT_Constraints {
	ret := &CT_Constraints{}
	return ret
}

func (m *CT_Constraints) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.Constr != nil {
		seconstr := xml.StartElement{Name: xml.Name{Local: "constr"}}
		for _, c := range m.Constr {
			e.EncodeElement(c, seconstr)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Constraints) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_Constraints:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "constr"}:
				tmp := NewCT_Constraint()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Constr = append(m.Constr, tmp)
			default:
				goffice.Log("skipping unsupported element on CT_Constraints %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Constraints
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Constraints and its children
func (m *CT_Constraints) Validate() error {
	return m.ValidateWithPath("CT_Constraints")
}

// ValidateWithPath validates the CT_Constraints and its children, prefixing error messages with path
func (m *CT_Constraints) ValidateWithPath(path string) error {
	for i, v := range m.Constr {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Constr[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
