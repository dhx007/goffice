package vml

import (
	"encoding/xml"
	"fmt"

	"github.com/dhx007/goffice"
)

type OfcCT_RelationTable struct {
	Rel     []*OfcCT_Relation
	ExtAttr ST_Ext
}

func NewOfcCT_RelationTable() *OfcCT_RelationTable {
	ret := &OfcCT_RelationTable{}
	return ret
}

func (m *OfcCT_RelationTable) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.ExtAttr != ST_ExtUnset {
		attr, err := m.ExtAttr.MarshalXMLAttr(xml.Name{Local: "ext"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	if m.Rel != nil {
		serel := xml.StartElement{Name: xml.Name{Local: "o:rel"}}
		for _, c := range m.Rel {
			e.EncodeElement(c, serel)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *OfcCT_RelationTable) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "ext" {
			m.ExtAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
lOfcCT_RelationTable:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "urn:schemas-microsoft-com:office:office", Local: "rel"}:
				tmp := NewOfcCT_Relation()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Rel = append(m.Rel, tmp)
			default:
				goffice.Log("skipping unsupported element on OfcCT_RelationTable %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lOfcCT_RelationTable
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the OfcCT_RelationTable and its children
func (m *OfcCT_RelationTable) Validate() error {
	return m.ValidateWithPath("OfcCT_RelationTable")
}

// ValidateWithPath validates the OfcCT_RelationTable and its children, prefixing error messages with path
func (m *OfcCT_RelationTable) ValidateWithPath(path string) error {
	for i, v := range m.Rel {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Rel[%d]", path, i)); err != nil {
			return err
		}
	}
	if err := m.ExtAttr.ValidateWithPath(path + "/ExtAttr"); err != nil {
		return err
	}
	return nil
}
