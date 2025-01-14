package pml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/dhx007/goffice"
)

type CT_OleObjectLink struct {
	// Update Linked Embedded Objects Automatically
	UpdateAutomaticAttr *bool
	ExtLst              *CT_ExtensionList
}

func NewCT_OleObjectLink() *CT_OleObjectLink {
	ret := &CT_OleObjectLink{}
	return ret
}

func (m *CT_OleObjectLink) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.UpdateAutomaticAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "updateAutomatic"},
			Value: fmt.Sprintf("%d", b2i(*m.UpdateAutomaticAttr))})
	}
	e.EncodeToken(start)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "p:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_OleObjectLink) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "updateAutomatic" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.UpdateAutomaticAttr = &parsed
			continue
		}
	}
lCT_OleObjectLink:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				goffice.Log("skipping unsupported element on CT_OleObjectLink %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_OleObjectLink
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_OleObjectLink and its children
func (m *CT_OleObjectLink) Validate() error {
	return m.ValidateWithPath("CT_OleObjectLink")
}

// ValidateWithPath validates the CT_OleObjectLink and its children, prefixing error messages with path
func (m *CT_OleObjectLink) ValidateWithPath(path string) error {
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
