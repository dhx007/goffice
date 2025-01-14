package core_properties

import (
	"encoding/xml"
	"time"

	"github.com/dhx007/goffice"
)

type CoreProperties struct {
	CT_CoreProperties
}

func NewCoreProperties() *CoreProperties {
	ret := &CoreProperties{}
	ret.CT_CoreProperties = *NewCT_CoreProperties()
	return ret
}

func (m *CoreProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/package/2006/metadata/core-properties"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:cp"}, Value: "http://schemas.openxmlformats.org/package/2006/metadata/core-properties"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:dc"}, Value: "http://purl.org/dc/elements/1.1/"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:dcterms"}, Value: "http://purl.org/dc/terms/"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "cp:coreProperties"
	return m.CT_CoreProperties.MarshalXML(e, start)
}

func (m *CoreProperties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_CoreProperties = *NewCT_CoreProperties()
lCoreProperties:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/package/2006/metadata/core-properties", Local: "category"}:
				m.Category = new(string)
				if err := d.DecodeElement(m.Category, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/package/2006/metadata/core-properties", Local: "contentStatus"}:
				m.ContentStatus = new(string)
				if err := d.DecodeElement(m.ContentStatus, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://purl.org/dc/terms/", Local: "created"}:
				m.Created = new(goffice.XSDAny)
				if err := d.DecodeElement(m.Created, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://purl.org/dc/elements/1.1/", Local: "creator"}:
				m.Creator = new(goffice.XSDAny)
				if err := d.DecodeElement(m.Creator, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://purl.org/dc/elements/1.1/", Local: "description"}:
				m.Description = new(goffice.XSDAny)
				if err := d.DecodeElement(m.Description, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://purl.org/dc/elements/1.1/", Local: "identifier"}:
				m.Identifier = new(goffice.XSDAny)
				if err := d.DecodeElement(m.Identifier, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/package/2006/metadata/core-properties", Local: "keywords"}:
				m.Keywords = NewCT_Keywords()
				if err := d.DecodeElement(m.Keywords, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://purl.org/dc/elements/1.1/", Local: "language"}:
				m.Language = new(goffice.XSDAny)
				if err := d.DecodeElement(m.Language, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/package/2006/metadata/core-properties", Local: "lastModifiedBy"}:
				m.LastModifiedBy = new(string)
				if err := d.DecodeElement(m.LastModifiedBy, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/package/2006/metadata/core-properties", Local: "lastPrinted"}:
				m.LastPrinted = new(time.Time)
				if err := d.DecodeElement(m.LastPrinted, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://purl.org/dc/terms/", Local: "modified"}:
				m.Modified = new(goffice.XSDAny)
				if err := d.DecodeElement(m.Modified, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/package/2006/metadata/core-properties", Local: "revision"}:
				m.Revision = new(string)
				if err := d.DecodeElement(m.Revision, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://purl.org/dc/elements/1.1/", Local: "subject"}:
				m.Subject = new(goffice.XSDAny)
				if err := d.DecodeElement(m.Subject, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://purl.org/dc/elements/1.1/", Local: "title"}:
				m.Title = new(goffice.XSDAny)
				if err := d.DecodeElement(m.Title, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/package/2006/metadata/core-properties", Local: "version"}:
				m.Version = new(string)
				if err := d.DecodeElement(m.Version, &el); err != nil {
					return err
				}
			default:
				goffice.Log("skipping unsupported element on CoreProperties %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCoreProperties
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CoreProperties and its children
func (m *CoreProperties) Validate() error {
	return m.ValidateWithPath("CoreProperties")
}

// ValidateWithPath validates the CoreProperties and its children, prefixing error messages with path
func (m *CoreProperties) ValidateWithPath(path string) error {
	if err := m.CT_CoreProperties.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
