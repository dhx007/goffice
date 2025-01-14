package dml

import (
	"encoding/xml"
	"fmt"

	"github.com/dhx007/goffice"
)

type CT_AdjustHandleList struct {
	AhXY    []*CT_XYAdjustHandle
	AhPolar []*CT_PolarAdjustHandle
}

func NewCT_AdjustHandleList() *CT_AdjustHandleList {
	ret := &CT_AdjustHandleList{}
	return ret
}

func (m *CT_AdjustHandleList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.AhXY != nil {
		seahXY := xml.StartElement{Name: xml.Name{Local: "a:ahXY"}}
		for _, c := range m.AhXY {
			e.EncodeElement(c, seahXY)
		}
	}
	if m.AhPolar != nil {
		seahPolar := xml.StartElement{Name: xml.Name{Local: "a:ahPolar"}}
		for _, c := range m.AhPolar {
			e.EncodeElement(c, seahPolar)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_AdjustHandleList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_AdjustHandleList:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "ahXY"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "ahXY"}:
				tmp := NewCT_XYAdjustHandle()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.AhXY = append(m.AhXY, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "ahPolar"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "ahPolar"}:
				tmp := NewCT_PolarAdjustHandle()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.AhPolar = append(m.AhPolar, tmp)
			default:
				goffice.Log("skipping unsupported element on CT_AdjustHandleList %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_AdjustHandleList
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_AdjustHandleList and its children
func (m *CT_AdjustHandleList) Validate() error {
	return m.ValidateWithPath("CT_AdjustHandleList")
}

// ValidateWithPath validates the CT_AdjustHandleList and its children, prefixing error messages with path
func (m *CT_AdjustHandleList) ValidateWithPath(path string) error {
	for i, v := range m.AhXY {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/AhXY[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.AhPolar {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/AhPolar[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
