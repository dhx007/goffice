package pml

import (
	"encoding/xml"

	"github.com/dhx007/goffice"
	"github.com/dhx007/goffice/schema/soo/dml"
)

type CT_TLTimeTargetElement struct {
	// Slide Target
	SldTgt *CT_Empty
	// Sound Target
	SndTgt *dml.CT_EmbeddedWAVAudioFile
	// Shape Target
	SpTgt *CT_TLShapeTargetElement
	// Ink Target
	InkTgt *CT_TLSubShapeId
}

func NewCT_TLTimeTargetElement() *CT_TLTimeTargetElement {
	ret := &CT_TLTimeTargetElement{}
	return ret
}

func (m *CT_TLTimeTargetElement) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.SldTgt != nil {
		sesldTgt := xml.StartElement{Name: xml.Name{Local: "p:sldTgt"}}
		e.EncodeElement(m.SldTgt, sesldTgt)
	}
	if m.SndTgt != nil {
		sesndTgt := xml.StartElement{Name: xml.Name{Local: "p:sndTgt"}}
		e.EncodeElement(m.SndTgt, sesndTgt)
	}
	if m.SpTgt != nil {
		sespTgt := xml.StartElement{Name: xml.Name{Local: "p:spTgt"}}
		e.EncodeElement(m.SpTgt, sespTgt)
	}
	if m.InkTgt != nil {
		seinkTgt := xml.StartElement{Name: xml.Name{Local: "p:inkTgt"}}
		e.EncodeElement(m.InkTgt, seinkTgt)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TLTimeTargetElement) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_TLTimeTargetElement:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "sldTgt"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "sldTgt"}:
				m.SldTgt = NewCT_Empty()
				if err := d.DecodeElement(m.SldTgt, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "sndTgt"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "sndTgt"}:
				m.SndTgt = dml.NewCT_EmbeddedWAVAudioFile()
				if err := d.DecodeElement(m.SndTgt, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "spTgt"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "spTgt"}:
				m.SpTgt = NewCT_TLShapeTargetElement()
				if err := d.DecodeElement(m.SpTgt, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "inkTgt"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "inkTgt"}:
				m.InkTgt = NewCT_TLSubShapeId()
				if err := d.DecodeElement(m.InkTgt, &el); err != nil {
					return err
				}
			default:
				goffice.Log("skipping unsupported element on CT_TLTimeTargetElement %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TLTimeTargetElement
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TLTimeTargetElement and its children
func (m *CT_TLTimeTargetElement) Validate() error {
	return m.ValidateWithPath("CT_TLTimeTargetElement")
}

// ValidateWithPath validates the CT_TLTimeTargetElement and its children, prefixing error messages with path
func (m *CT_TLTimeTargetElement) ValidateWithPath(path string) error {
	if m.SldTgt != nil {
		if err := m.SldTgt.ValidateWithPath(path + "/SldTgt"); err != nil {
			return err
		}
	}
	if m.SndTgt != nil {
		if err := m.SndTgt.ValidateWithPath(path + "/SndTgt"); err != nil {
			return err
		}
	}
	if m.SpTgt != nil {
		if err := m.SpTgt.ValidateWithPath(path + "/SpTgt"); err != nil {
			return err
		}
	}
	if m.InkTgt != nil {
		if err := m.InkTgt.ValidateWithPath(path + "/InkTgt"); err != nil {
			return err
		}
	}
	return nil
}
