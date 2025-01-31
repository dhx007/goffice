package pml

import (
	"encoding/xml"

	"github.com/dhx007/goffice"
	"github.com/dhx007/goffice/schema/soo/dml"
)

type EG_TopLevelSlide struct {
	// Color Scheme Map
	ClrMap *dml.CT_ColorMapping
}

func NewEG_TopLevelSlide() *EG_TopLevelSlide {
	ret := &EG_TopLevelSlide{}
	ret.ClrMap = dml.NewCT_ColorMapping()
	return ret
}

func (m *EG_TopLevelSlide) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	seclrMap := xml.StartElement{Name: xml.Name{Local: "p:clrMap"}}
	e.EncodeElement(m.ClrMap, seclrMap)
	return nil
}

func (m *EG_TopLevelSlide) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.ClrMap = dml.NewCT_ColorMapping()
lEG_TopLevelSlide:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "clrMap"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "clrMap"}:
				if err := d.DecodeElement(m.ClrMap, &el); err != nil {
					return err
				}
			default:
				goffice.Log("skipping unsupported element on EG_TopLevelSlide %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lEG_TopLevelSlide
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the EG_TopLevelSlide and its children
func (m *EG_TopLevelSlide) Validate() error {
	return m.ValidateWithPath("EG_TopLevelSlide")
}

// ValidateWithPath validates the EG_TopLevelSlide and its children, prefixing error messages with path
func (m *EG_TopLevelSlide) ValidateWithPath(path string) error {
	if err := m.ClrMap.ValidateWithPath(path + "/ClrMap"); err != nil {
		return err
	}
	return nil
}
