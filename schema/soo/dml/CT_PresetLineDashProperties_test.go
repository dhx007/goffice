package dml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/dml"
)

func TestCT_PresetLineDashPropertiesConstructor(t *testing.T) {
	v := dml.NewCT_PresetLineDashProperties()
	if v == nil {
		t.Errorf("dml.NewCT_PresetLineDashProperties must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_PresetLineDashProperties should validate: %s", err)
	}
}

func TestCT_PresetLineDashPropertiesMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_PresetLineDashProperties()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_PresetLineDashProperties()
	xml.Unmarshal(buf, v2)
}
