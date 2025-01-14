package vml

import (
	"encoding/xml"
	"fmt"

	"github.com/dhx007/goffice"
)

type OfcCT_RegroupTable struct {
	Entry   []*OfcCT_Entry
	ExtAttr ST_Ext
}

func NewOfcCT_RegroupTable() *OfcCT_RegroupTable {
	ret := &OfcCT_RegroupTable{}
	return ret
}

func (m *OfcCT_RegroupTable) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.ExtAttr != ST_ExtUnset {
		attr, err := m.ExtAttr.MarshalXMLAttr(xml.Name{Local: "ext"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	if m.Entry != nil {
		seentry := xml.StartElement{Name: xml.Name{Local: "o:entry"}}
		for _, c := range m.Entry {
			e.EncodeElement(c, seentry)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *OfcCT_RegroupTable) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "ext" {
			m.ExtAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
lOfcCT_RegroupTable:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "urn:schemas-microsoft-com:office:office", Local: "entry"}:
				tmp := NewOfcCT_Entry()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Entry = append(m.Entry, tmp)
			default:
				goffice.Log("skipping unsupported element on OfcCT_RegroupTable %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lOfcCT_RegroupTable
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the OfcCT_RegroupTable and its children
func (m *OfcCT_RegroupTable) Validate() error {
	return m.ValidateWithPath("OfcCT_RegroupTable")
}

// ValidateWithPath validates the OfcCT_RegroupTable and its children, prefixing error messages with path
func (m *OfcCT_RegroupTable) ValidateWithPath(path string) error {
	for i, v := range m.Entry {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Entry[%d]", path, i)); err != nil {
			return err
		}
	}
	if err := m.ExtAttr.ValidateWithPath(path + "/ExtAttr"); err != nil {
		return err
	}
	return nil
}
