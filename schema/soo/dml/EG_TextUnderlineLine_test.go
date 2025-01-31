package dml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/dml"
)

func TestEG_TextUnderlineLineConstructor(t *testing.T) {
	v := dml.NewEG_TextUnderlineLine()
	if v == nil {
		t.Errorf("dml.NewEG_TextUnderlineLine must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.EG_TextUnderlineLine should validate: %s", err)
	}
}

func TestEG_TextUnderlineLineMarshalUnmarshal(t *testing.T) {
	v := dml.NewEG_TextUnderlineLine()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewEG_TextUnderlineLine()
	xml.Unmarshal(buf, v2)
}
