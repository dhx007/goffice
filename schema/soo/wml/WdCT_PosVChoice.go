package wml

import (
	"encoding/xml"

	"github.com/dhx007/goffice"
)

type WdCT_PosVChoice struct {
	Align     WdST_AlignV
	PosOffset *int32
}

func NewWdCT_PosVChoice() *WdCT_PosVChoice {
	ret := &WdCT_PosVChoice{}
	return ret
}

func (m *WdCT_PosVChoice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.Align != WdST_AlignVUnset {
		sealign := xml.StartElement{Name: xml.Name{Local: "wp:align"}}
		e.EncodeElement(m.Align, sealign)
	}
	if m.PosOffset != nil {
		seposOffset := xml.StartElement{Name: xml.Name{Local: "wp:posOffset"}}
		e.EncodeElement(m.PosOffset, seposOffset)
	}
	return nil
}

func (m *WdCT_PosVChoice) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lWdCT_PosVChoice:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "align"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/wordprocessingDrawing", Local: "align"}:
				m.Align = WdST_AlignVUnset
				if err := d.DecodeElement(&m.Align, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "posOffset"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/wordprocessingDrawing", Local: "posOffset"}:
				m.PosOffset = new(int32)
				if err := d.DecodeElement(m.PosOffset, &el); err != nil {
					return err
				}
			default:
				goffice.Log("skipping unsupported element on WdCT_PosVChoice %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lWdCT_PosVChoice
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the WdCT_PosVChoice and its children
func (m *WdCT_PosVChoice) Validate() error {
	return m.ValidateWithPath("WdCT_PosVChoice")
}

// ValidateWithPath validates the WdCT_PosVChoice and its children, prefixing error messages with path
func (m *WdCT_PosVChoice) ValidateWithPath(path string) error {
	if err := m.Align.ValidateWithPath(path + "/Align"); err != nil {
		return err
	}
	return nil
}
