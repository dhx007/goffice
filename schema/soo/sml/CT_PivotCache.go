package sml

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type CT_PivotCache struct {
	// PivotCache Id
	CacheIdAttr uint32
	IdAttr      string
}

func NewCT_PivotCache() *CT_PivotCache {
	ret := &CT_PivotCache{}
	return ret
}

func (m *CT_PivotCache) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "cacheId"},
		Value: fmt.Sprintf("%v", m.CacheIdAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r:id"},
		Value: fmt.Sprintf("%v", m.IdAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_PivotCache) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Space == "http://schemas.openxmlformats.org/officeDocument/2006/relationships" && attr.Name.Local == "id" ||
			attr.Name.Space == "http://purl.oclc.org/ooxml/officeDocument/relationships" && attr.Name.Local == "id" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IdAttr = parsed
			continue
		}
		if attr.Name.Local == "cacheId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.CacheIdAttr = uint32(parsed)
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_PivotCache: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_PivotCache and its children
func (m *CT_PivotCache) Validate() error {
	return m.ValidateWithPath("CT_PivotCache")
}

// ValidateWithPath validates the CT_PivotCache and its children, prefixing error messages with path
func (m *CT_PivotCache) ValidateWithPath(path string) error {
	return nil
}
