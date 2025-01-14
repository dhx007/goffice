package dml

import (
	"encoding/xml"
	"fmt"
)

type CT_AlphaBiLevelEffect struct {
	ThreshAttr ST_PositiveFixedPercentage
}

func NewCT_AlphaBiLevelEffect() *CT_AlphaBiLevelEffect {
	ret := &CT_AlphaBiLevelEffect{}
	return ret
}

func (m *CT_AlphaBiLevelEffect) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "thresh"},
		Value: fmt.Sprintf("%v", m.ThreshAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_AlphaBiLevelEffect) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "thresh" {
			parsed, err := ParseUnionST_PositiveFixedPercentage(attr.Value)
			if err != nil {
				return err
			}
			m.ThreshAttr = parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_AlphaBiLevelEffect: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_AlphaBiLevelEffect and its children
func (m *CT_AlphaBiLevelEffect) Validate() error {
	return m.ValidateWithPath("CT_AlphaBiLevelEffect")
}

// ValidateWithPath validates the CT_AlphaBiLevelEffect and its children, prefixing error messages with path
func (m *CT_AlphaBiLevelEffect) ValidateWithPath(path string) error {
	if err := m.ThreshAttr.ValidateWithPath(path + "/ThreshAttr"); err != nil {
		return err
	}
	return nil
}
