package wml

import (
	"encoding/xml"
	"fmt"

	"github.com/dhx007/goffice/schema/soo/ofc/sharedTypes"
)

type CT_LvlLegacy struct {
	// Use Legacy Numbering Properties
	LegacyAttr *sharedTypes.ST_OnOff
	// Legacy Spacing
	LegacySpaceAttr *sharedTypes.ST_TwipsMeasure
	// Legacy Indent
	LegacyIndentAttr *ST_SignedTwipsMeasure
}

func NewCT_LvlLegacy() *CT_LvlLegacy {
	ret := &CT_LvlLegacy{}
	return ret
}

func (m *CT_LvlLegacy) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.LegacyAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:legacy"},
			Value: fmt.Sprintf("%v", *m.LegacyAttr)})
	}
	if m.LegacySpaceAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:legacySpace"},
			Value: fmt.Sprintf("%v", *m.LegacySpaceAttr)})
	}
	if m.LegacyIndentAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:legacyIndent"},
			Value: fmt.Sprintf("%v", *m.LegacyIndentAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_LvlLegacy) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "legacy" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.LegacyAttr = &parsed
			continue
		}
		if attr.Name.Local == "legacySpace" {
			parsed, err := ParseUnionST_TwipsMeasure(attr.Value)
			if err != nil {
				return err
			}
			m.LegacySpaceAttr = &parsed
			continue
		}
		if attr.Name.Local == "legacyIndent" {
			parsed, err := ParseUnionST_SignedTwipsMeasure(attr.Value)
			if err != nil {
				return err
			}
			m.LegacyIndentAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_LvlLegacy: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_LvlLegacy and its children
func (m *CT_LvlLegacy) Validate() error {
	return m.ValidateWithPath("CT_LvlLegacy")
}

// ValidateWithPath validates the CT_LvlLegacy and its children, prefixing error messages with path
func (m *CT_LvlLegacy) ValidateWithPath(path string) error {
	if m.LegacyAttr != nil {
		if err := m.LegacyAttr.ValidateWithPath(path + "/LegacyAttr"); err != nil {
			return err
		}
	}
	if m.LegacySpaceAttr != nil {
		if err := m.LegacySpaceAttr.ValidateWithPath(path + "/LegacySpaceAttr"); err != nil {
			return err
		}
	}
	if m.LegacyIndentAttr != nil {
		if err := m.LegacyIndentAttr.ValidateWithPath(path + "/LegacyIndentAttr"); err != nil {
			return err
		}
	}
	return nil
}
