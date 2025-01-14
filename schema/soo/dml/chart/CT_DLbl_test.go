package chart_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/dml/chart"
)

func TestCT_DLblConstructor(t *testing.T) {
	v := chart.NewCT_DLbl()
	if v == nil {
		t.Errorf("chart.NewCT_DLbl must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_DLbl should validate: %s", err)
	}
}

func TestCT_DLblMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_DLbl()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_DLbl()
	xml.Unmarshal(buf, v2)
}
