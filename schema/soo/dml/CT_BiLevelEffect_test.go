package dml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/dml"
)

func TestCT_BiLevelEffectConstructor(t *testing.T) {
	v := dml.NewCT_BiLevelEffect()
	if v == nil {
		t.Errorf("dml.NewCT_BiLevelEffect must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_BiLevelEffect should validate: %s", err)
	}
}

func TestCT_BiLevelEffectMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_BiLevelEffect()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_BiLevelEffect()
	xml.Unmarshal(buf, v2)
}
