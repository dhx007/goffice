package pml

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type CT_ShowInfoBrowse struct {
	// Show Scroll Bar in Window
	ShowScrollbarAttr *bool
}

func NewCT_ShowInfoBrowse() *CT_ShowInfoBrowse {
	ret := &CT_ShowInfoBrowse{}
	return ret
}

func (m *CT_ShowInfoBrowse) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.ShowScrollbarAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showScrollbar"},
			Value: fmt.Sprintf("%d", b2i(*m.ShowScrollbarAttr))})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ShowInfoBrowse) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "showScrollbar" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowScrollbarAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_ShowInfoBrowse: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_ShowInfoBrowse and its children
func (m *CT_ShowInfoBrowse) Validate() error {
	return m.ValidateWithPath("CT_ShowInfoBrowse")
}

// ValidateWithPath validates the CT_ShowInfoBrowse and its children, prefixing error messages with path
func (m *CT_ShowInfoBrowse) ValidateWithPath(path string) error {
	return nil
}
