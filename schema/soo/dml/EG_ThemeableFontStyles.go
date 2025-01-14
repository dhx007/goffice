package dml

import (
	"encoding/xml"

	"github.com/dhx007/goffice"
)

type EG_ThemeableFontStyles struct {
	Font    *CT_FontCollection
	FontRef *CT_FontReference
}

func NewEG_ThemeableFontStyles() *EG_ThemeableFontStyles {
	ret := &EG_ThemeableFontStyles{}
	return ret
}

func (m *EG_ThemeableFontStyles) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.Font != nil {
		sefont := xml.StartElement{Name: xml.Name{Local: "a:font"}}
		e.EncodeElement(m.Font, sefont)
	}
	if m.FontRef != nil {
		sefontRef := xml.StartElement{Name: xml.Name{Local: "a:fontRef"}}
		e.EncodeElement(m.FontRef, sefontRef)
	}
	return nil
}

func (m *EG_ThemeableFontStyles) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lEG_ThemeableFontStyles:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "font"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "font"}:
				m.Font = NewCT_FontCollection()
				if err := d.DecodeElement(m.Font, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "fontRef"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "fontRef"}:
				m.FontRef = NewCT_FontReference()
				if err := d.DecodeElement(m.FontRef, &el); err != nil {
					return err
				}
			default:
				goffice.Log("skipping unsupported element on EG_ThemeableFontStyles %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lEG_ThemeableFontStyles
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the EG_ThemeableFontStyles and its children
func (m *EG_ThemeableFontStyles) Validate() error {
	return m.ValidateWithPath("EG_ThemeableFontStyles")
}

// ValidateWithPath validates the EG_ThemeableFontStyles and its children, prefixing error messages with path
func (m *EG_ThemeableFontStyles) ValidateWithPath(path string) error {
	if m.Font != nil {
		if err := m.Font.ValidateWithPath(path + "/Font"); err != nil {
			return err
		}
	}
	if m.FontRef != nil {
		if err := m.FontRef.ValidateWithPath(path + "/FontRef"); err != nil {
			return err
		}
	}
	return nil
}
