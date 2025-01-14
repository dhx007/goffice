package pml

import (
	"encoding/xml"

	"github.com/dhx007/goffice"
	"github.com/dhx007/goffice/schema/soo/dml"
)

type EG_Background struct {
	// Background Properties
	BgPr *CT_BackgroundProperties
	// Background Style Reference
	BgRef *dml.CT_StyleMatrixReference
}

func NewEG_Background() *EG_Background {
	ret := &EG_Background{}
	return ret
}

func (m *EG_Background) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.BgPr != nil {
		sebgPr := xml.StartElement{Name: xml.Name{Local: "p:bgPr"}}
		e.EncodeElement(m.BgPr, sebgPr)
	}
	if m.BgRef != nil {
		sebgRef := xml.StartElement{Name: xml.Name{Local: "p:bgRef"}}
		e.EncodeElement(m.BgRef, sebgRef)
	}
	return nil
}

func (m *EG_Background) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lEG_Background:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "bgPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "bgPr"}:
				m.BgPr = NewCT_BackgroundProperties()
				if err := d.DecodeElement(m.BgPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "bgRef"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "bgRef"}:
				m.BgRef = dml.NewCT_StyleMatrixReference()
				if err := d.DecodeElement(m.BgRef, &el); err != nil {
					return err
				}
			default:
				goffice.Log("skipping unsupported element on EG_Background %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lEG_Background
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the EG_Background and its children
func (m *EG_Background) Validate() error {
	return m.ValidateWithPath("EG_Background")
}

// ValidateWithPath validates the EG_Background and its children, prefixing error messages with path
func (m *EG_Background) ValidateWithPath(path string) error {
	if m.BgPr != nil {
		if err := m.BgPr.ValidateWithPath(path + "/BgPr"); err != nil {
			return err
		}
	}
	if m.BgRef != nil {
		if err := m.BgRef.ValidateWithPath(path + "/BgRef"); err != nil {
			return err
		}
	}
	return nil
}
