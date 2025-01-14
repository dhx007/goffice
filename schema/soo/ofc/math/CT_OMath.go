package math

import (
	"encoding/xml"
	"fmt"

	"github.com/dhx007/goffice"
)

type CT_OMath struct {
	EG_OMathMathElements []*EG_OMathMathElements
}

func NewCT_OMath() *CT_OMath {
	ret := &CT_OMath{}
	return ret
}

func (m *CT_OMath) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.EG_OMathMathElements != nil {
		for _, c := range m.EG_OMathMathElements {
			c.MarshalXML(e, xml.StartElement{})
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_OMath) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_OMath:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "acc"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "acc"}:
				tmpomathmathelements := NewEG_OMathMathElements()
				tmpomathmathelements.Acc = NewCT_Acc()
				if err := d.DecodeElement(tmpomathmathelements.Acc, &el); err != nil {
					return err
				}
				m.EG_OMathMathElements = append(m.EG_OMathMathElements, tmpomathmathelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "bar"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "bar"}:
				tmpomathmathelements := NewEG_OMathMathElements()
				tmpomathmathelements.Bar = NewCT_Bar()
				if err := d.DecodeElement(tmpomathmathelements.Bar, &el); err != nil {
					return err
				}
				m.EG_OMathMathElements = append(m.EG_OMathMathElements, tmpomathmathelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "box"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "box"}:
				tmpomathmathelements := NewEG_OMathMathElements()
				tmpomathmathelements.Box = NewCT_Box()
				if err := d.DecodeElement(tmpomathmathelements.Box, &el); err != nil {
					return err
				}
				m.EG_OMathMathElements = append(m.EG_OMathMathElements, tmpomathmathelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "borderBox"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "borderBox"}:
				tmpomathmathelements := NewEG_OMathMathElements()
				tmpomathmathelements.BorderBox = NewCT_BorderBox()
				if err := d.DecodeElement(tmpomathmathelements.BorderBox, &el); err != nil {
					return err
				}
				m.EG_OMathMathElements = append(m.EG_OMathMathElements, tmpomathmathelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "d"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "d"}:
				tmpomathmathelements := NewEG_OMathMathElements()
				tmpomathmathelements.D = NewCT_D()
				if err := d.DecodeElement(tmpomathmathelements.D, &el); err != nil {
					return err
				}
				m.EG_OMathMathElements = append(m.EG_OMathMathElements, tmpomathmathelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "eqArr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "eqArr"}:
				tmpomathmathelements := NewEG_OMathMathElements()
				tmpomathmathelements.EqArr = NewCT_EqArr()
				if err := d.DecodeElement(tmpomathmathelements.EqArr, &el); err != nil {
					return err
				}
				m.EG_OMathMathElements = append(m.EG_OMathMathElements, tmpomathmathelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "f"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "f"}:
				tmpomathmathelements := NewEG_OMathMathElements()
				tmpomathmathelements.F = NewCT_F()
				if err := d.DecodeElement(tmpomathmathelements.F, &el); err != nil {
					return err
				}
				m.EG_OMathMathElements = append(m.EG_OMathMathElements, tmpomathmathelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "func"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "func"}:
				tmpomathmathelements := NewEG_OMathMathElements()
				tmpomathmathelements.Func = NewCT_Func()
				if err := d.DecodeElement(tmpomathmathelements.Func, &el); err != nil {
					return err
				}
				m.EG_OMathMathElements = append(m.EG_OMathMathElements, tmpomathmathelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "groupChr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "groupChr"}:
				tmpomathmathelements := NewEG_OMathMathElements()
				tmpomathmathelements.GroupChr = NewCT_GroupChr()
				if err := d.DecodeElement(tmpomathmathelements.GroupChr, &el); err != nil {
					return err
				}
				m.EG_OMathMathElements = append(m.EG_OMathMathElements, tmpomathmathelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "limLow"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "limLow"}:
				tmpomathmathelements := NewEG_OMathMathElements()
				tmpomathmathelements.LimLow = NewCT_LimLow()
				if err := d.DecodeElement(tmpomathmathelements.LimLow, &el); err != nil {
					return err
				}
				m.EG_OMathMathElements = append(m.EG_OMathMathElements, tmpomathmathelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "limUpp"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "limUpp"}:
				tmpomathmathelements := NewEG_OMathMathElements()
				tmpomathmathelements.LimUpp = NewCT_LimUpp()
				if err := d.DecodeElement(tmpomathmathelements.LimUpp, &el); err != nil {
					return err
				}
				m.EG_OMathMathElements = append(m.EG_OMathMathElements, tmpomathmathelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "m"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "m"}:
				tmpomathmathelements := NewEG_OMathMathElements()
				tmpomathmathelements.M = NewCT_M()
				if err := d.DecodeElement(tmpomathmathelements.M, &el); err != nil {
					return err
				}
				m.EG_OMathMathElements = append(m.EG_OMathMathElements, tmpomathmathelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "nary"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "nary"}:
				tmpomathmathelements := NewEG_OMathMathElements()
				tmpomathmathelements.Nary = NewCT_Nary()
				if err := d.DecodeElement(tmpomathmathelements.Nary, &el); err != nil {
					return err
				}
				m.EG_OMathMathElements = append(m.EG_OMathMathElements, tmpomathmathelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "phant"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "phant"}:
				tmpomathmathelements := NewEG_OMathMathElements()
				tmpomathmathelements.Phant = NewCT_Phant()
				if err := d.DecodeElement(tmpomathmathelements.Phant, &el); err != nil {
					return err
				}
				m.EG_OMathMathElements = append(m.EG_OMathMathElements, tmpomathmathelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "rad"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "rad"}:
				tmpomathmathelements := NewEG_OMathMathElements()
				tmpomathmathelements.Rad = NewCT_Rad()
				if err := d.DecodeElement(tmpomathmathelements.Rad, &el); err != nil {
					return err
				}
				m.EG_OMathMathElements = append(m.EG_OMathMathElements, tmpomathmathelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "sPre"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "sPre"}:
				tmpomathmathelements := NewEG_OMathMathElements()
				tmpomathmathelements.SPre = NewCT_SPre()
				if err := d.DecodeElement(tmpomathmathelements.SPre, &el); err != nil {
					return err
				}
				m.EG_OMathMathElements = append(m.EG_OMathMathElements, tmpomathmathelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "sSub"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "sSub"}:
				tmpomathmathelements := NewEG_OMathMathElements()
				tmpomathmathelements.SSub = NewCT_SSub()
				if err := d.DecodeElement(tmpomathmathelements.SSub, &el); err != nil {
					return err
				}
				m.EG_OMathMathElements = append(m.EG_OMathMathElements, tmpomathmathelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "sSubSup"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "sSubSup"}:
				tmpomathmathelements := NewEG_OMathMathElements()
				tmpomathmathelements.SSubSup = NewCT_SSubSup()
				if err := d.DecodeElement(tmpomathmathelements.SSubSup, &el); err != nil {
					return err
				}
				m.EG_OMathMathElements = append(m.EG_OMathMathElements, tmpomathmathelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "sSup"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "sSup"}:
				tmpomathmathelements := NewEG_OMathMathElements()
				tmpomathmathelements.SSup = NewCT_SSup()
				if err := d.DecodeElement(tmpomathmathelements.SSup, &el); err != nil {
					return err
				}
				m.EG_OMathMathElements = append(m.EG_OMathMathElements, tmpomathmathelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "r"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "r"}:
				tmpomathmathelements := NewEG_OMathMathElements()
				tmpomathmathelements.R = NewCT_R()
				if err := d.DecodeElement(tmpomathmathelements.R, &el); err != nil {
					return err
				}
				m.EG_OMathMathElements = append(m.EG_OMathMathElements, tmpomathmathelements)
			default:
				goffice.Log("skipping unsupported element on CT_OMath %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_OMath
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_OMath and its children
func (m *CT_OMath) Validate() error {
	return m.ValidateWithPath("CT_OMath")
}

// ValidateWithPath validates the CT_OMath and its children, prefixing error messages with path
func (m *CT_OMath) ValidateWithPath(path string) error {
	for i, v := range m.EG_OMathMathElements {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/EG_OMathMathElements[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
