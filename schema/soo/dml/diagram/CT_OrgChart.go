package diagram

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type CT_OrgChart struct {
	ValAttr *bool
}

func NewCT_OrgChart() *CT_OrgChart {
	ret := &CT_OrgChart{}
	return ret
}

func (m *CT_OrgChart) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.ValAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "val"},
			Value: fmt.Sprintf("%d", b2i(*m.ValAttr))})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_OrgChart) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "val" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ValAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_OrgChart: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_OrgChart and its children
func (m *CT_OrgChart) Validate() error {
	return m.ValidateWithPath("CT_OrgChart")
}

// ValidateWithPath validates the CT_OrgChart and its children, prefixing error messages with path
func (m *CT_OrgChart) ValidateWithPath(path string) error {
	return nil
}
