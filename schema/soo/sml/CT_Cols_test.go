package sml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/sml"
)

func TestCT_ColsConstructor(t *testing.T) {
	v := sml.NewCT_Cols()
	if v == nil {
		t.Errorf("sml.NewCT_Cols must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Cols should validate: %s", err)
	}
}

func TestCT_ColsMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Cols()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Cols()
	xml.Unmarshal(buf, v2)
}
