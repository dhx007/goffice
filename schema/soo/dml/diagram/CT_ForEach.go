package diagram

import (
	"encoding/xml"
	"fmt"

	"github.com/dhx007/goffice"
	"github.com/dhx007/goffice/schema/soo/dml"
)

type CT_ForEach struct {
	NameAttr          *string
	RefAttr           *string
	Alg               []*CT_Algorithm
	Shape             []*CT_Shape
	PresOf            []*CT_PresentationOf
	ConstrLst         []*CT_Constraints
	RuleLst           []*CT_Rules
	ForEach           []*CT_ForEach
	LayoutNode        []*CT_LayoutNode
	Choose            []*CT_Choose
	ExtLst            []*dml.CT_OfficeArtExtensionList
	AxisAttr          *ST_AxisTypes
	PtTypeAttr        *ST_ElementTypes
	HideLastTransAttr *ST_Booleans
	StAttr            *ST_Ints
	CntAttr           *ST_UnsignedInts
	StepAttr          *ST_Ints
}

func NewCT_ForEach() *CT_ForEach {
	ret := &CT_ForEach{}
	return ret
}

func (m *CT_ForEach) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.NameAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "name"},
			Value: fmt.Sprintf("%v", *m.NameAttr)})
	}
	if m.RefAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "ref"},
			Value: fmt.Sprintf("%v", *m.RefAttr)})
	}
	if m.AxisAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "axis"},
			Value: fmt.Sprintf("%v", *m.AxisAttr)})
	}
	if m.PtTypeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "ptType"},
			Value: fmt.Sprintf("%v", *m.PtTypeAttr)})
	}
	if m.HideLastTransAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "hideLastTrans"},
			Value: fmt.Sprintf("%v", *m.HideLastTransAttr)})
	}
	if m.StAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "st"},
			Value: fmt.Sprintf("%v", *m.StAttr)})
	}
	if m.CntAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "cnt"},
			Value: fmt.Sprintf("%v", *m.CntAttr)})
	}
	if m.StepAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "step"},
			Value: fmt.Sprintf("%v", *m.StepAttr)})
	}
	e.EncodeToken(start)
	if m.Alg != nil {
		sealg := xml.StartElement{Name: xml.Name{Local: "alg"}}
		for _, c := range m.Alg {
			e.EncodeElement(c, sealg)
		}
	}
	if m.Shape != nil {
		seshape := xml.StartElement{Name: xml.Name{Local: "shape"}}
		for _, c := range m.Shape {
			e.EncodeElement(c, seshape)
		}
	}
	if m.PresOf != nil {
		sepresOf := xml.StartElement{Name: xml.Name{Local: "presOf"}}
		for _, c := range m.PresOf {
			e.EncodeElement(c, sepresOf)
		}
	}
	if m.ConstrLst != nil {
		seconstrLst := xml.StartElement{Name: xml.Name{Local: "constrLst"}}
		for _, c := range m.ConstrLst {
			e.EncodeElement(c, seconstrLst)
		}
	}
	if m.RuleLst != nil {
		seruleLst := xml.StartElement{Name: xml.Name{Local: "ruleLst"}}
		for _, c := range m.RuleLst {
			e.EncodeElement(c, seruleLst)
		}
	}
	if m.ForEach != nil {
		seforEach := xml.StartElement{Name: xml.Name{Local: "forEach"}}
		for _, c := range m.ForEach {
			e.EncodeElement(c, seforEach)
		}
	}
	if m.LayoutNode != nil {
		selayoutNode := xml.StartElement{Name: xml.Name{Local: "layoutNode"}}
		for _, c := range m.LayoutNode {
			e.EncodeElement(c, selayoutNode)
		}
	}
	if m.Choose != nil {
		sechoose := xml.StartElement{Name: xml.Name{Local: "choose"}}
		for _, c := range m.Choose {
			e.EncodeElement(c, sechoose)
		}
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "extLst"}}
		for _, c := range m.ExtLst {
			e.EncodeElement(c, seextLst)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ForEach) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "ref" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RefAttr = &parsed
			continue
		}
		if attr.Name.Local == "name" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NameAttr = &parsed
			continue
		}
		if attr.Name.Local == "axis" {
			parsed, err := ParseSliceST_AxisTypes(attr.Value)
			if err != nil {
				return err
			}
			m.AxisAttr = &parsed
			continue
		}
		if attr.Name.Local == "ptType" {
			parsed, err := ParseSliceST_ElementTypes(attr.Value)
			if err != nil {
				return err
			}
			m.PtTypeAttr = &parsed
			continue
		}
		if attr.Name.Local == "hideLastTrans" {
			parsed, err := ParseSliceST_Booleans(attr.Value)
			if err != nil {
				return err
			}
			m.HideLastTransAttr = &parsed
			continue
		}
		if attr.Name.Local == "st" {
			parsed, err := ParseSliceST_Ints(attr.Value)
			if err != nil {
				return err
			}
			m.StAttr = &parsed
			continue
		}
		if attr.Name.Local == "cnt" {
			parsed, err := ParseSliceST_UnsignedInts(attr.Value)
			if err != nil {
				return err
			}
			m.CntAttr = &parsed
			continue
		}
		if attr.Name.Local == "step" {
			parsed, err := ParseSliceST_Ints(attr.Value)
			if err != nil {
				return err
			}
			m.StepAttr = &parsed
			continue
		}
	}
lCT_ForEach:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "alg"}:
				tmp := NewCT_Algorithm()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Alg = append(m.Alg, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "shape"}:
				tmp := NewCT_Shape()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Shape = append(m.Shape, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "presOf"}:
				tmp := NewCT_PresentationOf()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.PresOf = append(m.PresOf, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "constrLst"}:
				tmp := NewCT_Constraints()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.ConstrLst = append(m.ConstrLst, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "ruleLst"}:
				tmp := NewCT_Rules()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.RuleLst = append(m.RuleLst, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "forEach"}:
				tmp := NewCT_ForEach()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.ForEach = append(m.ForEach, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "layoutNode"}:
				tmp := NewCT_LayoutNode()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.LayoutNode = append(m.LayoutNode, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "choose"}:
				tmp := NewCT_Choose()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Choose = append(m.Choose, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "extLst"}:
				tmp := dml.NewCT_OfficeArtExtensionList()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.ExtLst = append(m.ExtLst, tmp)
			default:
				goffice.Log("skipping unsupported element on CT_ForEach %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ForEach
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_ForEach and its children
func (m *CT_ForEach) Validate() error {
	return m.ValidateWithPath("CT_ForEach")
}

// ValidateWithPath validates the CT_ForEach and its children, prefixing error messages with path
func (m *CT_ForEach) ValidateWithPath(path string) error {
	for i, v := range m.Alg {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Alg[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Shape {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Shape[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.PresOf {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/PresOf[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.ConstrLst {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/ConstrLst[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.RuleLst {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/RuleLst[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.ForEach {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/ForEach[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.LayoutNode {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/LayoutNode[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Choose {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Choose[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.ExtLst {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/ExtLst[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
