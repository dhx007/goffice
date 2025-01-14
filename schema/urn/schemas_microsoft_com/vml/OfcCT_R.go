package vml

import (
	"encoding/xml"
	"fmt"

	"github.com/dhx007/goffice"
)

type OfcCT_R struct {
	IdAttr    string
	TypeAttr  OfcST_RType
	HowAttr   OfcST_How
	IdrefAttr *string
	Proxy     []*OfcCT_Proxy
}

func NewOfcCT_R() *OfcCT_R {
	ret := &OfcCT_R{}
	return ret
}

func (m *OfcCT_R) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "id"},
		Value: fmt.Sprintf("%v", m.IdAttr)})
	if m.TypeAttr != OfcST_RTypeUnset {
		attr, err := m.TypeAttr.MarshalXMLAttr(xml.Name{Local: "type"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.HowAttr != OfcST_HowUnset {
		attr, err := m.HowAttr.MarshalXMLAttr(xml.Name{Local: "how"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.IdrefAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "idref"},
			Value: fmt.Sprintf("%v", *m.IdrefAttr)})
	}
	e.EncodeToken(start)
	if m.Proxy != nil {
		seproxy := xml.StartElement{Name: xml.Name{Local: "o:proxy"}}
		for _, c := range m.Proxy {
			e.EncodeElement(c, seproxy)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *OfcCT_R) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "id" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IdAttr = parsed
			continue
		}
		if attr.Name.Local == "type" {
			m.TypeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "how" {
			m.HowAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "idref" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IdrefAttr = &parsed
			continue
		}
	}
lOfcCT_R:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "urn:schemas-microsoft-com:office:office", Local: "proxy"}:
				tmp := NewOfcCT_Proxy()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Proxy = append(m.Proxy, tmp)
			default:
				goffice.Log("skipping unsupported element on OfcCT_R %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lOfcCT_R
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the OfcCT_R and its children
func (m *OfcCT_R) Validate() error {
	return m.ValidateWithPath("OfcCT_R")
}

// ValidateWithPath validates the OfcCT_R and its children, prefixing error messages with path
func (m *OfcCT_R) ValidateWithPath(path string) error {
	if err := m.TypeAttr.ValidateWithPath(path + "/TypeAttr"); err != nil {
		return err
	}
	if err := m.HowAttr.ValidateWithPath(path + "/HowAttr"); err != nil {
		return err
	}
	for i, v := range m.Proxy {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Proxy[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
