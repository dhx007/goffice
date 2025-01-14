package chart

import (
	"encoding/xml"
	"fmt"

	"github.com/dhx007/goffice"
)

type CT_BubbleChart struct {
	VaryColors     *CT_Boolean
	Ser            []*CT_BubbleSer
	DLbls          *CT_DLbls
	Bubble3D       *CT_Boolean
	BubbleScale    *CT_BubbleScale
	ShowNegBubbles *CT_Boolean
	SizeRepresents *CT_SizeRepresents
	AxId           []*CT_UnsignedInt
	ExtLst         *CT_ExtensionList
}

func NewCT_BubbleChart() *CT_BubbleChart {
	ret := &CT_BubbleChart{}
	return ret
}

func (m *CT_BubbleChart) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
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
	if m.Bubble3D != nil {
		sebubble3D := xml.StartElement{Name: xml.Name{Local: "c:bubble3D"}}
		e.EncodeElement(m.Bubble3D, sebubble3D)
	}
	if m.BubbleScale != nil {
		sebubbleScale := xml.StartElement{Name: xml.Name{Local: "c:bubbleScale"}}
		e.EncodeElement(m.BubbleScale, sebubbleScale)
	}
	if m.ShowNegBubbles != nil {
		seshowNegBubbles := xml.StartElement{Name: xml.Name{Local: "c:showNegBubbles"}}
		e.EncodeElement(m.ShowNegBubbles, seshowNegBubbles)
	}
	if m.SizeRepresents != nil {
		sesizeRepresents := xml.StartElement{Name: xml.Name{Local: "c:sizeRepresents"}}
		e.EncodeElement(m.SizeRepresents, sesizeRepresents)
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

func (m *CT_BubbleChart) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_BubbleChart:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "varyColors"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "varyColors"}:
				m.VaryColors = NewCT_Boolean()
				if err := d.DecodeElement(m.VaryColors, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "ser"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "ser"}:
				tmp := NewCT_BubbleSer()
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
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "bubble3D"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "bubble3D"}:
				m.Bubble3D = NewCT_Boolean()
				if err := d.DecodeElement(m.Bubble3D, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "bubbleScale"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "bubbleScale"}:
				m.BubbleScale = NewCT_BubbleScale()
				if err := d.DecodeElement(m.BubbleScale, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "showNegBubbles"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "showNegBubbles"}:
				m.ShowNegBubbles = NewCT_Boolean()
				if err := d.DecodeElement(m.ShowNegBubbles, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "sizeRepresents"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "sizeRepresents"}:
				m.SizeRepresents = NewCT_SizeRepresents()
				if err := d.DecodeElement(m.SizeRepresents, &el); err != nil {
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
				goffice.Log("skipping unsupported element on CT_BubbleChart %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_BubbleChart
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_BubbleChart and its children
func (m *CT_BubbleChart) Validate() error {
	return m.ValidateWithPath("CT_BubbleChart")
}

// ValidateWithPath validates the CT_BubbleChart and its children, prefixing error messages with path
func (m *CT_BubbleChart) ValidateWithPath(path string) error {
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
	if m.Bubble3D != nil {
		if err := m.Bubble3D.ValidateWithPath(path + "/Bubble3D"); err != nil {
			return err
		}
	}
	if m.BubbleScale != nil {
		if err := m.BubbleScale.ValidateWithPath(path + "/BubbleScale"); err != nil {
			return err
		}
	}
	if m.ShowNegBubbles != nil {
		if err := m.ShowNegBubbles.ValidateWithPath(path + "/ShowNegBubbles"); err != nil {
			return err
		}
	}
	if m.SizeRepresents != nil {
		if err := m.SizeRepresents.ValidateWithPath(path + "/SizeRepresents"); err != nil {
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
