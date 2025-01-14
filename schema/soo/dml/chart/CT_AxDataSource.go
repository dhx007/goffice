package chart

import (
	"encoding/xml"

	"github.com/dhx007/goffice"
)

type CT_AxDataSource struct {
	Choice *CT_AxDataSourceChoice
}

func NewCT_AxDataSource() *CT_AxDataSource {
	ret := &CT_AxDataSource{}
	ret.Choice = NewCT_AxDataSourceChoice()
	return ret
}

func (m *CT_AxDataSource) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	m.Choice.MarshalXML(e, xml.StartElement{})
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_AxDataSource) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Choice = NewCT_AxDataSourceChoice()
lCT_AxDataSource:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "multiLvlStrRef"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "multiLvlStrRef"}:
				m.Choice = NewCT_AxDataSourceChoice()
				if err := d.DecodeElement(&m.Choice.MultiLvlStrRef, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "numRef"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "numRef"}:
				m.Choice = NewCT_AxDataSourceChoice()
				if err := d.DecodeElement(&m.Choice.NumRef, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "numLit"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "numLit"}:
				m.Choice = NewCT_AxDataSourceChoice()
				if err := d.DecodeElement(&m.Choice.NumLit, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "strRef"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "strRef"}:
				m.Choice = NewCT_AxDataSourceChoice()
				if err := d.DecodeElement(&m.Choice.StrRef, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "strLit"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "strLit"}:
				m.Choice = NewCT_AxDataSourceChoice()
				if err := d.DecodeElement(&m.Choice.StrLit, &el); err != nil {
					return err
				}
			default:
				goffice.Log("skipping unsupported element on CT_AxDataSource %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_AxDataSource
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_AxDataSource and its children
func (m *CT_AxDataSource) Validate() error {
	return m.ValidateWithPath("CT_AxDataSource")
}

// ValidateWithPath validates the CT_AxDataSource and its children, prefixing error messages with path
func (m *CT_AxDataSource) ValidateWithPath(path string) error {
	if err := m.Choice.ValidateWithPath(path + "/Choice"); err != nil {
		return err
	}
	return nil
}
