package chart_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/dml/chart"
)

func TestCT_SurfaceConstructor(t *testing.T) {
	v := chart.NewCT_Surface()
	if v == nil {
		t.Errorf("chart.NewCT_Surface must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_Surface should validate: %s", err)
	}
}

func TestCT_SurfaceMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_Surface()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_Surface()
	xml.Unmarshal(buf, v2)
}
