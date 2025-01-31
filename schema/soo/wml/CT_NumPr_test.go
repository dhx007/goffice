package wml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/wml"
)

func TestCT_NumPrConstructor(t *testing.T) {
	v := wml.NewCT_NumPr()
	if v == nil {
		t.Errorf("wml.NewCT_NumPr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_NumPr should validate: %s", err)
	}
}

func TestCT_NumPrMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_NumPr()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_NumPr()
	xml.Unmarshal(buf, v2)
}
