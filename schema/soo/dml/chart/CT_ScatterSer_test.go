package chart_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/dml/chart"
)

func TestCT_ScatterSerConstructor(t *testing.T) {
	v := chart.NewCT_ScatterSer()
	if v == nil {
		t.Errorf("chart.NewCT_ScatterSer must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_ScatterSer should validate: %s", err)
	}
}

func TestCT_ScatterSerMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_ScatterSer()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_ScatterSer()
	xml.Unmarshal(buf, v2)
}
