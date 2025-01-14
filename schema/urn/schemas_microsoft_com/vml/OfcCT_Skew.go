package vml

import (
	"encoding/xml"
	"fmt"

	"github.com/dhx007/goffice/schema/soo/ofc/sharedTypes"
)

type OfcCT_Skew struct {
	IdAttr     *string
	OnAttr     sharedTypes.ST_TrueFalse
	OffsetAttr *string
	OriginAttr *string
	MatrixAttr *string
	ExtAttr    ST_Ext
}

func NewOfcCT_Skew() *OfcCT_Skew {
	ret := &OfcCT_Skew{}
	return ret
}

func (m *OfcCT_Skew) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.IdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "id"},
			Value: fmt.Sprintf("%v", *m.IdAttr)})
	}
	if m.OnAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.OnAttr.MarshalXMLAttr(xml.Name{Local: "on"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.OffsetAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "offset"},
			Value: fmt.Sprintf("%v", *m.OffsetAttr)})
	}
	if m.OriginAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "origin"},
			Value: fmt.Sprintf("%v", *m.OriginAttr)})
	}
	if m.MatrixAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "matrix"},
			Value: fmt.Sprintf("%v", *m.MatrixAttr)})
	}
	if m.ExtAttr != ST_ExtUnset {
		attr, err := m.ExtAttr.MarshalXMLAttr(xml.Name{Local: "ext"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *OfcCT_Skew) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "id" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IdAttr = &parsed
			continue
		}
		if attr.Name.Local == "on" {
			m.OnAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "offset" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.OffsetAttr = &parsed
			continue
		}
		if attr.Name.Local == "origin" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.OriginAttr = &parsed
			continue
		}
		if attr.Name.Local == "matrix" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.MatrixAttr = &parsed
			continue
		}
		if attr.Name.Local == "ext" {
			m.ExtAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing OfcCT_Skew: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the OfcCT_Skew and its children
func (m *OfcCT_Skew) Validate() error {
	return m.ValidateWithPath("OfcCT_Skew")
}

// ValidateWithPath validates the OfcCT_Skew and its children, prefixing error messages with path
func (m *OfcCT_Skew) ValidateWithPath(path string) error {
	if err := m.OnAttr.ValidateWithPath(path + "/OnAttr"); err != nil {
		return err
	}
	if err := m.ExtAttr.ValidateWithPath(path + "/ExtAttr"); err != nil {
		return err
	}
	return nil
}
