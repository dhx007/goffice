package dml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/dml"
)

func TestCT_XYAdjustHandleConstructor(t *testing.T) {
	v := dml.NewCT_XYAdjustHandle()
	if v == nil {
		t.Errorf("dml.NewCT_XYAdjustHandle must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_XYAdjustHandle should validate: %s", err)
	}
}

func TestCT_XYAdjustHandleMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_XYAdjustHandle()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_XYAdjustHandle()
	xml.Unmarshal(buf, v2)
}
