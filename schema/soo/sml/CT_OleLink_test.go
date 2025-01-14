package sml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/sml"
)

func TestCT_OleLinkConstructor(t *testing.T) {
	v := sml.NewCT_OleLink()
	if v == nil {
		t.Errorf("sml.NewCT_OleLink must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_OleLink should validate: %s", err)
	}
}

func TestCT_OleLinkMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_OleLink()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_OleLink()
	xml.Unmarshal(buf, v2)
}
