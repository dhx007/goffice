package wml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/wml"
)

func TestEG_PContentBaseConstructor(t *testing.T) {
	v := wml.NewEG_PContentBase()
	if v == nil {
		t.Errorf("wml.NewEG_PContentBase must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.EG_PContentBase should validate: %s", err)
	}
}

func TestEG_PContentBaseMarshalUnmarshal(t *testing.T) {
	v := wml.NewEG_PContentBase()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewEG_PContentBase()
	xml.Unmarshal(buf, v2)
}
