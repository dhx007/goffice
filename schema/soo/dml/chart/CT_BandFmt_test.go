package chart_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/dml/chart"
)

func TestCT_BandFmtConstructor(t *testing.T) {
	v := chart.NewCT_BandFmt()
	if v == nil {
		t.Errorf("chart.NewCT_BandFmt must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_BandFmt should validate: %s", err)
	}
}

func TestCT_BandFmtMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_BandFmt()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_BandFmt()
	xml.Unmarshal(buf, v2)
}
