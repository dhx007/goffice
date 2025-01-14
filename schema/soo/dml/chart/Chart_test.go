package chart_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/dml/chart"
)

func TestChartConstructor(t *testing.T) {
	v := chart.NewChart()
	if v == nil {
		t.Errorf("chart.NewChart must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.Chart should validate: %s", err)
	}
}

func TestChartMarshalUnmarshal(t *testing.T) {
	v := chart.NewChart()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewChart()
	xml.Unmarshal(buf, v2)
}
