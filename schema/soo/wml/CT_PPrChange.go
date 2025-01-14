package wml

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"time"

	"github.com/dhx007/goffice"
)

type CT_PPrChange struct {
	AuthorAttr string
	DateAttr   *time.Time
	// Annotation Identifier
	IdAttr int64
	PPr    *CT_PPrBase
}

func NewCT_PPrChange() *CT_PPrChange {
	ret := &CT_PPrChange{}
	ret.PPr = NewCT_PPrBase()
	return ret
}

func (m *CT_PPrChange) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:author"},
		Value: fmt.Sprintf("%v", m.AuthorAttr)})
	if m.DateAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:date"},
			Value: fmt.Sprintf("%v", *m.DateAttr)})
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:id"},
		Value: fmt.Sprintf("%v", m.IdAttr)})
	e.EncodeToken(start)
	sepPr := xml.StartElement{Name: xml.Name{Local: "w:pPr"}}
	e.EncodeElement(m.PPr, sepPr)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_PPrChange) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.PPr = NewCT_PPrBase()
	for _, attr := range start.Attr {
		if attr.Name.Local == "author" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.AuthorAttr = parsed
			continue
		}
		if attr.Name.Local == "date" {
			parsed, err := ParseStdlibTime(attr.Value)
			if err != nil {
				return err
			}
			m.DateAttr = &parsed
			continue
		}
		if attr.Name.Local == "id" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.IdAttr = parsed
			continue
		}
	}
lCT_PPrChange:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "pPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "pPr"}:
				if err := d.DecodeElement(m.PPr, &el); err != nil {
					return err
				}
			default:
				goffice.Log("skipping unsupported element on CT_PPrChange %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_PPrChange
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_PPrChange and its children
func (m *CT_PPrChange) Validate() error {
	return m.ValidateWithPath("CT_PPrChange")
}

// ValidateWithPath validates the CT_PPrChange and its children, prefixing error messages with path
func (m *CT_PPrChange) ValidateWithPath(path string) error {
	if err := m.PPr.ValidateWithPath(path + "/PPr"); err != nil {
		return err
	}
	return nil
}
