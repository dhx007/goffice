package dml

import (
	"encoding/xml"
	"fmt"

	"github.com/dhx007/goffice"
)

type CT_ConnectionSiteList struct {
	Cxn []*CT_ConnectionSite
}

func NewCT_ConnectionSiteList() *CT_ConnectionSiteList {
	ret := &CT_ConnectionSiteList{}
	return ret
}

func (m *CT_ConnectionSiteList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.Cxn != nil {
		secxn := xml.StartElement{Name: xml.Name{Local: "a:cxn"}}
		for _, c := range m.Cxn {
			e.EncodeElement(c, secxn)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ConnectionSiteList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_ConnectionSiteList:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "cxn"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "cxn"}:
				tmp := NewCT_ConnectionSite()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Cxn = append(m.Cxn, tmp)
			default:
				goffice.Log("skipping unsupported element on CT_ConnectionSiteList %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ConnectionSiteList
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_ConnectionSiteList and its children
func (m *CT_ConnectionSiteList) Validate() error {
	return m.ValidateWithPath("CT_ConnectionSiteList")
}

// ValidateWithPath validates the CT_ConnectionSiteList and its children, prefixing error messages with path
func (m *CT_ConnectionSiteList) ValidateWithPath(path string) error {
	for i, v := range m.Cxn {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Cxn[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
