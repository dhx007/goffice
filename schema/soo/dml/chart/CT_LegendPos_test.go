package chart_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/dml/chart"
)

func TestCT_LegendPosConstructor(t *testing.T) {
	v := chart.NewCT_LegendPos()
	if v == nil {
		t.Errorf("chart.NewCT_LegendPos must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_LegendPos should validate: %s", err)
	}
}

func TestCT_LegendPosMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_LegendPos()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_LegendPos()
	xml.Unmarshal(buf, v2)
}
