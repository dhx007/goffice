package wml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/dhx007/goffice/schema/soo/ofc/sharedTypes"
)

type CT_ReadingModeInkLockDown struct {
	// Use Actual Pages, Not Virtual Pages
	ActualPgAttr sharedTypes.ST_OnOff
	// Virtual Page Width
	WAttr uint64
	// Virtual Page Height
	HAttr uint64
	// Font Size Scaling
	FontSzAttr ST_DecimalNumberOrPercent
}

func NewCT_ReadingModeInkLockDown() *CT_ReadingModeInkLockDown {
	ret := &CT_ReadingModeInkLockDown{}
	return ret
}

func (m *CT_ReadingModeInkLockDown) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:actualPg"},
		Value: fmt.Sprintf("%v", m.ActualPgAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:w"},
		Value: fmt.Sprintf("%v", m.WAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:h"},
		Value: fmt.Sprintf("%v", m.HAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:fontSz"},
		Value: fmt.Sprintf("%v", m.FontSzAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ReadingModeInkLockDown) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "actualPg" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.ActualPgAttr = parsed
			continue
		}
		if attr.Name.Local == "w" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.WAttr = parsed
			continue
		}
		if attr.Name.Local == "h" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.HAttr = parsed
			continue
		}
		if attr.Name.Local == "fontSz" {
			parsed, err := ParseUnionST_DecimalNumberOrPercent(attr.Value)
			if err != nil {
				return err
			}
			m.FontSzAttr = parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_ReadingModeInkLockDown: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_ReadingModeInkLockDown and its children
func (m *CT_ReadingModeInkLockDown) Validate() error {
	return m.ValidateWithPath("CT_ReadingModeInkLockDown")
}

// ValidateWithPath validates the CT_ReadingModeInkLockDown and its children, prefixing error messages with path
func (m *CT_ReadingModeInkLockDown) ValidateWithPath(path string) error {
	if err := m.ActualPgAttr.ValidateWithPath(path + "/ActualPgAttr"); err != nil {
		return err
	}
	if err := m.FontSzAttr.ValidateWithPath(path + "/FontSzAttr"); err != nil {
		return err
	}
	return nil
}
