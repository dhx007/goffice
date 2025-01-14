package chart

import (
	"encoding/xml"

	"github.com/dhx007/goffice"
)

type CT_NumDataSourceChoice struct {
	NumRef *CT_NumRef
	NumLit *CT_NumData
}

func NewCT_NumDataSourceChoice() *CT_NumDataSourceChoice {
	ret := &CT_NumDataSourceChoice{}
	return ret
}

func (m *CT_NumDataSourceChoice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.NumRef != nil {
		senumRef := xml.StartElement{Name: xml.Name{Local: "c:numRef"}}
		e.EncodeElement(m.NumRef, senumRef)
	}
	if m.NumLit != nil {
		senumLit := xml.StartElement{Name: xml.Name{Local: "c:numLit"}}
		e.EncodeElement(m.NumLit, senumLit)
	}
	return nil
}

func (m *CT_NumDataSourceChoice) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_NumDataSourceChoice:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "numRef"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "numRef"}:
				m.NumRef = NewCT_NumRef()
				if err := d.DecodeElement(m.NumRef, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "numLit"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "numLit"}:
				m.NumLit = NewCT_NumData()
				if err := d.DecodeElement(m.NumLit, &el); err != nil {
					return err
				}
			default:
				goffice.Log("skipping unsupported element on CT_NumDataSourceChoice %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_NumDataSourceChoice
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_NumDataSourceChoice and its children
func (m *CT_NumDataSourceChoice) Validate() error {
	return m.ValidateWithPath("CT_NumDataSourceChoice")
}

// ValidateWithPath validates the CT_NumDataSourceChoice and its children, prefixing error messages with path
func (m *CT_NumDataSourceChoice) ValidateWithPath(path string) error {
	if m.NumRef != nil {
		if err := m.NumRef.ValidateWithPath(path + "/NumRef"); err != nil {
			return err
		}
	}
	if m.NumLit != nil {
		if err := m.NumLit.ValidateWithPath(path + "/NumLit"); err != nil {
			return err
		}
	}
	return nil
}
