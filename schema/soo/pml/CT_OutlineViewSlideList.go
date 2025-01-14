package pml

import (
	"encoding/xml"
	"fmt"

	"github.com/dhx007/goffice"
)

type CT_OutlineViewSlideList struct {
	// Presentation Slide
	Sld []*CT_OutlineViewSlideEntry
}

func NewCT_OutlineViewSlideList() *CT_OutlineViewSlideList {
	ret := &CT_OutlineViewSlideList{}
	return ret
}

func (m *CT_OutlineViewSlideList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.Sld != nil {
		sesld := xml.StartElement{Name: xml.Name{Local: "p:sld"}}
		for _, c := range m.Sld {
			e.EncodeElement(c, sesld)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_OutlineViewSlideList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_OutlineViewSlideList:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "sld"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "sld"}:
				tmp := NewCT_OutlineViewSlideEntry()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Sld = append(m.Sld, tmp)
			default:
				goffice.Log("skipping unsupported element on CT_OutlineViewSlideList %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_OutlineViewSlideList
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_OutlineViewSlideList and its children
func (m *CT_OutlineViewSlideList) Validate() error {
	return m.ValidateWithPath("CT_OutlineViewSlideList")
}

// ValidateWithPath validates the CT_OutlineViewSlideList and its children, prefixing error messages with path
func (m *CT_OutlineViewSlideList) ValidateWithPath(path string) error {
	for i, v := range m.Sld {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Sld[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
