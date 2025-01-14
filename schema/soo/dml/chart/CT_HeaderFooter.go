package chart

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/dhx007/goffice"
)

type CT_HeaderFooter struct {
	AlignWithMarginsAttr *bool
	DifferentOddEvenAttr *bool
	DifferentFirstAttr   *bool
	OddHeader            *string
	OddFooter            *string
	EvenHeader           *string
	EvenFooter           *string
	FirstHeader          *string
	FirstFooter          *string
}

func NewCT_HeaderFooter() *CT_HeaderFooter {
	ret := &CT_HeaderFooter{}
	return ret
}

func (m *CT_HeaderFooter) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.AlignWithMarginsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "alignWithMargins"},
			Value: fmt.Sprintf("%d", b2i(*m.AlignWithMarginsAttr))})
	}
	if m.DifferentOddEvenAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "differentOddEven"},
			Value: fmt.Sprintf("%d", b2i(*m.DifferentOddEvenAttr))})
	}
	if m.DifferentFirstAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "differentFirst"},
			Value: fmt.Sprintf("%d", b2i(*m.DifferentFirstAttr))})
	}
	e.EncodeToken(start)
	if m.OddHeader != nil {
		seoddHeader := xml.StartElement{Name: xml.Name{Local: "c:oddHeader"}}
		goffice.AddPreserveSpaceAttr(&seoddHeader, *m.OddHeader)
		e.EncodeElement(m.OddHeader, seoddHeader)
	}
	if m.OddFooter != nil {
		seoddFooter := xml.StartElement{Name: xml.Name{Local: "c:oddFooter"}}
		goffice.AddPreserveSpaceAttr(&seoddFooter, *m.OddFooter)
		e.EncodeElement(m.OddFooter, seoddFooter)
	}
	if m.EvenHeader != nil {
		seevenHeader := xml.StartElement{Name: xml.Name{Local: "c:evenHeader"}}
		goffice.AddPreserveSpaceAttr(&seevenHeader, *m.EvenHeader)
		e.EncodeElement(m.EvenHeader, seevenHeader)
	}
	if m.EvenFooter != nil {
		seevenFooter := xml.StartElement{Name: xml.Name{Local: "c:evenFooter"}}
		goffice.AddPreserveSpaceAttr(&seevenFooter, *m.EvenFooter)
		e.EncodeElement(m.EvenFooter, seevenFooter)
	}
	if m.FirstHeader != nil {
		sefirstHeader := xml.StartElement{Name: xml.Name{Local: "c:firstHeader"}}
		goffice.AddPreserveSpaceAttr(&sefirstHeader, *m.FirstHeader)
		e.EncodeElement(m.FirstHeader, sefirstHeader)
	}
	if m.FirstFooter != nil {
		sefirstFooter := xml.StartElement{Name: xml.Name{Local: "c:firstFooter"}}
		goffice.AddPreserveSpaceAttr(&sefirstFooter, *m.FirstFooter)
		e.EncodeElement(m.FirstFooter, sefirstFooter)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_HeaderFooter) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "alignWithMargins" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.AlignWithMarginsAttr = &parsed
			continue
		}
		if attr.Name.Local == "differentOddEven" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.DifferentOddEvenAttr = &parsed
			continue
		}
		if attr.Name.Local == "differentFirst" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.DifferentFirstAttr = &parsed
			continue
		}
	}
lCT_HeaderFooter:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "oddHeader"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "oddHeader"}:
				m.OddHeader = new(string)
				if err := d.DecodeElement(m.OddHeader, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "oddFooter"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "oddFooter"}:
				m.OddFooter = new(string)
				if err := d.DecodeElement(m.OddFooter, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "evenHeader"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "evenHeader"}:
				m.EvenHeader = new(string)
				if err := d.DecodeElement(m.EvenHeader, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "evenFooter"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "evenFooter"}:
				m.EvenFooter = new(string)
				if err := d.DecodeElement(m.EvenFooter, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "firstHeader"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "firstHeader"}:
				m.FirstHeader = new(string)
				if err := d.DecodeElement(m.FirstHeader, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "firstFooter"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "firstFooter"}:
				m.FirstFooter = new(string)
				if err := d.DecodeElement(m.FirstFooter, &el); err != nil {
					return err
				}
			default:
				goffice.Log("skipping unsupported element on CT_HeaderFooter %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_HeaderFooter
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_HeaderFooter and its children
func (m *CT_HeaderFooter) Validate() error {
	return m.ValidateWithPath("CT_HeaderFooter")
}

// ValidateWithPath validates the CT_HeaderFooter and its children, prefixing error messages with path
func (m *CT_HeaderFooter) ValidateWithPath(path string) error {
	return nil
}
