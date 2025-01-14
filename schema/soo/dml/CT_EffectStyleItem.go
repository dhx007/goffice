package dml

import (
	"encoding/xml"

	"github.com/dhx007/goffice"
)

type CT_EffectStyleItem struct {
	EffectLst *CT_EffectList
	EffectDag *CT_EffectContainer
	Scene3d   *CT_Scene3D
	Sp3d      *CT_Shape3D
}

func NewCT_EffectStyleItem() *CT_EffectStyleItem {
	ret := &CT_EffectStyleItem{}
	return ret
}

func (m *CT_EffectStyleItem) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.EffectLst != nil {
		seeffectLst := xml.StartElement{Name: xml.Name{Local: "a:effectLst"}}
		e.EncodeElement(m.EffectLst, seeffectLst)
	}
	if m.EffectDag != nil {
		seeffectDag := xml.StartElement{Name: xml.Name{Local: "a:effectDag"}}
		e.EncodeElement(m.EffectDag, seeffectDag)
	}
	if m.Scene3d != nil {
		sescene3d := xml.StartElement{Name: xml.Name{Local: "a:scene3d"}}
		e.EncodeElement(m.Scene3d, sescene3d)
	}
	if m.Sp3d != nil {
		sesp3d := xml.StartElement{Name: xml.Name{Local: "a:sp3d"}}
		e.EncodeElement(m.Sp3d, sesp3d)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_EffectStyleItem) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_EffectStyleItem:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "effectLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "effectLst"}:
				m.EffectLst = NewCT_EffectList()
				if err := d.DecodeElement(m.EffectLst, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "effectDag"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "effectDag"}:
				m.EffectDag = NewCT_EffectContainer()
				if err := d.DecodeElement(m.EffectDag, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "scene3d"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "scene3d"}:
				m.Scene3d = NewCT_Scene3D()
				if err := d.DecodeElement(m.Scene3d, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "sp3d"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "sp3d"}:
				m.Sp3d = NewCT_Shape3D()
				if err := d.DecodeElement(m.Sp3d, &el); err != nil {
					return err
				}
			default:
				goffice.Log("skipping unsupported element on CT_EffectStyleItem %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_EffectStyleItem
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_EffectStyleItem and its children
func (m *CT_EffectStyleItem) Validate() error {
	return m.ValidateWithPath("CT_EffectStyleItem")
}

// ValidateWithPath validates the CT_EffectStyleItem and its children, prefixing error messages with path
func (m *CT_EffectStyleItem) ValidateWithPath(path string) error {
	if m.EffectLst != nil {
		if err := m.EffectLst.ValidateWithPath(path + "/EffectLst"); err != nil {
			return err
		}
	}
	if m.EffectDag != nil {
		if err := m.EffectDag.ValidateWithPath(path + "/EffectDag"); err != nil {
			return err
		}
	}
	if m.Scene3d != nil {
		if err := m.Scene3d.ValidateWithPath(path + "/Scene3d"); err != nil {
			return err
		}
	}
	if m.Sp3d != nil {
		if err := m.Sp3d.ValidateWithPath(path + "/Sp3d"); err != nil {
			return err
		}
	}
	return nil
}
