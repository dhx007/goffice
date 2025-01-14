package chart

import (
	"encoding/xml"
	"fmt"

	"github.com/dhx007/goffice"
)

type CT_PivotSource struct {
	Name   string
	FmtId  *CT_UnsignedInt
	ExtLst []*CT_ExtensionList
}

func NewCT_PivotSource() *CT_PivotSource {
	ret := &CT_PivotSource{}
	ret.FmtId = NewCT_UnsignedInt()
	return ret
}

func (m *CT_PivotSource) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	sename := xml.StartElement{Name: xml.Name{Local: "c:name"}}
	goffice.AddPreserveSpaceAttr(&sename, m.Name)
	e.EncodeElement(m.Name, sename)
	sefmtId := xml.StartElement{Name: xml.Name{Local: "c:fmtId"}}
	e.EncodeElement(m.FmtId, sefmtId)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "c:extLst"}}
		for _, c := range m.ExtLst {
			e.EncodeElement(c, seextLst)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_PivotSource) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.FmtId = NewCT_UnsignedInt()
lCT_PivotSource:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "name"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "name"}:
				if err := d.DecodeElement(&m.Name, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "fmtId"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "fmtId"}:
				if err := d.DecodeElement(m.FmtId, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "extLst"}:
				tmp := NewCT_ExtensionList()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.ExtLst = append(m.ExtLst, tmp)
			default:
				goffice.Log("skipping unsupported element on CT_PivotSource %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_PivotSource
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_PivotSource and its children
func (m *CT_PivotSource) Validate() error {
	return m.ValidateWithPath("CT_PivotSource")
}

// ValidateWithPath validates the CT_PivotSource and its children, prefixing error messages with path
func (m *CT_PivotSource) ValidateWithPath(path string) error {
	if err := m.FmtId.ValidateWithPath(path + "/FmtId"); err != nil {
		return err
	}
	for i, v := range m.ExtLst {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/ExtLst[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
