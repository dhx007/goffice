package sml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/dhx007/goffice"
)

type CT_WebPublishItems struct {
	// Web Publishing Items Count
	CountAttr *uint32
	// Web Publishing Item
	WebPublishItem []*CT_WebPublishItem
}

func NewCT_WebPublishItems() *CT_WebPublishItems {
	ret := &CT_WebPublishItems{}
	return ret
}

func (m *CT_WebPublishItems) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.CountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "count"},
			Value: fmt.Sprintf("%v", *m.CountAttr)})
	}
	e.EncodeToken(start)
	sewebPublishItem := xml.StartElement{Name: xml.Name{Local: "ma:webPublishItem"}}
	for _, c := range m.WebPublishItem {
		e.EncodeElement(c, sewebPublishItem)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_WebPublishItems) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "count" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.CountAttr = &pt
			continue
		}
	}
lCT_WebPublishItems:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "webPublishItem"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "webPublishItem"}:
				tmp := NewCT_WebPublishItem()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.WebPublishItem = append(m.WebPublishItem, tmp)
			default:
				goffice.Log("skipping unsupported element on CT_WebPublishItems %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_WebPublishItems
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_WebPublishItems and its children
func (m *CT_WebPublishItems) Validate() error {
	return m.ValidateWithPath("CT_WebPublishItems")
}

// ValidateWithPath validates the CT_WebPublishItems and its children, prefixing error messages with path
func (m *CT_WebPublishItems) ValidateWithPath(path string) error {
	for i, v := range m.WebPublishItem {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/WebPublishItem[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
