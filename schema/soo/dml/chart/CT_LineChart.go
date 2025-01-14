package chart

import (
	"encoding/xml"
	"fmt"

	"github.com/dhx007/goffice"
)

type CT_LineChart struct {
	Grouping   *CT_Grouping
	VaryColors *CT_Boolean
	Ser        []*CT_LineSer
	DLbls      *CT_DLbls
	DropLines  *CT_ChartLines
	HiLowLines *CT_ChartLines
	UpDownBars *CT_UpDownBars
	Marker     *CT_Boolean
	Smooth     *CT_Boolean
	AxId       []*CT_UnsignedInt
	ExtLst     *CT_ExtensionList
}

func NewCT_LineChart() *CT_LineChart {
	ret := &CT_LineChart{}
	ret.Grouping = NewCT_Grouping()
	return ret
}

func (m *CT_LineChart) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	segrouping := xml.StartElement{Name: xml.Name{Local: "c:grouping"}}
	e.EncodeElement(m.Grouping, segrouping)
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
	if m.DropLines != nil {
		sedropLines := xml.StartElement{Name: xml.Name{Local: "c:dropLines"}}
		e.EncodeElement(m.DropLines, sedropLines)
	}
	if m.HiLowLines != nil {
		sehiLowLines := xml.StartElement{Name: xml.Name{Local: "c:hiLowLines"}}
		e.EncodeElement(m.HiLowLines, sehiLowLines)
	}
	if m.UpDownBars != nil {
		seupDownBars := xml.StartElement{Name: xml.Name{Local: "c:upDownBars"}}
		e.EncodeElement(m.UpDownBars, seupDownBars)
	}
	if m.Marker != nil {
		semarker := xml.StartElement{Name: xml.Name{Local: "c:marker"}}
		e.EncodeElement(m.Marker, semarker)
	}
	if m.Smooth != nil {
		sesmooth := xml.StartElement{Name: xml.Name{Local: "c:smooth"}}
		e.EncodeElement(m.Smooth, sesmooth)
	}
	seaxId := xml.StartElement{Name: xml.Name{Local: "c:axId"}}
	for _, c := range m.AxId {
		e.EncodeElement(c, seaxId)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "c:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_LineChart) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Grouping = NewCT_Grouping()
lCT_LineChart:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "grouping"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "grouping"}:
				if err := d.DecodeElement(m.Grouping, &el); err != nil {
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
				tmp := NewCT_LineSer()
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
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "dropLines"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "dropLines"}:
				m.DropLines = NewCT_ChartLines()
				if err := d.DecodeElement(m.DropLines, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "hiLowLines"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "hiLowLines"}:
				m.HiLowLines = NewCT_ChartLines()
				if err := d.DecodeElement(m.HiLowLines, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "upDownBars"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "upDownBars"}:
				m.UpDownBars = NewCT_UpDownBars()
				if err := d.DecodeElement(m.UpDownBars, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "marker"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "marker"}:
				m.Marker = NewCT_Boolean()
				if err := d.DecodeElement(m.Marker, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "smooth"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "smooth"}:
				m.Smooth = NewCT_Boolean()
				if err := d.DecodeElement(m.Smooth, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "axId"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "axId"}:
				tmp := NewCT_UnsignedInt()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.AxId = append(m.AxId, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				goffice.Log("skipping unsupported element on CT_LineChart %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_LineChart
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_LineChart and its children
func (m *CT_LineChart) Validate() error {
	return m.ValidateWithPath("CT_LineChart")
}

// ValidateWithPath validates the CT_LineChart and its children, prefixing error messages with path
func (m *CT_LineChart) ValidateWithPath(path string) error {
	if err := m.Grouping.ValidateWithPath(path + "/Grouping"); err != nil {
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
	if m.DropLines != nil {
		if err := m.DropLines.ValidateWithPath(path + "/DropLines"); err != nil {
			return err
		}
	}
	if m.HiLowLines != nil {
		if err := m.HiLowLines.ValidateWithPath(path + "/HiLowLines"); err != nil {
			return err
		}
	}
	if m.UpDownBars != nil {
		if err := m.UpDownBars.ValidateWithPath(path + "/UpDownBars"); err != nil {
			return err
		}
	}
	if m.Marker != nil {
		if err := m.Marker.ValidateWithPath(path + "/Marker"); err != nil {
			return err
		}
	}
	if m.Smooth != nil {
		if err := m.Smooth.ValidateWithPath(path + "/Smooth"); err != nil {
			return err
		}
	}
	for i, v := range m.AxId {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/AxId[%d]", path, i)); err != nil {
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
