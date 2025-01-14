package chart

import (
	"encoding/xml"
	"fmt"

	"github.com/dhx007/goffice"
)

type CT_OfPieChart struct {
	OfPieType     *CT_OfPieType
	VaryColors    *CT_Boolean
	Ser           []*CT_PieSer
	DLbls         *CT_DLbls
	GapWidth      *CT_GapAmount
	SplitType     *CT_SplitType
	SplitPos      *CT_Double
	CustSplit     *CT_CustSplit
	SecondPieSize *CT_SecondPieSize
	SerLines      []*CT_ChartLines
	ExtLst        *CT_ExtensionList
}

func NewCT_OfPieChart() *CT_OfPieChart {
	ret := &CT_OfPieChart{}
	ret.OfPieType = NewCT_OfPieType()
	return ret
}

func (m *CT_OfPieChart) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	seofPieType := xml.StartElement{Name: xml.Name{Local: "c:ofPieType"}}
	e.EncodeElement(m.OfPieType, seofPieType)
	if m.VaryColors != nil {
		sevaryColors := xml.StartElement{Name: xml.Name{Local: "c:varyColors"}}
		e.EncodeElement(m.VaryColors, sevaryColors)
	}
	if m.Ser != nil {
		seser := xml.StartElement{Name: xml.Name{Local: "c:ser"}}
		for _, c := range m.Ser {
			e.EncodeElement(c, seser)
		}
	}
	if m.DLbls != nil {
		sedLbls := xml.StartElement{Name: xml.Name{Local: "c:dLbls"}}
		e.EncodeElement(m.DLbls, sedLbls)
	}
	if m.GapWidth != nil {
		segapWidth := xml.StartElement{Name: xml.Name{Local: "c:gapWidth"}}
		e.EncodeElement(m.GapWidth, segapWidth)
	}
	if m.SplitType != nil {
		sesplitType := xml.StartElement{Name: xml.Name{Local: "c:splitType"}}
		e.EncodeElement(m.SplitType, sesplitType)
	}
	if m.SplitPos != nil {
		sesplitPos := xml.StartElement{Name: xml.Name{Local: "c:splitPos"}}
		e.EncodeElement(m.SplitPos, sesplitPos)
	}
	if m.CustSplit != nil {
		secustSplit := xml.StartElement{Name: xml.Name{Local: "c:custSplit"}}
		e.EncodeElement(m.CustSplit, secustSplit)
	}
	if m.SecondPieSize != nil {
		sesecondPieSize := xml.StartElement{Name: xml.Name{Local: "c:secondPieSize"}}
		e.EncodeElement(m.SecondPieSize, sesecondPieSize)
	}
	if m.SerLines != nil {
		seserLines := xml.StartElement{Name: xml.Name{Local: "c:serLines"}}
		for _, c := range m.SerLines {
			e.EncodeElement(c, seserLines)
		}
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "c:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_OfPieChart) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.OfPieType = NewCT_OfPieType()
lCT_OfPieChart:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "ofPieType"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "ofPieType"}:
				if err := d.DecodeElement(m.OfPieType, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "varyColors"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "varyColors"}:
				m.VaryColors = NewCT_Boolean()
				if err := d.DecodeElement(m.VaryColors, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "ser"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "ser"}:
				tmp := NewCT_PieSer()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Ser = append(m.Ser, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "dLbls"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "dLbls"}:
				m.DLbls = NewCT_DLbls()
				if err := d.DecodeElement(m.DLbls, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "gapWidth"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "gapWidth"}:
				m.GapWidth = NewCT_GapAmount()
				if err := d.DecodeElement(m.GapWidth, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "splitType"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "splitType"}:
				m.SplitType = NewCT_SplitType()
				if err := d.DecodeElement(m.SplitType, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "splitPos"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "splitPos"}:
				m.SplitPos = NewCT_Double()
				if err := d.DecodeElement(m.SplitPos, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "custSplit"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "custSplit"}:
				m.CustSplit = NewCT_CustSplit()
				if err := d.DecodeElement(m.CustSplit, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "secondPieSize"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "secondPieSize"}:
				m.SecondPieSize = NewCT_SecondPieSize()
				if err := d.DecodeElement(m.SecondPieSize, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "serLines"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "serLines"}:
				tmp := NewCT_ChartLines()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.SerLines = append(m.SerLines, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				goffice.Log("skipping unsupported element on CT_OfPieChart %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_OfPieChart
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_OfPieChart and its children
func (m *CT_OfPieChart) Validate() error {
	return m.ValidateWithPath("CT_OfPieChart")
}

// ValidateWithPath validates the CT_OfPieChart and its children, prefixing error messages with path
func (m *CT_OfPieChart) ValidateWithPath(path string) error {
	if err := m.OfPieType.ValidateWithPath(path + "/OfPieType"); err != nil {
		return err
	}
	if m.VaryColors != nil {
		if err := m.VaryColors.ValidateWithPath(path + "/VaryColors"); err != nil {
			return err
		}
	}
	for i, v := range m.Ser {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Ser[%d]", path, i)); err != nil {
			return err
		}
	}
	if m.DLbls != nil {
		if err := m.DLbls.ValidateWithPath(path + "/DLbls"); err != nil {
			return err
		}
	}
	if m.GapWidth != nil {
		if err := m.GapWidth.ValidateWithPath(path + "/GapWidth"); err != nil {
			return err
		}
	}
	if m.SplitType != nil {
		if err := m.SplitType.ValidateWithPath(path + "/SplitType"); err != nil {
			return err
		}
	}
	if m.SplitPos != nil {
		if err := m.SplitPos.ValidateWithPath(path + "/SplitPos"); err != nil {
			return err
		}
	}
	if m.CustSplit != nil {
		if err := m.CustSplit.ValidateWithPath(path + "/CustSplit"); err != nil {
			return err
		}
	}
	if m.SecondPieSize != nil {
		if err := m.SecondPieSize.ValidateWithPath(path + "/SecondPieSize"); err != nil {
			return err
		}
	}
	for i, v := range m.SerLines {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/SerLines[%d]", path, i)); err != nil {
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
