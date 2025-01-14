package wml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/wml"
)

func TestCT_MathCtrlDelConstructor(t *testing.T) {
	v := wml.NewCT_MathCtrlDel()
	if v == nil {
		t.Errorf("wml.NewCT_MathCtrlDel must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_MathCtrlDel should validate: %s", err)
	}
}

func TestCT_MathCtrlDelMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_MathCtrlDel()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_MathCtrlDel()
	xml.Unmarshal(buf, v2)
}
