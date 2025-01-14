package wml

import (
	"encoding/xml"
	"fmt"

	"github.com/dhx007/goffice"
)

type WdCT_TxbxContent struct {
	// Anchor for Imported External Content
	AltChunk               []*CT_AltChunk
	EG_ContentBlockContent []*EG_ContentBlockContent
}

func NewWdCT_TxbxContent() *WdCT_TxbxContent {
	ret := &WdCT_TxbxContent{}
	return ret
}

func (m *WdCT_TxbxContent) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.AltChunk != nil {
		sealtChunk := xml.StartElement{Name: xml.Name{Local: "wp:altChunk"}}
		for _, c := range m.AltChunk {
			e.EncodeElement(c, sealtChunk)
		}
	}
	if m.EG_ContentBlockContent != nil {
		for _, c := range m.EG_ContentBlockContent {
			c.MarshalXML(e, xml.StartElement{})
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *WdCT_TxbxContent) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lWdCT_TxbxContent:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "altChunk"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "altChunk"}:
				tmp := NewCT_AltChunk()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.AltChunk = append(m.AltChunk, tmp)
			default:
				goffice.Log("skipping unsupported element on WdCT_TxbxContent %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lWdCT_TxbxContent
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the WdCT_TxbxContent and its children
func (m *WdCT_TxbxContent) Validate() error {
	return m.ValidateWithPath("WdCT_TxbxContent")
}

// ValidateWithPath validates the WdCT_TxbxContent and its children, prefixing error messages with path
func (m *WdCT_TxbxContent) ValidateWithPath(path string) error {
	for i, v := range m.AltChunk {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/AltChunk[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.EG_ContentBlockContent {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/EG_ContentBlockContent[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
