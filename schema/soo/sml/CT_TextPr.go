package sml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/dhx007/goffice"
)

type CT_TextPr struct {
	// Prompt for File Name
	PromptAttr *bool
	// File Type
	FileTypeAttr ST_FileType
	// Code Page
	CodePageAttr *uint32
	// Character Set
	CharacterSetAttr *string
	// First Row
	FirstRowAttr *uint32
	// Source File Name
	SourceFileAttr *string
	// Delimited File
	DelimitedAttr *bool
	// Decimal Separator
	DecimalAttr *string
	// Thousands Separator
	ThousandsAttr *string
	// Tab as Delimiter
	TabAttr *bool
	// Space is Delimiter
	SpaceAttr *bool
	// Comma is Delimiter
	CommaAttr *bool
	// Semicolon is Delimiter
	SemicolonAttr *bool
	// Consecutive Delimiters
	ConsecutiveAttr *bool
	// Qualifier
	QualifierAttr ST_Qualifier
	// Custom Delimiter
	DelimiterAttr *string
	// Fields
	TextFields *CT_TextFields
}

func NewCT_TextPr() *CT_TextPr {
	ret := &CT_TextPr{}
	return ret
}

func (m *CT_TextPr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.PromptAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "prompt"},
			Value: fmt.Sprintf("%d", b2i(*m.PromptAttr))})
	}
	if m.FileTypeAttr != ST_FileTypeUnset {
		attr, err := m.FileTypeAttr.MarshalXMLAttr(xml.Name{Local: "fileType"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.CodePageAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "codePage"},
			Value: fmt.Sprintf("%v", *m.CodePageAttr)})
	}
	if m.CharacterSetAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "characterSet"},
			Value: fmt.Sprintf("%v", *m.CharacterSetAttr)})
	}
	if m.FirstRowAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "firstRow"},
			Value: fmt.Sprintf("%v", *m.FirstRowAttr)})
	}
	if m.SourceFileAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "sourceFile"},
			Value: fmt.Sprintf("%v", *m.SourceFileAttr)})
	}
	if m.DelimitedAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "delimited"},
			Value: fmt.Sprintf("%d", b2i(*m.DelimitedAttr))})
	}
	if m.DecimalAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "decimal"},
			Value: fmt.Sprintf("%v", *m.DecimalAttr)})
	}
	if m.ThousandsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "thousands"},
			Value: fmt.Sprintf("%v", *m.ThousandsAttr)})
	}
	if m.TabAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "tab"},
			Value: fmt.Sprintf("%d", b2i(*m.TabAttr))})
	}
	if m.SpaceAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "space"},
			Value: fmt.Sprintf("%d", b2i(*m.SpaceAttr))})
	}
	if m.CommaAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "comma"},
			Value: fmt.Sprintf("%d", b2i(*m.CommaAttr))})
	}
	if m.SemicolonAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "semicolon"},
			Value: fmt.Sprintf("%d", b2i(*m.SemicolonAttr))})
	}
	if m.ConsecutiveAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "consecutive"},
			Value: fmt.Sprintf("%d", b2i(*m.ConsecutiveAttr))})
	}
	if m.QualifierAttr != ST_QualifierUnset {
		attr, err := m.QualifierAttr.MarshalXMLAttr(xml.Name{Local: "qualifier"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.DelimiterAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "delimiter"},
			Value: fmt.Sprintf("%v", *m.DelimiterAttr)})
	}
	e.EncodeToken(start)
	if m.TextFields != nil {
		setextFields := xml.StartElement{Name: xml.Name{Local: "ma:textFields"}}
		e.EncodeElement(m.TextFields, setextFields)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TextPr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "thousands" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ThousandsAttr = &parsed
			continue
		}
		if attr.Name.Local == "tab" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.TabAttr = &parsed
			continue
		}
		if attr.Name.Local == "fileType" {
			m.FileTypeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "space" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.SpaceAttr = &parsed
			continue
		}
		if attr.Name.Local == "characterSet" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.CharacterSetAttr = &parsed
			continue
		}
		if attr.Name.Local == "sourceFile" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.SourceFileAttr = &parsed
			continue
		}
		if attr.Name.Local == "delimited" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.DelimitedAttr = &parsed
			continue
		}
		if attr.Name.Local == "decimal" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.DecimalAttr = &parsed
			continue
		}
		if attr.Name.Local == "prompt" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.PromptAttr = &parsed
			continue
		}
		if attr.Name.Local == "codePage" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.CodePageAttr = &pt
			continue
		}
		if attr.Name.Local == "firstRow" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.FirstRowAttr = &pt
			continue
		}
		if attr.Name.Local == "comma" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.CommaAttr = &parsed
			continue
		}
		if attr.Name.Local == "semicolon" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.SemicolonAttr = &parsed
			continue
		}
		if attr.Name.Local == "consecutive" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ConsecutiveAttr = &parsed
			continue
		}
		if attr.Name.Local == "qualifier" {
			m.QualifierAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "delimiter" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.DelimiterAttr = &parsed
			continue
		}
	}
lCT_TextPr:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "textFields"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "textFields"}:
				m.TextFields = NewCT_TextFields()
				if err := d.DecodeElement(m.TextFields, &el); err != nil {
					return err
				}
			default:
				goffice.Log("skipping unsupported element on CT_TextPr %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TextPr
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TextPr and its children
func (m *CT_TextPr) Validate() error {
	return m.ValidateWithPath("CT_TextPr")
}

// ValidateWithPath validates the CT_TextPr and its children, prefixing error messages with path
func (m *CT_TextPr) ValidateWithPath(path string) error {
	if err := m.FileTypeAttr.ValidateWithPath(path + "/FileTypeAttr"); err != nil {
		return err
	}
	if err := m.QualifierAttr.ValidateWithPath(path + "/QualifierAttr"); err != nil {
		return err
	}
	if m.TextFields != nil {
		if err := m.TextFields.ValidateWithPath(path + "/TextFields"); err != nil {
			return err
		}
	}
	return nil
}
