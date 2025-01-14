package wml

import (
	"encoding/xml"

	"github.com/dhx007/goffice"
	"github.com/dhx007/goffice/schema/soo/dml"
)

type WdCT_WordprocessingShapeChoice struct {
	CNvSpPr *dml.CT_NonVisualDrawingShapeProps
	CNvCnPr *dml.CT_NonVisualConnectorProperties
}

func NewWdCT_WordprocessingShapeChoice() *WdCT_WordprocessingShapeChoice {
	ret := &WdCT_WordprocessingShapeChoice{}
	return ret
}

func (m *WdCT_WordprocessingShapeChoice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.CNvSpPr != nil {
		secNvSpPr := xml.StartElement{Name: xml.Name{Local: "wp:cNvSpPr"}}
		e.EncodeElement(m.CNvSpPr, secNvSpPr)
	}
	if m.CNvCnPr != nil {
		secNvCnPr := xml.StartElement{Name: xml.Name{Local: "wp:cNvCnPr"}}
		e.EncodeElement(m.CNvCnPr, secNvCnPr)
	}
	return nil
}

func (m *WdCT_WordprocessingShapeChoice) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lWdCT_WordprocessingShapeChoice:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "cNvSpPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/wordprocessingDrawing", Local: "cNvSpPr"}:
				m.CNvSpPr = dml.NewCT_NonVisualDrawingShapeProps()
				if err := d.DecodeElement(m.CNvSpPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "cNvCnPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/wordprocessingDrawing", Local: "cNvCnPr"}:
				m.CNvCnPr = dml.NewCT_NonVisualConnectorProperties()
				if err := d.DecodeElement(m.CNvCnPr, &el); err != nil {
					return err
				}
			default:
				goffice.Log("skipping unsupported element on WdCT_WordprocessingShapeChoice %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lWdCT_WordprocessingShapeChoice
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the WdCT_WordprocessingShapeChoice and its children
func (m *WdCT_WordprocessingShapeChoice) Validate() error {
	return m.ValidateWithPath("WdCT_WordprocessingShapeChoice")
}

// ValidateWithPath validates the WdCT_WordprocessingShapeChoice and its children, prefixing error messages with path
func (m *WdCT_WordprocessingShapeChoice) ValidateWithPath(path string) error {
	if m.CNvSpPr != nil {
		if err := m.CNvSpPr.ValidateWithPath(path + "/CNvSpPr"); err != nil {
			return err
		}
	}
	if m.CNvCnPr != nil {
		if err := m.CNvCnPr.ValidateWithPath(path + "/CNvCnPr"); err != nil {
			return err
		}
	}
	return nil
}
