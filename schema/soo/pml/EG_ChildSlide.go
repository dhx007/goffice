package pml

import (
	"encoding/xml"

	"github.com/dhx007/goffice"
	"github.com/dhx007/goffice/schema/soo/dml"
)

type EG_ChildSlide struct {
	// Color Scheme Map Override
	ClrMapOvr *dml.CT_ColorMappingOverride
}

func NewEG_ChildSlide() *EG_ChildSlide {
	ret := &EG_ChildSlide{}
	return ret
}

func (m *EG_ChildSlide) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.ClrMapOvr != nil {
		seclrMapOvr := xml.StartElement{Name: xml.Name{Local: "p:clrMapOvr"}}
		e.EncodeElement(m.ClrMapOvr, seclrMapOvr)
	}
	return nil
}

func (m *EG_ChildSlide) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lEG_ChildSlide:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "clrMapOvr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "clrMapOvr"}:
				m.ClrMapOvr = dml.NewCT_ColorMappingOverride()
				if err := d.DecodeElement(m.ClrMapOvr, &el); err != nil {
					return err
				}
			default:
				goffice.Log("skipping unsupported element on EG_ChildSlide %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lEG_ChildSlide
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the EG_ChildSlide and its children
func (m *EG_ChildSlide) Validate() error {
	return m.ValidateWithPath("EG_ChildSlide")
}

// ValidateWithPath validates the EG_ChildSlide and its children, prefixing error messages with path
func (m *EG_ChildSlide) ValidateWithPath(path string) error {
	if m.ClrMapOvr != nil {
		if err := m.ClrMapOvr.ValidateWithPath(path + "/ClrMapOvr"); err != nil {
			return err
		}
	}
	return nil
}
