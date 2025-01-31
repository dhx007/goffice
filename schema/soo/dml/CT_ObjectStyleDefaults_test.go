package dml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/dml"
)

func TestCT_ObjectStyleDefaultsConstructor(t *testing.T) {
	v := dml.NewCT_ObjectStyleDefaults()
	if v == nil {
		t.Errorf("dml.NewCT_ObjectStyleDefaults must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_ObjectStyleDefaults should validate: %s", err)
	}
}

func TestCT_ObjectStyleDefaultsMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_ObjectStyleDefaults()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_ObjectStyleDefaults()
	xml.Unmarshal(buf, v2)
}
