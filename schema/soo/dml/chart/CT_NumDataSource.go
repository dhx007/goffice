package chart

import (
	"encoding/xml"

	"github.com/dhx007/goffice"
)

type CT_NumDataSource struct {
	Choice *CT_NumDataSourceChoice
}

func NewCT_NumDataSource() *CT_NumDataSource {
	ret := &CT_NumDataSource{}
	ret.Choice = NewCT_NumDataSourceChoice()
	return ret
}

func (m *CT_NumDataSource) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	m.Choice.MarshalXML(e, xml.StartElement{})
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_NumDataSource) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Choice = NewCT_NumDataSourceChoice()
lCT_NumDataSource:
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
				m.Choice = NewCT_NumDataSourceChoice()
				if err := d.DecodeElement(&m.Choice.NumRef, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "numLit"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "numLit"}:
				m.Choice = NewCT_NumDataSourceChoice()
				if err := d.DecodeElement(&m.Choice.NumLit, &el); err != nil {
					return err
				}
			default:
				goffice.Log("skipping unsupported element on CT_NumDataSource %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_NumDataSource
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_NumDataSource and its children
func (m *CT_NumDataSource) Validate() error {
	return m.ValidateWithPath("CT_NumDataSource")
}

// ValidateWithPath validates the CT_NumDataSource and its children, prefixing error messages with path
func (m *CT_NumDataSource) ValidateWithPath(path string) error {
	if err := m.Choice.ValidateWithPath(path + "/Choice"); err != nil {
		return err
	}
	return nil
}
