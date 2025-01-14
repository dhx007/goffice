package wml

import (
	"encoding/xml"
	"fmt"
)

type CT_MailMergeDataType struct {
	// Value
	ValAttr string
}

func NewCT_MailMergeDataType() *CT_MailMergeDataType {
	ret := &CT_MailMergeDataType{}
	return ret
}

func (m *CT_MailMergeDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"},
		Value: fmt.Sprintf("%v", m.ValAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_MailMergeDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "val" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ValAttr = parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_MailMergeDataType: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_MailMergeDataType and its children
func (m *CT_MailMergeDataType) Validate() error {
	return m.ValidateWithPath("CT_MailMergeDataType")
}

// ValidateWithPath validates the CT_MailMergeDataType and its children, prefixing error messages with path
func (m *CT_MailMergeDataType) ValidateWithPath(path string) error {
	return nil
}
