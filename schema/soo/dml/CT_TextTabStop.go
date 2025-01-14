package dml

import (
	"encoding/xml"
	"fmt"
)

type CT_TextTabStop struct {
	PosAttr  *ST_Coordinate32
	AlgnAttr ST_TextTabAlignType
}

func NewCT_TextTabStop() *CT_TextTabStop {
	ret := &CT_TextTabStop{}
	return ret
}

func (m *CT_TextTabStop) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.PosAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "pos"},
			Value: fmt.Sprintf("%v", *m.PosAttr)})
	}
	if m.AlgnAttr != ST_TextTabAlignTypeUnset {
		attr, err := m.AlgnAttr.MarshalXMLAttr(xml.Name{Local: "algn"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TextTabStop) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "pos" {
			parsed, err := ParseUnionST_Coordinate32(attr.Value)
			if err != nil {
				return err
			}
			m.PosAttr = &parsed
			continue
		}
		if attr.Name.Local == "algn" {
			m.AlgnAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_TextTabStop: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_TextTabStop and its children
func (m *CT_TextTabStop) Validate() error {
	return m.ValidateWithPath("CT_TextTabStop")
}

// ValidateWithPath validates the CT_TextTabStop and its children, prefixing error messages with path
func (m *CT_TextTabStop) ValidateWithPath(path string) error {
	if m.PosAttr != nil {
		if err := m.PosAttr.ValidateWithPath(path + "/PosAttr"); err != nil {
			return err
		}
	}
	if err := m.AlgnAttr.ValidateWithPath(path + "/AlgnAttr"); err != nil {
		return err
	}
	return nil
}
