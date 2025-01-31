package sml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/sml"
)

func TestCT_IConstructor(t *testing.T) {
	v := sml.NewCT_I()
	if v == nil {
		t.Errorf("sml.NewCT_I must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_I should validate: %s", err)
	}
}

func TestCT_IMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_I()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_I()
	xml.Unmarshal(buf, v2)
}
