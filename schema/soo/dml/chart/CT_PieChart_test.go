package chart_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/dml/chart"
)

func TestCT_PieChartConstructor(t *testing.T) {
	v := chart.NewCT_PieChart()
	if v == nil {
		t.Errorf("chart.NewCT_PieChart must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_PieChart should validate: %s", err)
	}
}

func TestCT_PieChartMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_PieChart()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_PieChart()
	xml.Unmarshal(buf, v2)
}
