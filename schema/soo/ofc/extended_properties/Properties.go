package extended_properties

import (
	"encoding/xml"

	"github.com/dhx007/goffice"
)

type Properties struct {
	CT_Properties
}

func NewProperties() *Properties {
	ret := &Properties{}
	ret.CT_Properties = *NewCT_Properties()
	return ret
}

func (m *Properties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/extended-properties"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:vt"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "Properties"
	return m.CT_Properties.MarshalXML(e, start)
}

func (m *Properties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_Properties = *NewCT_Properties()
lProperties:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/extended-properties", Local: "Template"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/extendedProperties", Local: "Template"}:
				m.Template = new(string)
				if err := d.DecodeElement(m.Template, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/extended-properties", Local: "Manager"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/extendedProperties", Local: "Manager"}:
				m.Manager = new(string)
				if err := d.DecodeElement(m.Manager, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/extended-properties", Local: "Company"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/extendedProperties", Local: "Company"}:
				m.Company = new(string)
				if err := d.DecodeElement(m.Company, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/extended-properties", Local: "Pages"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/extendedProperties", Local: "Pages"}:
				m.Pages = new(int32)
				if err := d.DecodeElement(m.Pages, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/extended-properties", Local: "Words"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/extendedProperties", Local: "Words"}:
				m.Words = new(int32)
				if err := d.DecodeElement(m.Words, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/extended-properties", Local: "Characters"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/extendedProperties", Local: "Characters"}:
				m.Characters = new(int32)
				if err := d.DecodeElement(m.Characters, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/extended-properties", Local: "PresentationFormat"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/extendedProperties", Local: "PresentationFormat"}:
				m.PresentationFormat = new(string)
				if err := d.DecodeElement(m.PresentationFormat, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/extended-properties", Local: "Lines"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/extendedProperties", Local: "Lines"}:
				m.Lines = new(int32)
				if err := d.DecodeElement(m.Lines, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/extended-properties", Local: "Paragraphs"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/extendedProperties", Local: "Paragraphs"}:
				m.Paragraphs = new(int32)
				if err := d.DecodeElement(m.Paragraphs, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/extended-properties", Local: "Slides"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/extendedProperties", Local: "Slides"}:
				m.Slides = new(int32)
				if err := d.DecodeElement(m.Slides, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/extended-properties", Local: "Notes"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/extendedProperties", Local: "Notes"}:
				m.Notes = new(int32)
				if err := d.DecodeElement(m.Notes, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/extended-properties", Local: "TotalTime"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/extendedProperties", Local: "TotalTime"}:
				m.TotalTime = new(int32)
				if err := d.DecodeElement(m.TotalTime, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/extended-properties", Local: "HiddenSlides"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/extendedProperties", Local: "HiddenSlides"}:
				m.HiddenSlides = new(int32)
				if err := d.DecodeElement(m.HiddenSlides, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/extended-properties", Local: "MMClips"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/extendedProperties", Local: "MMClips"}:
				m.MMClips = new(int32)
				if err := d.DecodeElement(m.MMClips, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/extended-properties", Local: "ScaleCrop"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/extendedProperties", Local: "ScaleCrop"}:
				m.ScaleCrop = new(bool)
				if err := d.DecodeElement(m.ScaleCrop, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/extended-properties", Local: "HeadingPairs"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/extendedProperties", Local: "HeadingPairs"}:
				m.HeadingPairs = NewCT_VectorVariant()
				if err := d.DecodeElement(m.HeadingPairs, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/extended-properties", Local: "TitlesOfParts"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/extendedProperties", Local: "TitlesOfParts"}:
				m.TitlesOfParts = NewCT_VectorLpstr()
				if err := d.DecodeElement(m.TitlesOfParts, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/extended-properties", Local: "LinksUpToDate"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/extendedProperties", Local: "LinksUpToDate"}:
				m.LinksUpToDate = new(bool)
				if err := d.DecodeElement(m.LinksUpToDate, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/extended-properties", Local: "CharactersWithSpaces"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/extendedProperties", Local: "CharactersWithSpaces"}:
				m.CharactersWithSpaces = new(int32)
				if err := d.DecodeElement(m.CharactersWithSpaces, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/extended-properties", Local: "SharedDoc"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/extendedProperties", Local: "SharedDoc"}:
				m.SharedDoc = new(bool)
				if err := d.DecodeElement(m.SharedDoc, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/extended-properties", Local: "HyperlinkBase"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/extendedProperties", Local: "HyperlinkBase"}:
				m.HyperlinkBase = new(string)
				if err := d.DecodeElement(m.HyperlinkBase, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/extended-properties", Local: "HLinks"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/extendedProperties", Local: "HLinks"}:
				m.HLinks = NewCT_VectorVariant()
				if err := d.DecodeElement(m.HLinks, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/extended-properties", Local: "HyperlinksChanged"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/extendedProperties", Local: "HyperlinksChanged"}:
				m.HyperlinksChanged = new(bool)
				if err := d.DecodeElement(m.HyperlinksChanged, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/extended-properties", Local: "DigSig"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/extendedProperties", Local: "DigSig"}:
				m.DigSig = NewCT_DigSigBlob()
				if err := d.DecodeElement(m.DigSig, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/extended-properties", Local: "Application"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/extendedProperties", Local: "Application"}:
				m.Application = new(string)
				if err := d.DecodeElement(m.Application, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/extended-properties", Local: "AppVersion"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/extendedProperties", Local: "AppVersion"}:
				m.AppVersion = new(string)
				if err := d.DecodeElement(m.AppVersion, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/extended-properties", Local: "DocSecurity"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/extendedProperties", Local: "DocSecurity"}:
				m.DocSecurity = new(int32)
				if err := d.DecodeElement(m.DocSecurity, &el); err != nil {
					return err
				}
			default:
				goffice.Log("skipping unsupported element on Properties %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lProperties
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the Properties and its children
func (m *Properties) Validate() error {
	return m.ValidateWithPath("Properties")
}

// ValidateWithPath validates the Properties and its children, prefixing error messages with path
func (m *Properties) ValidateWithPath(path string) error {
	if err := m.CT_Properties.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
