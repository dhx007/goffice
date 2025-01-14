package wml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/wml"
)

func TestCT_SectPrConstructor(t *testing.T) {
	v := wml.NewCT_SectPr()
	if v == nil {
		t.Errorf("wml.NewCT_SectPr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_SectPr should validate: %s", err)
	}
}

func TestCT_SectPrMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_SectPr()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_SectPr()
	xml.Unmarshal(buf, v2)
}
