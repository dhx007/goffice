package wml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/wml"
)

func TestWdCT_PosVConstructor(t *testing.T) {
	v := wml.NewWdCT_PosV()
	if v == nil {
		t.Errorf("wml.NewWdCT_PosV must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.WdCT_PosV should validate: %s", err)
	}
}

func TestWdCT_PosVMarshalUnmarshal(t *testing.T) {
	v := wml.NewWdCT_PosV()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewWdCT_PosV()
	xml.Unmarshal(buf, v2)
}
