package sml

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type CT_Top10 struct {
	// Top
	TopAttr *bool
	// Filter by Percent
	PercentAttr *bool
	// Top or Bottom Value
	ValAttr float64
	// Filter Value
	FilterValAttr *float64
}

func NewCT_Top10() *CT_Top10 {
	ret := &CT_Top10{}
	return ret
}

func (m *CT_Top10) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.TopAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "top"},
			Value: fmt.Sprintf("%d", b2i(*m.TopAttr))})
	}
	if m.PercentAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "percent"},
			Value: fmt.Sprintf("%d", b2i(*m.PercentAttr))})
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "val"},
		Value: fmt.Sprintf("%v", m.ValAttr)})
	if m.FilterValAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "filterVal"},
			Value: fmt.Sprintf("%v", *m.FilterValAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Top10) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "top" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.TopAttr = &parsed
			continue
		}
		if attr.Name.Local == "percent" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.PercentAttr = &parsed
			continue
		}
		if attr.Name.Local == "val" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return err
			}
			m.ValAttr = parsed
			continue
		}
		if attr.Name.Local == "filterVal" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return err
			}
			m.FilterValAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_Top10: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_Top10 and its children
func (m *CT_Top10) Validate() error {
	return m.ValidateWithPath("CT_Top10")
}

// ValidateWithPath validates the CT_Top10 and its children, prefixing error messages with path
func (m *CT_Top10) ValidateWithPath(path string) error {
	return nil
}
