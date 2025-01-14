package pml

import (
	"encoding/xml"

	"github.com/dhx007/goffice"
	"github.com/dhx007/goffice/schema/soo/dml"
)

type CT_PresentationProperties struct {
	// HTML Publishing Properties
	HtmlPubPr *CT_HtmlPublishProperties
	// Web Properties
	WebPr *CT_WebProperties
	// Printing Properties
	PrnPr *CT_PrintProperties
	// Presentation-wide Show Properties
	ShowPr *CT_ShowProperties
	// Color MRU
	ClrMru *dml.CT_ColorMRU
	ExtLst *CT_ExtensionList
}

func NewCT_PresentationProperties() *CT_PresentationProperties {
	ret := &CT_PresentationProperties{}
	return ret
}

func (m *CT_PresentationProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.HtmlPubPr != nil {
		sehtmlPubPr := xml.StartElement{Name: xml.Name{Local: "p:htmlPubPr"}}
		e.EncodeElement(m.HtmlPubPr, sehtmlPubPr)
	}
	if m.WebPr != nil {
		sewebPr := xml.StartElement{Name: xml.Name{Local: "p:webPr"}}
		e.EncodeElement(m.WebPr, sewebPr)
	}
	if m.PrnPr != nil {
		seprnPr := xml.StartElement{Name: xml.Name{Local: "p:prnPr"}}
		e.EncodeElement(m.PrnPr, seprnPr)
	}
	if m.ShowPr != nil {
		seshowPr := xml.StartElement{Name: xml.Name{Local: "p:showPr"}}
		e.EncodeElement(m.ShowPr, seshowPr)
	}
	if m.ClrMru != nil {
		seclrMru := xml.StartElement{Name: xml.Name{Local: "p:clrMru"}}
		e.EncodeElement(m.ClrMru, seclrMru)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "p:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_PresentationProperties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_PresentationProperties:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "htmlPubPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "htmlPubPr"}:
				m.HtmlPubPr = NewCT_HtmlPublishProperties()
				if err := d.DecodeElement(m.HtmlPubPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "webPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "webPr"}:
				m.WebPr = NewCT_WebProperties()
				if err := d.DecodeElement(m.WebPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "prnPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "prnPr"}:
				m.PrnPr = NewCT_PrintProperties()
				if err := d.DecodeElement(m.PrnPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "showPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "showPr"}:
				m.ShowPr = NewCT_ShowProperties()
				if err := d.DecodeElement(m.ShowPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "clrMru"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "clrMru"}:
				m.ClrMru = dml.NewCT_ColorMRU()
				if err := d.DecodeElement(m.ClrMru, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				goffice.Log("skipping unsupported element on CT_PresentationProperties %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_PresentationProperties
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_PresentationProperties and its children
func (m *CT_PresentationProperties) Validate() error {
	return m.ValidateWithPath("CT_PresentationProperties")
}

// ValidateWithPath validates the CT_PresentationProperties and its children, prefixing error messages with path
func (m *CT_PresentationProperties) ValidateWithPath(path string) error {
	if m.HtmlPubPr != nil {
		if err := m.HtmlPubPr.ValidateWithPath(path + "/HtmlPubPr"); err != nil {
			return err
		}
	}
	if m.WebPr != nil {
		if err := m.WebPr.ValidateWithPath(path + "/WebPr"); err != nil {
			return err
		}
	}
	if m.PrnPr != nil {
		if err := m.PrnPr.ValidateWithPath(path + "/PrnPr"); err != nil {
			return err
		}
	}
	if m.ShowPr != nil {
		if err := m.ShowPr.ValidateWithPath(path + "/ShowPr"); err != nil {
			return err
		}
	}
	if m.ClrMru != nil {
		if err := m.ClrMru.ValidateWithPath(path + "/ClrMru"); err != nil {
			return err
		}
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
