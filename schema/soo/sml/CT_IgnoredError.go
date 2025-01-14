package sml

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type CT_IgnoredError struct {
	// Sequence of References
	SqrefAttr ST_Sqref
	// Evaluation Error
	EvalErrorAttr *bool
	// Two Digit Text Year
	TwoDigitTextYearAttr *bool
	// Number Stored As Text
	NumberStoredAsTextAttr *bool
	// Formula
	FormulaAttr *bool
	// Formula Range
	FormulaRangeAttr *bool
	// Unlocked Formula
	UnlockedFormulaAttr *bool
	// Empty Cell Reference
	EmptyCellReferenceAttr *bool
	// List Data Validation
	ListDataValidationAttr *bool
	// Calculated Column
	CalculatedColumnAttr *bool
}

func NewCT_IgnoredError() *CT_IgnoredError {
	ret := &CT_IgnoredError{}
	return ret
}

func (m *CT_IgnoredError) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "sqref"},
		Value: fmt.Sprintf("%v", m.SqrefAttr)})
	if m.EvalErrorAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "evalError"},
			Value: fmt.Sprintf("%d", b2i(*m.EvalErrorAttr))})
	}
	if m.TwoDigitTextYearAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "twoDigitTextYear"},
			Value: fmt.Sprintf("%d", b2i(*m.TwoDigitTextYearAttr))})
	}
	if m.NumberStoredAsTextAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "numberStoredAsText"},
			Value: fmt.Sprintf("%d", b2i(*m.NumberStoredAsTextAttr))})
	}
	if m.FormulaAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "formula"},
			Value: fmt.Sprintf("%d", b2i(*m.FormulaAttr))})
	}
	if m.FormulaRangeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "formulaRange"},
			Value: fmt.Sprintf("%d", b2i(*m.FormulaRangeAttr))})
	}
	if m.UnlockedFormulaAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "unlockedFormula"},
			Value: fmt.Sprintf("%d", b2i(*m.UnlockedFormulaAttr))})
	}
	if m.EmptyCellReferenceAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "emptyCellReference"},
			Value: fmt.Sprintf("%d", b2i(*m.EmptyCellReferenceAttr))})
	}
	if m.ListDataValidationAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "listDataValidation"},
			Value: fmt.Sprintf("%d", b2i(*m.ListDataValidationAttr))})
	}
	if m.CalculatedColumnAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "calculatedColumn"},
			Value: fmt.Sprintf("%d", b2i(*m.CalculatedColumnAttr))})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_IgnoredError) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "sqref" {
			parsed, err := ParseSliceST_Sqref(attr.Value)
			if err != nil {
				return err
			}
			m.SqrefAttr = parsed
			continue
		}
		if attr.Name.Local == "evalError" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.EvalErrorAttr = &parsed
			continue
		}
		if attr.Name.Local == "twoDigitTextYear" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.TwoDigitTextYearAttr = &parsed
			continue
		}
		if attr.Name.Local == "numberStoredAsText" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.NumberStoredAsTextAttr = &parsed
			continue
		}
		if attr.Name.Local == "formula" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.FormulaAttr = &parsed
			continue
		}
		if attr.Name.Local == "formulaRange" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.FormulaRangeAttr = &parsed
			continue
		}
		if attr.Name.Local == "unlockedFormula" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.UnlockedFormulaAttr = &parsed
			continue
		}
		if attr.Name.Local == "emptyCellReference" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.EmptyCellReferenceAttr = &parsed
			continue
		}
		if attr.Name.Local == "listDataValidation" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ListDataValidationAttr = &parsed
			continue
		}
		if attr.Name.Local == "calculatedColumn" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.CalculatedColumnAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_IgnoredError: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_IgnoredError and its children
func (m *CT_IgnoredError) Validate() error {
	return m.ValidateWithPath("CT_IgnoredError")
}

// ValidateWithPath validates the CT_IgnoredError and its children, prefixing error messages with path
func (m *CT_IgnoredError) ValidateWithPath(path string) error {
	return nil
}
