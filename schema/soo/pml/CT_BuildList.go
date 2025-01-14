package pml

import (
	"encoding/xml"
	"fmt"

	"github.com/dhx007/goffice"
)

type CT_BuildList struct {
	// Build Paragraph
	BldP []*CT_TLBuildParagraph
	// Build Diagram
	BldDgm []*CT_TLBuildDiagram
	// Build Embedded Chart
	BldOleChart []*CT_TLOleBuildChart
	// Build Graphics
	BldGraphic []*CT_TLGraphicalObjectBuild
}

func NewCT_BuildList() *CT_BuildList {
	ret := &CT_BuildList{}
	return ret
}

func (m *CT_BuildList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.BldP != nil {
		sebldP := xml.StartElement{Name: xml.Name{Local: "p:bldP"}}
		for _, c := range m.BldP {
			e.EncodeElement(c, sebldP)
		}
	}
	if m.BldDgm != nil {
		sebldDgm := xml.StartElement{Name: xml.Name{Local: "p:bldDgm"}}
		for _, c := range m.BldDgm {
			e.EncodeElement(c, sebldDgm)
		}
	}
	if m.BldOleChart != nil {
		sebldOleChart := xml.StartElement{Name: xml.Name{Local: "p:bldOleChart"}}
		for _, c := range m.BldOleChart {
			e.EncodeElement(c, sebldOleChart)
		}
	}
	if m.BldGraphic != nil {
		sebldGraphic := xml.StartElement{Name: xml.Name{Local: "p:bldGraphic"}}
		for _, c := range m.BldGraphic {
			e.EncodeElement(c, sebldGraphic)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_BuildList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_BuildList:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "bldP"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "bldP"}:
				tmp := NewCT_TLBuildParagraph()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.BldP = append(m.BldP, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "bldDgm"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "bldDgm"}:
				tmp := NewCT_TLBuildDiagram()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.BldDgm = append(m.BldDgm, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "bldOleChart"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "bldOleChart"}:
				tmp := NewCT_TLOleBuildChart()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.BldOleChart = append(m.BldOleChart, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "bldGraphic"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "bldGraphic"}:
				tmp := NewCT_TLGraphicalObjectBuild()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.BldGraphic = append(m.BldGraphic, tmp)
			default:
				goffice.Log("skipping unsupported element on CT_BuildList %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_BuildList
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_BuildList and its children
func (m *CT_BuildList) Validate() error {
	return m.ValidateWithPath("CT_BuildList")
}

// ValidateWithPath validates the CT_BuildList and its children, prefixing error messages with path
func (m *CT_BuildList) ValidateWithPath(path string) error {
	for i, v := range m.BldP {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/BldP[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.BldDgm {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/BldDgm[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.BldOleChart {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/BldOleChart[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.BldGraphic {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/BldGraphic[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
