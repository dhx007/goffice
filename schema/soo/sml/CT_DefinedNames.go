package sml

import (
	"encoding/xml"
	"fmt"

	"github.com/dhx007/goffice"
)

type CT_DefinedNames struct {
	// Defined Name
	DefinedName []*CT_DefinedName
}

func NewCT_DefinedNames() *CT_DefinedNames {
	ret := &CT_DefinedNames{}
	return ret
}

func (m *CT_DefinedNames) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.DefinedName != nil {
		sedefinedName := xml.StartElement{Name: xml.Name{Local: "ma:definedName"}}
		for _, c := range m.DefinedName {
			e.EncodeElement(c, sedefinedName)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_DefinedNames) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_DefinedNames:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "definedName"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "definedName"}:
				tmp := NewCT_DefinedName()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.DefinedName = append(m.DefinedName, tmp)
			default:
				goffice.Log("skipping unsupported element on CT_DefinedNames %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_DefinedNames
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_DefinedNames and its children
func (m *CT_DefinedNames) Validate() error {
	return m.ValidateWithPath("CT_DefinedNames")
}

// ValidateWithPath validates the CT_DefinedNames and its children, prefixing error messages with path
func (m *CT_DefinedNames) ValidateWithPath(path string) error {
	for i, v := range m.DefinedName {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/DefinedName[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
