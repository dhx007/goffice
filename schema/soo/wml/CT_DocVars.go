package wml

import (
	"encoding/xml"
	"fmt"

	"github.com/dhx007/goffice"
)

type CT_DocVars struct {
	// Single Document Variable
	DocVar []*CT_DocVar
}

func NewCT_DocVars() *CT_DocVars {
	ret := &CT_DocVars{}
	return ret
}

func (m *CT_DocVars) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.DocVar != nil {
		sedocVar := xml.StartElement{Name: xml.Name{Local: "w:docVar"}}
		for _, c := range m.DocVar {
			e.EncodeElement(c, sedocVar)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_DocVars) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_DocVars:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "docVar"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "docVar"}:
				tmp := NewCT_DocVar()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.DocVar = append(m.DocVar, tmp)
			default:
				goffice.Log("skipping unsupported element on CT_DocVars %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_DocVars
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_DocVars and its children
func (m *CT_DocVars) Validate() error {
	return m.ValidateWithPath("CT_DocVars")
}

// ValidateWithPath validates the CT_DocVars and its children, prefixing error messages with path
func (m *CT_DocVars) ValidateWithPath(path string) error {
	for i, v := range m.DocVar {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/DocVar[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
