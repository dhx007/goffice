package chart_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/dml/chart"
)

func TestCT_MarkerSizeConstructor(t *testing.T) {
	v := chart.NewCT_MarkerSize()
	if v == nil {
		t.Errorf("chart.NewCT_MarkerSize must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_MarkerSize should validate: %s", err)
	}
}

func TestCT_MarkerSizeMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_MarkerSize()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_MarkerSize()
	xml.Unmarshal(buf, v2)
}
