package content_types

import (
	"encoding/xml"
	"fmt"
)

type CT_Override struct {
	ContentTypeAttr string
	PartNameAttr    string
}

func NewCT_Override() *CT_Override {
	ret := &CT_Override{}
	ret.ContentTypeAttr = "application/xml"
	return ret
}

func (m *CT_Override) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "ContentType"},
		Value: fmt.Sprintf("%v", m.ContentTypeAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "PartName"},
		Value: fmt.Sprintf("%v", m.PartNameAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Override) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.ContentTypeAttr = "application/xml"
	for _, attr := range start.Attr {
		if attr.Name.Local == "ContentType" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ContentTypeAttr = parsed
			continue
		}
		if attr.Name.Local == "PartName" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.PartNameAttr = parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_Override: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_Override and its children
func (m *CT_Override) Validate() error {
	return m.ValidateWithPath("CT_Override")
}

// ValidateWithPath validates the CT_Override and its children, prefixing error messages with path
func (m *CT_Override) ValidateWithPath(path string) error {
	if !ST_ContentTypePatternRe.MatchString(m.ContentTypeAttr) {
		return fmt.Errorf(`%s/m.ContentTypeAttr must match '%s' (have %v)`, path, ST_ContentTypePatternRe, m.ContentTypeAttr)
	}
	return nil
}
