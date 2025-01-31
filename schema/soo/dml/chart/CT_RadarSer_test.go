package chart_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/dml/chart"
)

func TestCT_RadarSerConstructor(t *testing.T) {
	v := chart.NewCT_RadarSer()
	if v == nil {
		t.Errorf("chart.NewCT_RadarSer must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_RadarSer should validate: %s", err)
	}
}

func TestCT_RadarSerMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_RadarSer()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_RadarSer()
	xml.Unmarshal(buf, v2)
}
