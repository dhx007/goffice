package dml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/dml"
)

func TestEG_ThemeableFillStyleConstructor(t *testing.T) {
	v := dml.NewEG_ThemeableFillStyle()
	if v == nil {
		t.Errorf("dml.NewEG_ThemeableFillStyle must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.EG_ThemeableFillStyle should validate: %s", err)
	}
}

func TestEG_ThemeableFillStyleMarshalUnmarshal(t *testing.T) {
	v := dml.NewEG_ThemeableFillStyle()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewEG_ThemeableFillStyle()
	xml.Unmarshal(buf, v2)
}
