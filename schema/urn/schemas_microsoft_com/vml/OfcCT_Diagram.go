package vml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/dhx007/goffice"
	"github.com/dhx007/goffice/schema/soo/ofc/sharedTypes"
)

type OfcCT_Diagram struct {
	DgmstyleAttr         *int64
	AutoformatAttr       sharedTypes.ST_TrueFalse
	ReverseAttr          sharedTypes.ST_TrueFalse
	AutolayoutAttr       sharedTypes.ST_TrueFalse
	DgmscalexAttr        *int64
	DgmscaleyAttr        *int64
	DgmfontsizeAttr      *int64
	ConstrainboundsAttr  *string
	DgmbasetextscaleAttr *int64
	Relationtable        *OfcCT_RelationTable
	ExtAttr              ST_Ext
}

func NewOfcCT_Diagram() *OfcCT_Diagram {
	ret := &OfcCT_Diagram{}
	return ret
}

func (m *OfcCT_Diagram) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.DgmstyleAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "dgmstyle"},
			Value: fmt.Sprintf("%v", *m.DgmstyleAttr)})
	}
	if m.AutoformatAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.AutoformatAttr.MarshalXMLAttr(xml.Name{Local: "autoformat"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.ReverseAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.ReverseAttr.MarshalXMLAttr(xml.Name{Local: "reverse"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.AutolayoutAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.AutolayoutAttr.MarshalXMLAttr(xml.Name{Local: "autolayout"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.DgmscalexAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "dgmscalex"},
			Value: fmt.Sprintf("%v", *m.DgmscalexAttr)})
	}
	if m.DgmscaleyAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "dgmscaley"},
			Value: fmt.Sprintf("%v", *m.DgmscaleyAttr)})
	}
	if m.DgmfontsizeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "dgmfontsize"},
			Value: fmt.Sprintf("%v", *m.DgmfontsizeAttr)})
	}
	if m.ConstrainboundsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "constrainbounds"},
			Value: fmt.Sprintf("%v", *m.ConstrainboundsAttr)})
	}
	if m.DgmbasetextscaleAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "dgmbasetextscale"},
			Value: fmt.Sprintf("%v", *m.DgmbasetextscaleAttr)})
	}
	if m.ExtAttr != ST_ExtUnset {
		attr, err := m.ExtAttr.MarshalXMLAttr(xml.Name{Local: "ext"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	if m.Relationtable != nil {
		serelationtable := xml.StartElement{Name: xml.Name{Local: "o:relationtable"}}
		e.EncodeElement(m.Relationtable, serelationtable)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *OfcCT_Diagram) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "dgmstyle" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.DgmstyleAttr = &parsed
			continue
		}
		if attr.Name.Local == "autoformat" {
			m.AutoformatAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "reverse" {
			m.ReverseAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "autolayout" {
			m.AutolayoutAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "dgmscalex" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.DgmscalexAttr = &parsed
			continue
		}
		if attr.Name.Local == "dgmscaley" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.DgmscaleyAttr = &parsed
			continue
		}
		if attr.Name.Local == "dgmfontsize" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.DgmfontsizeAttr = &parsed
			continue
		}
		if attr.Name.Local == "constrainbounds" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ConstrainboundsAttr = &parsed
			continue
		}
		if attr.Name.Local == "dgmbasetextscale" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.DgmbasetextscaleAttr = &parsed
			continue
		}
		if attr.Name.Local == "ext" {
			m.ExtAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
lOfcCT_Diagram:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "urn:schemas-microsoft-com:office:office", Local: "relationtable"}:
				m.Relationtable = NewOfcCT_RelationTable()
				if err := d.DecodeElement(m.Relationtable, &el); err != nil {
					return err
				}
			default:
				goffice.Log("skipping unsupported element on OfcCT_Diagram %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lOfcCT_Diagram
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the OfcCT_Diagram and its children
func (m *OfcCT_Diagram) Validate() error {
	return m.ValidateWithPath("OfcCT_Diagram")
}

// ValidateWithPath validates the OfcCT_Diagram and its children, prefixing error messages with path
func (m *OfcCT_Diagram) ValidateWithPath(path string) error {
	if err := m.AutoformatAttr.ValidateWithPath(path + "/AutoformatAttr"); err != nil {
		return err
	}
	if err := m.ReverseAttr.ValidateWithPath(path + "/ReverseAttr"); err != nil {
		return err
	}
	if err := m.AutolayoutAttr.ValidateWithPath(path + "/AutolayoutAttr"); err != nil {
		return err
	}
	if m.Relationtable != nil {
		if err := m.Relationtable.ValidateWithPath(path + "/Relationtable"); err != nil {
			return err
		}
	}
	if err := m.ExtAttr.ValidateWithPath(path + "/ExtAttr"); err != nil {
		return err
	}
	return nil
}
