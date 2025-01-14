package dml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/dml"
)

func TestCT_ColorMappingOverrideConstructor(t *testing.T) {
	v := dml.NewCT_ColorMappingOverride()
	if v == nil {
		t.Errorf("dml.NewCT_ColorMappingOverride must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_ColorMappingOverride should validate: %s", err)
	}
}

func TestCT_ColorMappingOverrideMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_ColorMappingOverride()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_ColorMappingOverride()
	xml.Unmarshal(buf, v2)
}
