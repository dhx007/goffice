package wml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/dhx007/goffice/schema/soo/ofc/sharedTypes"
)

type CT_TopPageBorder struct {
	TopLeftAttr  *string
	TopRightAttr *string
	IdAttr       *string
	// Border Style
	ValAttr ST_Border
	// Border Color
	ColorAttr *ST_HexColor
	// Border Theme Color
	ThemeColorAttr ST_ThemeColor
	// Border Theme Color Tint
	ThemeTintAttr *string
	// Border Theme Color Shade
	ThemeShadeAttr *string
	// Border Width
	SzAttr *uint64
	// Border Spacing Measurement
	SpaceAttr *uint64
	// Border Shadow
	ShadowAttr *sharedTypes.ST_OnOff
	// Create Frame Effect
	FrameAttr *sharedTypes.ST_OnOff
}

func NewCT_TopPageBorder() *CT_TopPageBorder {
	ret := &CT_TopPageBorder{}
	ret.ValAttr = ST_Border(1)
	return ret
}

func (m *CT_TopPageBorder) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.TopLeftAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r:topLeft"},
			Value: fmt.Sprintf("%v", *m.TopLeftAttr)})
	}
	if m.TopRightAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r:topRight"},
			Value: fmt.Sprintf("%v", *m.TopRightAttr)})
	}
	if m.IdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r:id"},
			Value: fmt.Sprintf("%v", *m.IdAttr)})
	}
	attr, err := m.ValAttr.MarshalXMLAttr(xml.Name{Local: "w:val"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	if m.ColorAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:color"},
			Value: fmt.Sprintf("%v", *m.ColorAttr)})
	}
	if m.ThemeColorAttr != ST_ThemeColorUnset {
		attr, err := m.ThemeColorAttr.MarshalXMLAttr(xml.Name{Local: "w:themeColor"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.ThemeTintAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:themeTint"},
			Value: fmt.Sprintf("%v", *m.ThemeTintAttr)})
	}
	if m.ThemeShadeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:themeShade"},
			Value: fmt.Sprintf("%v", *m.ThemeShadeAttr)})
	}
	if m.SzAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:sz"},
			Value: fmt.Sprintf("%v", *m.SzAttr)})
	}
	if m.SpaceAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:space"},
			Value: fmt.Sprintf("%v", *m.SpaceAttr)})
	}
	if m.ShadowAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:shadow"},
			Value: fmt.Sprintf("%v", *m.ShadowAttr)})
	}
	if m.FrameAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:frame"},
			Value: fmt.Sprintf("%v", *m.FrameAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TopPageBorder) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.ValAttr = ST_Border(1)
	for _, attr := range start.Attr {
		if attr.Name.Space == "http://schemas.openxmlformats.org/officeDocument/2006/relationships" && attr.Name.Local == "topLeft" ||
			attr.Name.Space == "http://purl.oclc.org/ooxml/officeDocument/relationships" && attr.Name.Local == "topLeft" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.TopLeftAttr = &parsed
			continue
		}
		if attr.Name.Space == "http://schemas.openxmlformats.org/officeDocument/2006/relationships" && attr.Name.Local == "topRight" ||
			attr.Name.Space == "http://purl.oclc.org/ooxml/officeDocument/relationships" && attr.Name.Local == "topRight" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.TopRightAttr = &parsed
			continue
		}
		if attr.Name.Space == "http://schemas.openxmlformats.org/officeDocument/2006/relationships" && attr.Name.Local == "id" ||
			attr.Name.Space == "http://purl.oclc.org/ooxml/officeDocument/relationships" && attr.Name.Local == "id" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IdAttr = &parsed
			continue
		}
		if attr.Name.Local == "val" {
			m.ValAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "color" {
			parsed, err := ParseUnionST_HexColor(attr.Value)
			if err != nil {
				return err
			}
			m.ColorAttr = &parsed
			continue
		}
		if attr.Name.Local == "themeColor" {
			m.ThemeColorAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "themeTint" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ThemeTintAttr = &parsed
			continue
		}
		if attr.Name.Local == "themeShade" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ThemeShadeAttr = &parsed
			continue
		}
		if attr.Name.Local == "sz" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.SzAttr = &parsed
			continue
		}
		if attr.Name.Local == "space" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.SpaceAttr = &parsed
			continue
		}
		if attr.Name.Local == "shadow" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.ShadowAttr = &parsed
			continue
		}
		if attr.Name.Local == "frame" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.FrameAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_TopPageBorder: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_TopPageBorder and its children
func (m *CT_TopPageBorder) Validate() error {
	return m.ValidateWithPath("CT_TopPageBorder")
}

// ValidateWithPath validates the CT_TopPageBorder and its children, prefixing error messages with path
func (m *CT_TopPageBorder) ValidateWithPath(path string) error {
	if m.ValAttr == ST_BorderUnset {
		return fmt.Errorf("%s/ValAttr is a mandatory field", path)
	}
	if err := m.ValAttr.ValidateWithPath(path + "/ValAttr"); err != nil {
		return err
	}
	if m.ColorAttr != nil {
		if err := m.ColorAttr.ValidateWithPath(path + "/ColorAttr"); err != nil {
			return err
		}
	}
	if err := m.ThemeColorAttr.ValidateWithPath(path + "/ThemeColorAttr"); err != nil {
		return err
	}
	if m.ShadowAttr != nil {
		if err := m.ShadowAttr.ValidateWithPath(path + "/ShadowAttr"); err != nil {
			return err
		}
	}
	if m.FrameAttr != nil {
		if err := m.FrameAttr.ValidateWithPath(path + "/FrameAttr"); err != nil {
			return err
		}
	}
	return nil
}
