package sml

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type AG_AutoFormat struct {
	AutoFormatIdAttr            *uint32
	ApplyNumberFormatsAttr      *bool
	ApplyBorderFormatsAttr      *bool
	ApplyFontFormatsAttr        *bool
	ApplyPatternFormatsAttr     *bool
	ApplyAlignmentFormatsAttr   *bool
	ApplyWidthHeightFormatsAttr *bool
}

func NewAG_AutoFormat() *AG_AutoFormat {
	ret := &AG_AutoFormat{}
	return ret
}

func (m *AG_AutoFormat) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.AutoFormatIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "autoFormatId"},
			Value: fmt.Sprintf("%v", *m.AutoFormatIdAttr)})
	}
	if m.ApplyNumberFormatsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "applyNumberFormats"},
			Value: fmt.Sprintf("%d", b2i(*m.ApplyNumberFormatsAttr))})
	}
	if m.ApplyBorderFormatsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "applyBorderFormats"},
			Value: fmt.Sprintf("%d", b2i(*m.ApplyBorderFormatsAttr))})
	}
	if m.ApplyFontFormatsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "applyFontFormats"},
			Value: fmt.Sprintf("%d", b2i(*m.ApplyFontFormatsAttr))})
	}
	if m.ApplyPatternFormatsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "applyPatternFormats"},
			Value: fmt.Sprintf("%d", b2i(*m.ApplyPatternFormatsAttr))})
	}
	if m.ApplyAlignmentFormatsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "applyAlignmentFormats"},
			Value: fmt.Sprintf("%d", b2i(*m.ApplyAlignmentFormatsAttr))})
	}
	if m.ApplyWidthHeightFormatsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "applyWidthHeightFormats"},
			Value: fmt.Sprintf("%d", b2i(*m.ApplyWidthHeightFormatsAttr))})
	}
	return nil
}

func (m *AG_AutoFormat) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "autoFormatId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.AutoFormatIdAttr = &pt
			continue
		}
		if attr.Name.Local == "applyNumberFormats" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ApplyNumberFormatsAttr = &parsed
			continue
		}
		if attr.Name.Local == "applyBorderFormats" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ApplyBorderFormatsAttr = &parsed
			continue
		}
		if attr.Name.Local == "applyFontFormats" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ApplyFontFormatsAttr = &parsed
			continue
		}
		if attr.Name.Local == "applyPatternFormats" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ApplyPatternFormatsAttr = &parsed
			continue
		}
		if attr.Name.Local == "applyAlignmentFormats" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ApplyAlignmentFormatsAttr = &parsed
			continue
		}
		if attr.Name.Local == "applyWidthHeightFormats" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ApplyWidthHeightFormatsAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing AG_AutoFormat: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the AG_AutoFormat and its children
func (m *AG_AutoFormat) Validate() error {
	return m.ValidateWithPath("AG_AutoFormat")
}

// ValidateWithPath validates the AG_AutoFormat and its children, prefixing error messages with path
func (m *AG_AutoFormat) ValidateWithPath(path string) error {
	return nil
}
