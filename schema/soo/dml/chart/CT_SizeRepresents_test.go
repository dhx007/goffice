package chart_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/dml/chart"
)

func TestCT_SizeRepresentsConstructor(t *testing.T) {
	v := chart.NewCT_SizeRepresents()
	if v == nil {
		t.Errorf("chart.NewCT_SizeRepresents must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_SizeRepresents should validate: %s", err)
	}
}

func TestCT_SizeRepresentsMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_SizeRepresents()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_SizeRepresents()
	xml.Unmarshal(buf, v2)
}
