package pml

import (
	"encoding/xml"
	"fmt"

	"github.com/dhx007/goffice"
)

type CT_SlideMasterIdList struct {
	// Slide Master ID
	SldMasterId []*CT_SlideMasterIdListEntry
}

func NewCT_SlideMasterIdList() *CT_SlideMasterIdList {
	ret := &CT_SlideMasterIdList{}
	return ret
}

func (m *CT_SlideMasterIdList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.SldMasterId != nil {
		sesldMasterId := xml.StartElement{Name: xml.Name{Local: "p:sldMasterId"}}
		for _, c := range m.SldMasterId {
			e.EncodeElement(c, sesldMasterId)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_SlideMasterIdList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_SlideMasterIdList:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "sldMasterId"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "sldMasterId"}:
				tmp := NewCT_SlideMasterIdListEntry()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.SldMasterId = append(m.SldMasterId, tmp)
			default:
				goffice.Log("skipping unsupported element on CT_SlideMasterIdList %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_SlideMasterIdList
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_SlideMasterIdList and its children
func (m *CT_SlideMasterIdList) Validate() error {
	return m.ValidateWithPath("CT_SlideMasterIdList")
}

// ValidateWithPath validates the CT_SlideMasterIdList and its children, prefixing error messages with path
func (m *CT_SlideMasterIdList) ValidateWithPath(path string) error {
	for i, v := range m.SldMasterId {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/SldMasterId[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
