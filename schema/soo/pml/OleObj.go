package pml

import (
	"encoding/xml"
	"strconv"

	"github.com/dhx007/goffice"
)

type OleObj struct {
	CT_OleObject
}

func NewOleObj() *OleObj {
	ret := &OleObj{}
	ret.CT_OleObject = *NewCT_OleObject()
	return ret
}

func (m *OleObj) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/presentationml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:a"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:p"}, Value: "http://schemas.openxmlformats.org/presentationml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:sh"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "p:oleObj"
	return m.CT_OleObject.MarshalXML(e, start)
}

func (m *OleObj) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_OleObject = *NewCT_OleObject()
	for _, attr := range start.Attr {
		if attr.Name.Space == "http://schemas.openxmlformats.org/officeDocument/2006/relationships" && attr.Name.Local == "id" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IdAttr = &parsed
			continue
		}
		if attr.Name.Local == "spid" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.SpidAttr = &parsed
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
		if attr.Name.Local == "showAsIcon" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowAsIconAttr = &parsed
			continue
		}
		if attr.Name.Local == "progId" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ProgIdAttr = &parsed
			continue
		}
		if attr.Name.Local == "imgW" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := int32(parsed)
			m.ImgWAttr = &pt
			continue
		}
		if attr.Name.Local == "imgH" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := int32(parsed)
			m.ImgHAttr = &pt
			continue
		}
	}
lOleObj:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "embed"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "embed"}:
				m.Choice = NewCT_OleObjectChoice()
				if err := d.DecodeElement(&m.Choice.Embed, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "link"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "link"}:
				m.Choice = NewCT_OleObjectChoice()
				if err := d.DecodeElement(&m.Choice.Link, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "pic"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "pic"}:
				m.Pic = NewCT_Picture()
				if err := d.DecodeElement(m.Pic, &el); err != nil {
					return err
				}
			default:
				goffice.Log("skipping unsupported element on OleObj %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lOleObj
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the OleObj and its children
func (m *OleObj) Validate() error {
	return m.ValidateWithPath("OleObj")
}

// ValidateWithPath validates the OleObj and its children, prefixing error messages with path
func (m *OleObj) ValidateWithPath(path string) error {
	if err := m.CT_OleObject.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
