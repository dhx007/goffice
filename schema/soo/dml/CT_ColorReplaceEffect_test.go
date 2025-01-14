package dml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/dml"
)

func TestCT_ColorReplaceEffectConstructor(t *testing.T) {
	v := dml.NewCT_ColorReplaceEffect()
	if v == nil {
		t.Errorf("dml.NewCT_ColorReplaceEffect must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_ColorReplaceEffect should validate: %s", err)
	}
}

func TestCT_ColorReplaceEffectMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_ColorReplaceEffect()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_ColorReplaceEffect()
	xml.Unmarshal(buf, v2)
}
