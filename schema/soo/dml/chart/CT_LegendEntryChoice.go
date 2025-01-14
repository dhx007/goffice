package chart

import (
	"encoding/xml"

	"github.com/dhx007/goffice"
	"github.com/dhx007/goffice/schema/soo/dml"
)

type CT_LegendEntryChoice struct {
	Delete *CT_Boolean
	TxPr   *dml.CT_TextBody
}

func NewCT_LegendEntryChoice() *CT_LegendEntryChoice {
	ret := &CT_LegendEntryChoice{}
	return ret
}

func (m *CT_LegendEntryChoice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.Delete != nil {
		sedelete := xml.StartElement{Name: xml.Name{Local: "c:delete"}}
		e.EncodeElement(m.Delete, sedelete)
	}
	if m.TxPr != nil {
		setxPr := xml.StartElement{Name: xml.Name{Local: "c:txPr"}}
		e.EncodeElement(m.TxPr, setxPr)
	}
	return nil
}

func (m *CT_LegendEntryChoice) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_LegendEntryChoice:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "delete"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "delete"}:
				m.Delete = NewCT_Boolean()
				if err := d.DecodeElement(m.Delete, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "txPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "txPr"}:
				m.TxPr = dml.NewCT_TextBody()
				if err := d.DecodeElement(m.TxPr, &el); err != nil {
					return err
				}
			default:
				goffice.Log("skipping unsupported element on CT_LegendEntryChoice %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_LegendEntryChoice
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_LegendEntryChoice and its children
func (m *CT_LegendEntryChoice) Validate() error {
	return m.ValidateWithPath("CT_LegendEntryChoice")
}

// ValidateWithPath validates the CT_LegendEntryChoice and its children, prefixing error messages with path
func (m *CT_LegendEntryChoice) ValidateWithPath(path string) error {
	if m.Delete != nil {
		if err := m.Delete.ValidateWithPath(path + "/Delete"); err != nil {
			return err
		}
	}
	if m.TxPr != nil {
		if err := m.TxPr.ValidateWithPath(path + "/TxPr"); err != nil {
			return err
		}
	}
	return nil
}
