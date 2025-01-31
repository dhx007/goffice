package sml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/sml"
)

func TestCT_ChartsheetPrConstructor(t *testing.T) {
	v := sml.NewCT_ChartsheetPr()
	if v == nil {
		t.Errorf("sml.NewCT_ChartsheetPr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_ChartsheetPr should validate: %s", err)
	}
}

func TestCT_ChartsheetPrMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_ChartsheetPr()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_ChartsheetPr()
	xml.Unmarshal(buf, v2)
}
