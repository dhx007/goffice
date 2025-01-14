package dml

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type CT_PositiveFixedAngle struct {
	ValAttr int32
}

func NewCT_PositiveFixedAngle() *CT_PositiveFixedAngle {
	ret := &CT_PositiveFixedAngle{}
	ret.ValAttr = 0
	return ret
}

func (m *CT_PositiveFixedAngle) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "val"},
		Value: fmt.Sprintf("%v", m.ValAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_PositiveFixedAngle) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.ValAttr = 0
	for _, attr := range start.Attr {
		if attr.Name.Local == "val" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.ValAttr = int32(parsed)
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_PositiveFixedAngle: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_PositiveFixedAngle and its children
func (m *CT_PositiveFixedAngle) Validate() error {
	return m.ValidateWithPath("CT_PositiveFixedAngle")
}

// ValidateWithPath validates the CT_PositiveFixedAngle and its children, prefixing error messages with path
func (m *CT_PositiveFixedAngle) ValidateWithPath(path string) error {
	if m.ValAttr < 0 {
		return fmt.Errorf("%s/m.ValAttr must be >= 0 (have %v)", path, m.ValAttr)
	}
	if m.ValAttr >= 21600000 {
		return fmt.Errorf("%s/m.ValAttr must be < 21600000 (have %v)", path, m.ValAttr)
	}
	return nil
}
