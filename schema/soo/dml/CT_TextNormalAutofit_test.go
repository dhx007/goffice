package dml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/dml"
)

func TestCT_TextNormalAutofitConstructor(t *testing.T) {
	v := dml.NewCT_TextNormalAutofit()
	if v == nil {
		t.Errorf("dml.NewCT_TextNormalAutofit must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_TextNormalAutofit should validate: %s", err)
	}
}

func TestCT_TextNormalAutofitMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_TextNormalAutofit()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_TextNormalAutofit()
	xml.Unmarshal(buf, v2)
}
