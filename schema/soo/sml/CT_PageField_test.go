package sml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/sml"
)

func TestCT_PageFieldConstructor(t *testing.T) {
	v := sml.NewCT_PageField()
	if v == nil {
		t.Errorf("sml.NewCT_PageField must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_PageField should validate: %s", err)
	}
}

func TestCT_PageFieldMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_PageField()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_PageField()
	xml.Unmarshal(buf, v2)
}
