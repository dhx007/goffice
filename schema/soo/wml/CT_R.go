package wml

import (
	"encoding/xml"
	"fmt"

	"github.com/dhx007/goffice"
)

type CT_R struct {
	// Revision Identifier for Run Properties
	RsidRPrAttr *string
	// Revision Identifier for Run Deletion
	RsidDelAttr *string
	// Revision Identifier for Run
	RsidRAttr *string
	// Run Properties
	RPr                *CT_RPr
	EG_RunInnerContent []*EG_RunInnerContent
	Extra              []goffice.Any
}

func NewCT_R() *CT_R {
	ret := &CT_R{}
	return ret
}

func (m *CT_R) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.RsidRPrAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:rsidRPr"},
			Value: fmt.Sprintf("%v", *m.RsidRPrAttr)})
	}
	if m.RsidDelAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:rsidDel"},
			Value: fmt.Sprintf("%v", *m.RsidDelAttr)})
	}
	if m.RsidRAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:rsidR"},
			Value: fmt.Sprintf("%v", *m.RsidRAttr)})
	}
	e.EncodeToken(start)
	if m.RPr != nil {
		serPr := xml.StartElement{Name: xml.Name{Local: "w:rPr"}}
		e.EncodeElement(m.RPr, serPr)
	}
	if m.EG_RunInnerContent != nil {
		for _, c := range m.EG_RunInnerContent {
			c.MarshalXML(e, xml.StartElement{})
		}
	}
	for _, any := range m.Extra {
		if err := any.MarshalXML(e, xml.StartElement{}); err != nil {
			return err
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_R) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "rsidRPr" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RsidRPrAttr = &parsed
			continue
		}
		if attr.Name.Local == "rsidDel" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RsidDelAttr = &parsed
			continue
		}
		if attr.Name.Local == "rsidR" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RsidRAttr = &parsed
			continue
		}
	}
lCT_R:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "rPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "rPr"}:
				m.RPr = NewCT_RPr()
				if err := d.DecodeElement(m.RPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "br"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "br"}:
				tmpruninnercontent := NewEG_RunInnerContent()
				tmpruninnercontent.Br = NewCT_Br()
				if err := d.DecodeElement(tmpruninnercontent.Br, &el); err != nil {
					return err
				}
				m.EG_RunInnerContent = append(m.EG_RunInnerContent, tmpruninnercontent)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "t"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "t"}:
				tmpruninnercontent := NewEG_RunInnerContent()
				tmpruninnercontent.T = NewCT_Text()
				if err := d.DecodeElement(tmpruninnercontent.T, &el); err != nil {
					return err
				}
				m.EG_RunInnerContent = append(m.EG_RunInnerContent, tmpruninnercontent)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "contentPart"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "contentPart"}:
				tmpruninnercontent := NewEG_RunInnerContent()
				tmpruninnercontent.ContentPart = NewCT_Rel()
				if err := d.DecodeElement(tmpruninnercontent.ContentPart, &el); err != nil {
					return err
				}
				m.EG_RunInnerContent = append(m.EG_RunInnerContent, tmpruninnercontent)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "delText"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "delText"}:
				tmpruninnercontent := NewEG_RunInnerContent()
				tmpruninnercontent.DelText = NewCT_Text()
				if err := d.DecodeElement(tmpruninnercontent.DelText, &el); err != nil {
					return err
				}
				m.EG_RunInnerContent = append(m.EG_RunInnerContent, tmpruninnercontent)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "instrText"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "instrText"}:
				tmpruninnercontent := NewEG_RunInnerContent()
				tmpruninnercontent.InstrText = NewCT_Text()
				if err := d.DecodeElement(tmpruninnercontent.InstrText, &el); err != nil {
					return err
				}
				m.EG_RunInnerContent = append(m.EG_RunInnerContent, tmpruninnercontent)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "delInstrText"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "delInstrText"}:
				tmpruninnercontent := NewEG_RunInnerContent()
				tmpruninnercontent.DelInstrText = NewCT_Text()
				if err := d.DecodeElement(tmpruninnercontent.DelInstrText, &el); err != nil {
					return err
				}
				m.EG_RunInnerContent = append(m.EG_RunInnerContent, tmpruninnercontent)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "noBreakHyphen"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "noBreakHyphen"}:
				tmpruninnercontent := NewEG_RunInnerContent()
				tmpruninnercontent.NoBreakHyphen = NewCT_Empty()
				if err := d.DecodeElement(tmpruninnercontent.NoBreakHyphen, &el); err != nil {
					return err
				}
				m.EG_RunInnerContent = append(m.EG_RunInnerContent, tmpruninnercontent)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "softHyphen"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "softHyphen"}:
				tmpruninnercontent := NewEG_RunInnerContent()
				tmpruninnercontent.SoftHyphen = NewCT_Empty()
				if err := d.DecodeElement(tmpruninnercontent.SoftHyphen, &el); err != nil {
					return err
				}
				m.EG_RunInnerContent = append(m.EG_RunInnerContent, tmpruninnercontent)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "dayShort"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "dayShort"}:
				tmpruninnercontent := NewEG_RunInnerContent()
				tmpruninnercontent.DayShort = NewCT_Empty()
				if err := d.DecodeElement(tmpruninnercontent.DayShort, &el); err != nil {
					return err
				}
				m.EG_RunInnerContent = append(m.EG_RunInnerContent, tmpruninnercontent)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "monthShort"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "monthShort"}:
				tmpruninnercontent := NewEG_RunInnerContent()
				tmpruninnercontent.MonthShort = NewCT_Empty()
				if err := d.DecodeElement(tmpruninnercontent.MonthShort, &el); err != nil {
					return err
				}
				m.EG_RunInnerContent = append(m.EG_RunInnerContent, tmpruninnercontent)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "yearShort"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "yearShort"}:
				tmpruninnercontent := NewEG_RunInnerContent()
				tmpruninnercontent.YearShort = NewCT_Empty()
				if err := d.DecodeElement(tmpruninnercontent.YearShort, &el); err != nil {
					return err
				}
				m.EG_RunInnerContent = append(m.EG_RunInnerContent, tmpruninnercontent)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "dayLong"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "dayLong"}:
				tmpruninnercontent := NewEG_RunInnerContent()
				tmpruninnercontent.DayLong = NewCT_Empty()
				if err := d.DecodeElement(tmpruninnercontent.DayLong, &el); err != nil {
					return err
				}
				m.EG_RunInnerContent = append(m.EG_RunInnerContent, tmpruninnercontent)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "monthLong"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "monthLong"}:
				tmpruninnercontent := NewEG_RunInnerContent()
				tmpruninnercontent.MonthLong = NewCT_Empty()
				if err := d.DecodeElement(tmpruninnercontent.MonthLong, &el); err != nil {
					return err
				}
				m.EG_RunInnerContent = append(m.EG_RunInnerContent, tmpruninnercontent)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "yearLong"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "yearLong"}:
				tmpruninnercontent := NewEG_RunInnerContent()
				tmpruninnercontent.YearLong = NewCT_Empty()
				if err := d.DecodeElement(tmpruninnercontent.YearLong, &el); err != nil {
					return err
				}
				m.EG_RunInnerContent = append(m.EG_RunInnerContent, tmpruninnercontent)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "annotationRef"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "annotationRef"}:
				tmpruninnercontent := NewEG_RunInnerContent()
				tmpruninnercontent.AnnotationRef = NewCT_Empty()
				if err := d.DecodeElement(tmpruninnercontent.AnnotationRef, &el); err != nil {
					return err
				}
				m.EG_RunInnerContent = append(m.EG_RunInnerContent, tmpruninnercontent)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "footnoteRef"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "footnoteRef"}:
				tmpruninnercontent := NewEG_RunInnerContent()
				tmpruninnercontent.FootnoteRef = NewCT_Empty()
				if err := d.DecodeElement(tmpruninnercontent.FootnoteRef, &el); err != nil {
					return err
				}
				m.EG_RunInnerContent = append(m.EG_RunInnerContent, tmpruninnercontent)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "endnoteRef"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "endnoteRef"}:
				tmpruninnercontent := NewEG_RunInnerContent()
				tmpruninnercontent.EndnoteRef = NewCT_Empty()
				if err := d.DecodeElement(tmpruninnercontent.EndnoteRef, &el); err != nil {
					return err
				}
				m.EG_RunInnerContent = append(m.EG_RunInnerContent, tmpruninnercontent)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "separator"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "separator"}:
				tmpruninnercontent := NewEG_RunInnerContent()
				tmpruninnercontent.Separator = NewCT_Empty()
				if err := d.DecodeElement(tmpruninnercontent.Separator, &el); err != nil {
					return err
				}
				m.EG_RunInnerContent = append(m.EG_RunInnerContent, tmpruninnercontent)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "continuationSeparator"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "continuationSeparator"}:
				tmpruninnercontent := NewEG_RunInnerContent()
				tmpruninnercontent.ContinuationSeparator = NewCT_Empty()
				if err := d.DecodeElement(tmpruninnercontent.ContinuationSeparator, &el); err != nil {
					return err
				}
				m.EG_RunInnerContent = append(m.EG_RunInnerContent, tmpruninnercontent)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "sym"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "sym"}:
				tmpruninnercontent := NewEG_RunInnerContent()
				tmpruninnercontent.Sym = NewCT_Sym()
				if err := d.DecodeElement(tmpruninnercontent.Sym, &el); err != nil {
					return err
				}
				m.EG_RunInnerContent = append(m.EG_RunInnerContent, tmpruninnercontent)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "pgNum"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "pgNum"}:
				tmpruninnercontent := NewEG_RunInnerContent()
				tmpruninnercontent.PgNum = NewCT_Empty()
				if err := d.DecodeElement(tmpruninnercontent.PgNum, &el); err != nil {
					return err
				}
				m.EG_RunInnerContent = append(m.EG_RunInnerContent, tmpruninnercontent)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "cr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "cr"}:
				tmpruninnercontent := NewEG_RunInnerContent()
				tmpruninnercontent.Cr = NewCT_Empty()
				if err := d.DecodeElement(tmpruninnercontent.Cr, &el); err != nil {
					return err
				}
				m.EG_RunInnerContent = append(m.EG_RunInnerContent, tmpruninnercontent)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "tab"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "tab"}:
				tmpruninnercontent := NewEG_RunInnerContent()
				tmpruninnercontent.Tab = NewCT_Empty()
				if err := d.DecodeElement(tmpruninnercontent.Tab, &el); err != nil {
					return err
				}
				m.EG_RunInnerContent = append(m.EG_RunInnerContent, tmpruninnercontent)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "object"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "object"}:
				tmpruninnercontent := NewEG_RunInnerContent()
				tmpruninnercontent.Object = NewCT_Object()
				if err := d.DecodeElement(tmpruninnercontent.Object, &el); err != nil {
					return err
				}
				m.EG_RunInnerContent = append(m.EG_RunInnerContent, tmpruninnercontent)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "pict"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "pict"}:
				tmpruninnercontent := NewEG_RunInnerContent()
				tmpruninnercontent.Pict = NewCT_Picture()
				if err := d.DecodeElement(tmpruninnercontent.Pict, &el); err != nil {
					return err
				}
				m.EG_RunInnerContent = append(m.EG_RunInnerContent, tmpruninnercontent)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "fldChar"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "fldChar"}:
				tmpruninnercontent := NewEG_RunInnerContent()
				tmpruninnercontent.FldChar = NewCT_FldChar()
				if err := d.DecodeElement(tmpruninnercontent.FldChar, &el); err != nil {
					return err
				}
				m.EG_RunInnerContent = append(m.EG_RunInnerContent, tmpruninnercontent)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "ruby"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "ruby"}:
				tmpruninnercontent := NewEG_RunInnerContent()
				tmpruninnercontent.Ruby = NewCT_Ruby()
				if err := d.DecodeElement(tmpruninnercontent.Ruby, &el); err != nil {
					return err
				}
				m.EG_RunInnerContent = append(m.EG_RunInnerContent, tmpruninnercontent)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "footnoteReference"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "footnoteReference"}:
				tmpruninnercontent := NewEG_RunInnerContent()
				tmpruninnercontent.FootnoteReference = NewCT_FtnEdnRef()
				if err := d.DecodeElement(tmpruninnercontent.FootnoteReference, &el); err != nil {
					return err
				}
				m.EG_RunInnerContent = append(m.EG_RunInnerContent, tmpruninnercontent)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "endnoteReference"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "endnoteReference"}:
				tmpruninnercontent := NewEG_RunInnerContent()
				tmpruninnercontent.EndnoteReference = NewCT_FtnEdnRef()
				if err := d.DecodeElement(tmpruninnercontent.EndnoteReference, &el); err != nil {
					return err
				}
				m.EG_RunInnerContent = append(m.EG_RunInnerContent, tmpruninnercontent)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "commentReference"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "commentReference"}:
				tmpruninnercontent := NewEG_RunInnerContent()
				tmpruninnercontent.CommentReference = NewCT_Markup()
				if err := d.DecodeElement(tmpruninnercontent.CommentReference, &el); err != nil {
					return err
				}
				m.EG_RunInnerContent = append(m.EG_RunInnerContent, tmpruninnercontent)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "drawing"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "drawing"}:
				tmpruninnercontent := NewEG_RunInnerContent()
				tmpruninnercontent.Drawing = NewCT_Drawing()
				if err := d.DecodeElement(tmpruninnercontent.Drawing, &el); err != nil {
					return err
				}
				m.EG_RunInnerContent = append(m.EG_RunInnerContent, tmpruninnercontent)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "ptab"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "ptab"}:
				tmpruninnercontent := NewEG_RunInnerContent()
				tmpruninnercontent.Ptab = NewCT_PTab()
				if err := d.DecodeElement(tmpruninnercontent.Ptab, &el); err != nil {
					return err
				}
				m.EG_RunInnerContent = append(m.EG_RunInnerContent, tmpruninnercontent)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "lastRenderedPageBreak"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "lastRenderedPageBreak"}:
				tmpruninnercontent := NewEG_RunInnerContent()
				tmpruninnercontent.LastRenderedPageBreak = NewCT_Empty()
				if err := d.DecodeElement(tmpruninnercontent.LastRenderedPageBreak, &el); err != nil {
					return err
				}
				m.EG_RunInnerContent = append(m.EG_RunInnerContent, tmpruninnercontent)
			default:
				any := &goffice.XSDAny{}
				if err := d.DecodeElement(any, &el); err != nil {
					return err
				}
				m.Extra = append(m.Extra, any)
			}
		case xml.EndElement:
			break lCT_R
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_R and its children
func (m *CT_R) Validate() error {
	return m.ValidateWithPath("CT_R")
}

// ValidateWithPath validates the CT_R and its children, prefixing error messages with path
func (m *CT_R) ValidateWithPath(path string) error {
	if m.RPr != nil {
		if err := m.RPr.ValidateWithPath(path + "/RPr"); err != nil {
			return err
		}
	}
	for i, v := range m.EG_RunInnerContent {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/EG_RunInnerContent[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
