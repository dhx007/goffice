package chart_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/dml/chart"
)

func TestEG_PieChartSharedConstructor(t *testing.T) {
	v := chart.NewEG_PieChartShared()
	if v == nil {
		t.Errorf("chart.NewEG_PieChartShared must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.EG_PieChartShared should validate: %s", err)
	}
}

func TestEG_PieChartSharedMarshalUnmarshal(t *testing.T) {
	v := chart.NewEG_PieChartShared()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewEG_PieChartShared()
	xml.Unmarshal(buf, v2)
}
