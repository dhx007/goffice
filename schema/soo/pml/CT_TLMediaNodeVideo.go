package pml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/dhx007/goffice"
)

type CT_TLMediaNodeVideo struct {
	// Full Screen
	FullScrnAttr *bool
	// Common Media Node Properties
	CMediaNode *CT_TLCommonMediaNodeData
}

func NewCT_TLMediaNodeVideo() *CT_TLMediaNodeVideo {
	ret := &CT_TLMediaNodeVideo{}
	ret.CMediaNode = NewCT_TLCommonMediaNodeData()
	return ret
}

func (m *CT_TLMediaNodeVideo) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.FullScrnAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "fullScrn"},
			Value: fmt.Sprintf("%d", b2i(*m.FullScrnAttr))})
	}
	e.EncodeToken(start)
	secMediaNode := xml.StartElement{Name: xml.Name{Local: "p:cMediaNode"}}
	e.EncodeElement(m.CMediaNode, secMediaNode)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TLMediaNodeVideo) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CMediaNode = NewCT_TLCommonMediaNodeData()
	for _, attr := range start.Attr {
		if attr.Name.Local == "fullScrn" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.FullScrnAttr = &parsed
			continue
		}
	}
lCT_TLMediaNodeVideo:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "cMediaNode"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "cMediaNode"}:
				if err := d.DecodeElement(m.CMediaNode, &el); err != nil {
					return err
				}
			default:
				goffice.Log("skipping unsupported element on CT_TLMediaNodeVideo %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TLMediaNodeVideo
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TLMediaNodeVideo and its children
func (m *CT_TLMediaNodeVideo) Validate() error {
	return m.ValidateWithPath("CT_TLMediaNodeVideo")
}

// ValidateWithPath validates the CT_TLMediaNodeVideo and its children, prefixing error messages with path
func (m *CT_TLMediaNodeVideo) ValidateWithPath(path string) error {
	if err := m.CMediaNode.ValidateWithPath(path + "/CMediaNode"); err != nil {
		return err
	}
	return nil
}
