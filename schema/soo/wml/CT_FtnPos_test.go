package wml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/wml"
)

func TestCT_FtnPosConstructor(t *testing.T) {
	v := wml.NewCT_FtnPos()
	if v == nil {
		t.Errorf("wml.NewCT_FtnPos must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_FtnPos should validate: %s", err)
	}
}

func TestCT_FtnPosMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_FtnPos()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_FtnPos()
	xml.Unmarshal(buf, v2)
}
