package wml

import (
	"encoding/xml"
	"fmt"

	"github.com/dhx007/goffice"
)

type CT_SdtDropDownList struct {
	// Drop-down List Last Saved Value
	LastValueAttr *string
	// Drop-Down List Item
	ListItem []*CT_SdtListItem
}

func NewCT_SdtDropDownList() *CT_SdtDropDownList {
	ret := &CT_SdtDropDownList{}
	return ret
}

func (m *CT_SdtDropDownList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.LastValueAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:lastValue"},
			Value: fmt.Sprintf("%v", *m.LastValueAttr)})
	}
	e.EncodeToken(start)
	if m.ListItem != nil {
		selistItem := xml.StartElement{Name: xml.Name{Local: "w:listItem"}}
		for _, c := range m.ListItem {
			e.EncodeElement(c, selistItem)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_SdtDropDownList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "lastValue" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.LastValueAttr = &parsed
			continue
		}
	}
lCT_SdtDropDownList:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "listItem"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "listItem"}:
				tmp := NewCT_SdtListItem()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.ListItem = append(m.ListItem, tmp)
			default:
				goffice.Log("skipping unsupported element on CT_SdtDropDownList %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_SdtDropDownList
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_SdtDropDownList and its children
func (m *CT_SdtDropDownList) Validate() error {
	return m.ValidateWithPath("CT_SdtDropDownList")
}

// ValidateWithPath validates the CT_SdtDropDownList and its children, prefixing error messages with path
func (m *CT_SdtDropDownList) ValidateWithPath(path string) error {
	for i, v := range m.ListItem {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/ListItem[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
