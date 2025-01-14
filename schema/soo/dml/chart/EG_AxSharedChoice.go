package chart

import (
	"encoding/xml"

	"github.com/dhx007/goffice"
)

type EG_AxSharedChoice struct {
	Crosses   *CT_Crosses
	CrossesAt *CT_Double
}

func NewEG_AxSharedChoice() *EG_AxSharedChoice {
	ret := &EG_AxSharedChoice{}
	return ret
}

func (m *EG_AxSharedChoice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.Crosses != nil {
		secrosses := xml.StartElement{Name: xml.Name{Local: "c:crosses"}}
		e.EncodeElement(m.Crosses, secrosses)
	}
	if m.CrossesAt != nil {
		secrossesAt := xml.StartElement{Name: xml.Name{Local: "c:crossesAt"}}
		e.EncodeElement(m.CrossesAt, secrossesAt)
	}
	return nil
}

func (m *EG_AxSharedChoice) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lEG_AxSharedChoice:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "crosses"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "crosses"}:
				m.Crosses = NewCT_Crosses()
				if err := d.DecodeElement(m.Crosses, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "crossesAt"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "crossesAt"}:
				m.CrossesAt = NewCT_Double()
				if err := d.DecodeElement(m.CrossesAt, &el); err != nil {
					return err
				}
			default:
				goffice.Log("skipping unsupported element on EG_AxSharedChoice %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lEG_AxSharedChoice
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the EG_AxSharedChoice and its children
func (m *EG_AxSharedChoice) Validate() error {
	return m.ValidateWithPath("EG_AxSharedChoice")
}

// ValidateWithPath validates the EG_AxSharedChoice and its children, prefixing error messages with path
func (m *EG_AxSharedChoice) ValidateWithPath(path string) error {
	if m.Crosses != nil {
		if err := m.Crosses.ValidateWithPath(path + "/Crosses"); err != nil {
			return err
		}
	}
	if m.CrossesAt != nil {
		if err := m.CrossesAt.ValidateWithPath(path + "/CrossesAt"); err != nil {
			return err
		}
	}
	return nil
}
