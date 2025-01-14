package wml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/wml"
)

func TestCT_TblPrConstructor(t *testing.T) {
	v := wml.NewCT_TblPr()
	if v == nil {
		t.Errorf("wml.NewCT_TblPr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_TblPr should validate: %s", err)
	}
}

func TestCT_TblPrMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_TblPr()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_TblPr()
	xml.Unmarshal(buf, v2)
}
