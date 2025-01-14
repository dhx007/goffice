package pml

import (
	"encoding/xml"
	"fmt"

	"github.com/dhx007/goffice"
)

type CT_TLCommonBehaviorData struct {
	// Additive
	AdditiveAttr ST_TLBehaviorAdditiveType
	// Accumulate
	AccumulateAttr ST_TLBehaviorAccumulateType
	// Transform Type
	XfrmTypeAttr ST_TLBehaviorTransformType
	// From
	FromAttr *string
	// To
	ToAttr *string
	// By
	ByAttr *string
	// Runtime Context
	RctxAttr *string
	// Override
	OverrideAttr ST_TLBehaviorOverrideType
	CTn          *CT_TLCommonTimeNodeData
	// Target Element
	TgtEl *CT_TLTimeTargetElement
	// Attribute Name List
	AttrNameLst *CT_TLBehaviorAttributeNameList
}

func NewCT_TLCommonBehaviorData() *CT_TLCommonBehaviorData {
	ret := &CT_TLCommonBehaviorData{}
	ret.CTn = NewCT_TLCommonTimeNodeData()
	ret.TgtEl = NewCT_TLTimeTargetElement()
	return ret
}

func (m *CT_TLCommonBehaviorData) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.AdditiveAttr != ST_TLBehaviorAdditiveTypeUnset {
		attr, err := m.AdditiveAttr.MarshalXMLAttr(xml.Name{Local: "additive"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.AccumulateAttr != ST_TLBehaviorAccumulateTypeUnset {
		attr, err := m.AccumulateAttr.MarshalXMLAttr(xml.Name{Local: "accumulate"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.XfrmTypeAttr != ST_TLBehaviorTransformTypeUnset {
		attr, err := m.XfrmTypeAttr.MarshalXMLAttr(xml.Name{Local: "xfrmType"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.FromAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "from"},
			Value: fmt.Sprintf("%v", *m.FromAttr)})
	}
	if m.ToAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "to"},
			Value: fmt.Sprintf("%v", *m.ToAttr)})
	}
	if m.ByAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "by"},
			Value: fmt.Sprintf("%v", *m.ByAttr)})
	}
	if m.RctxAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "rctx"},
			Value: fmt.Sprintf("%v", *m.RctxAttr)})
	}
	if m.OverrideAttr != ST_TLBehaviorOverrideTypeUnset {
		attr, err := m.OverrideAttr.MarshalXMLAttr(xml.Name{Local: "override"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	secTn := xml.StartElement{Name: xml.Name{Local: "p:cTn"}}
	e.EncodeElement(m.CTn, secTn)
	setgtEl := xml.StartElement{Name: xml.Name{Local: "p:tgtEl"}}
	e.EncodeElement(m.TgtEl, setgtEl)
	if m.AttrNameLst != nil {
		seattrNameLst := xml.StartElement{Name: xml.Name{Local: "p:attrNameLst"}}
		e.EncodeElement(m.AttrNameLst, seattrNameLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TLCommonBehaviorData) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CTn = NewCT_TLCommonTimeNodeData()
	m.TgtEl = NewCT_TLTimeTargetElement()
	for _, attr := range start.Attr {
		if attr.Name.Local == "additive" {
			m.AdditiveAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "accumulate" {
			m.AccumulateAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "xfrmType" {
			m.XfrmTypeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "from" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.FromAttr = &parsed
			continue
		}
		if attr.Name.Local == "to" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ToAttr = &parsed
			continue
		}
		if attr.Name.Local == "by" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ByAttr = &parsed
			continue
		}
		if attr.Name.Local == "rctx" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RctxAttr = &parsed
			continue
		}
		if attr.Name.Local == "override" {
			m.OverrideAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
lCT_TLCommonBehaviorData:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "cTn"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "cTn"}:
				if err := d.DecodeElement(m.CTn, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "tgtEl"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "tgtEl"}:
				if err := d.DecodeElement(m.TgtEl, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "attrNameLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "attrNameLst"}:
				m.AttrNameLst = NewCT_TLBehaviorAttributeNameList()
				if err := d.DecodeElement(m.AttrNameLst, &el); err != nil {
					return err
				}
			default:
				goffice.Log("skipping unsupported element on CT_TLCommonBehaviorData %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TLCommonBehaviorData
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TLCommonBehaviorData and its children
func (m *CT_TLCommonBehaviorData) Validate() error {
	return m.ValidateWithPath("CT_TLCommonBehaviorData")
}

// ValidateWithPath validates the CT_TLCommonBehaviorData and its children, prefixing error messages with path
func (m *CT_TLCommonBehaviorData) ValidateWithPath(path string) error {
	if err := m.AdditiveAttr.ValidateWithPath(path + "/AdditiveAttr"); err != nil {
		return err
	}
	if err := m.AccumulateAttr.ValidateWithPath(path + "/AccumulateAttr"); err != nil {
		return err
	}
	if err := m.XfrmTypeAttr.ValidateWithPath(path + "/XfrmTypeAttr"); err != nil {
		return err
	}
	if err := m.OverrideAttr.ValidateWithPath(path + "/OverrideAttr"); err != nil {
		return err
	}
	if err := m.CTn.ValidateWithPath(path + "/CTn"); err != nil {
		return err
	}
	if err := m.TgtEl.ValidateWithPath(path + "/TgtEl"); err != nil {
		return err
	}
	if m.AttrNameLst != nil {
		if err := m.AttrNameLst.ValidateWithPath(path + "/AttrNameLst"); err != nil {
			return err
		}
	}
	return nil
}
