package chart_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/dml/chart"
)

func TestCT_DLblChoiceConstructor(t *testing.T) {
	v := chart.NewCT_DLblChoice()
	if v == nil {
		t.Errorf("chart.NewCT_DLblChoice must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_DLblChoice should validate: %s", err)
	}
}

func TestCT_DLblChoiceMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_DLblChoice()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_DLblChoice()
	xml.Unmarshal(buf, v2)
}
