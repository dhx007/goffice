package sml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/sml"
)

func TestCT_OlapPrConstructor(t *testing.T) {
	v := sml.NewCT_OlapPr()
	if v == nil {
		t.Errorf("sml.NewCT_OlapPr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_OlapPr should validate: %s", err)
	}
}

func TestCT_OlapPrMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_OlapPr()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_OlapPr()
	xml.Unmarshal(buf, v2)
}
