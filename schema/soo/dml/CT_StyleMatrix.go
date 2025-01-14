package dml

import (
	"encoding/xml"
	"fmt"

	"github.com/dhx007/goffice"
)

type CT_StyleMatrix struct {
	NameAttr       *string
	FillStyleLst   *CT_FillStyleList
	LnStyleLst     *CT_LineStyleList
	EffectStyleLst *CT_EffectStyleList
	BgFillStyleLst *CT_BackgroundFillStyleList
}

func NewCT_StyleMatrix() *CT_StyleMatrix {
	ret := &CT_StyleMatrix{}
	ret.FillStyleLst = NewCT_FillStyleList()
	ret.LnStyleLst = NewCT_LineStyleList()
	ret.EffectStyleLst = NewCT_EffectStyleList()
	ret.BgFillStyleLst = NewCT_BackgroundFillStyleList()
	return ret
}

func (m *CT_StyleMatrix) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.NameAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "name"},
			Value: fmt.Sprintf("%v", *m.NameAttr)})
	}
	e.EncodeToken(start)
	sefillStyleLst := xml.StartElement{Name: xml.Name{Local: "a:fillStyleLst"}}
	e.EncodeElement(m.FillStyleLst, sefillStyleLst)
	selnStyleLst := xml.StartElement{Name: xml.Name{Local: "a:lnStyleLst"}}
	e.EncodeElement(m.LnStyleLst, selnStyleLst)
	seeffectStyleLst := xml.StartElement{Name: xml.Name{Local: "a:effectStyleLst"}}
	e.EncodeElement(m.EffectStyleLst, seeffectStyleLst)
	sebgFillStyleLst := xml.StartElement{Name: xml.Name{Local: "a:bgFillStyleLst"}}
	e.EncodeElement(m.BgFillStyleLst, sebgFillStyleLst)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_StyleMatrix) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.FillStyleLst = NewCT_FillStyleList()
	m.LnStyleLst = NewCT_LineStyleList()
	m.EffectStyleLst = NewCT_EffectStyleList()
	m.BgFillStyleLst = NewCT_BackgroundFillStyleList()
	for _, attr := range start.Attr {
		if attr.Name.Local == "name" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NameAttr = &parsed
			continue
		}
	}
lCT_StyleMatrix:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "fillStyleLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "fillStyleLst"}:
				if err := d.DecodeElement(m.FillStyleLst, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "lnStyleLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "lnStyleLst"}:
				if err := d.DecodeElement(m.LnStyleLst, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "effectStyleLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "effectStyleLst"}:
				if err := d.DecodeElement(m.EffectStyleLst, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "bgFillStyleLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "bgFillStyleLst"}:
				if err := d.DecodeElement(m.BgFillStyleLst, &el); err != nil {
					return err
				}
			default:
				goffice.Log("skipping unsupported element on CT_StyleMatrix %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_StyleMatrix
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_StyleMatrix and its children
func (m *CT_StyleMatrix) Validate() error {
	return m.ValidateWithPath("CT_StyleMatrix")
}

// ValidateWithPath validates the CT_StyleMatrix and its children, prefixing error messages with path
func (m *CT_StyleMatrix) ValidateWithPath(path string) error {
	if err := m.FillStyleLst.ValidateWithPath(path + "/FillStyleLst"); err != nil {
		return err
	}
	if err := m.LnStyleLst.ValidateWithPath(path + "/LnStyleLst"); err != nil {
		return err
	}
	if err := m.EffectStyleLst.ValidateWithPath(path + "/EffectStyleLst"); err != nil {
		return err
	}
	if err := m.BgFillStyleLst.ValidateWithPath(path + "/BgFillStyleLst"); err != nil {
		return err
	}
	return nil
}
