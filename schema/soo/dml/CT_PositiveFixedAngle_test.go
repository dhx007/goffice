package dml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/dml"
)

func TestCT_PositiveFixedAngleConstructor(t *testing.T) {
	v := dml.NewCT_PositiveFixedAngle()
	if v == nil {
		t.Errorf("dml.NewCT_PositiveFixedAngle must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_PositiveFixedAngle should validate: %s", err)
	}
}

func TestCT_PositiveFixedAngleMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_PositiveFixedAngle()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_PositiveFixedAngle()
	xml.Unmarshal(buf, v2)
}
