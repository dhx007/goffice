package chart_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/dml/chart"
)

func TestCT_RelIdConstructor(t *testing.T) {
	v := chart.NewCT_RelId()
	if v == nil {
		t.Errorf("chart.NewCT_RelId must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_RelId should validate: %s", err)
	}
}

func TestCT_RelIdMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_RelId()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_RelId()
	xml.Unmarshal(buf, v2)
}
