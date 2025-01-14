package elements

import (
	"encoding/xml"
	"fmt"

	"github.com/dhx007/goffice"
)

type ElementContainer struct {
	Choice []*ElementsGroupChoice
}

func NewElementContainer() *ElementContainer {
	ret := &ElementContainer{}
	return ret
}

func (m *ElementContainer) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "elementContainer"
	e.EncodeToken(start)
	if m.Choice != nil {
		for _, c := range m.Choice {
			c.MarshalXML(e, xml.StartElement{})
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *ElementContainer) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lElementContainer:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://purl.org/dc/elements/1.1/", Local: "any"}:
				tmp := NewElementsGroupChoice()
				if err := d.DecodeElement(&tmp.Any, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			default:
				goffice.Log("skipping unsupported element on ElementContainer %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lElementContainer
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the ElementContainer and its children
func (m *ElementContainer) Validate() error {
	return m.ValidateWithPath("ElementContainer")
}

// ValidateWithPath validates the ElementContainer and its children, prefixing error messages with path
func (m *ElementContainer) ValidateWithPath(path string) error {
	for i, v := range m.Choice {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Choice[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
