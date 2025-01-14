package chart

import (
	"encoding/xml"
	"fmt"

	"github.com/dhx007/goffice"
)

type CT_StrData struct {
	PtCount *CT_UnsignedInt
	Pt      []*CT_StrVal
	ExtLst  *CT_ExtensionList
}

func NewCT_StrData() *CT_StrData {
	ret := &CT_StrData{}
	return ret
}

func (m *CT_StrData) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.PtCount != nil {
		septCount := xml.StartElement{Name: xml.Name{Local: "c:ptCount"}}
		e.EncodeElement(m.PtCount, septCount)
	}
	if m.Pt != nil {
		sept := xml.StartElement{Name: xml.Name{Local: "c:pt"}}
		for _, c := range m.Pt {
			e.EncodeElement(c, sept)
		}
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "c:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_StrData) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_StrData:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "ptCount"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "ptCount"}:
				m.PtCount = NewCT_UnsignedInt()
				if err := d.DecodeElement(m.PtCount, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "pt"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "pt"}:
				tmp := NewCT_StrVal()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Pt = append(m.Pt, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				goffice.Log("skipping unsupported element on CT_StrData %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_StrData
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_StrData and its children
func (m *CT_StrData) Validate() error {
	return m.ValidateWithPath("CT_StrData")
}

// ValidateWithPath validates the CT_StrData and its children, prefixing error messages with path
func (m *CT_StrData) ValidateWithPath(path string) error {
	if m.PtCount != nil {
		if err := m.PtCount.ValidateWithPath(path + "/PtCount"); err != nil {
			return err
		}
	}
	for i, v := range m.Pt {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Pt[%d]", path, i)); err != nil {
			return err
		}
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
