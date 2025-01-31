package dml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/dml"
)

func TestCT_PercentageConstructor(t *testing.T) {
	v := dml.NewCT_Percentage()
	if v == nil {
		t.Errorf("dml.NewCT_Percentage must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_Percentage should validate: %s", err)
	}
}

func TestCT_PercentageMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_Percentage()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_Percentage()
	xml.Unmarshal(buf, v2)
}
