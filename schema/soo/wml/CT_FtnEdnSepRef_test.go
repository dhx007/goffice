package wml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/wml"
)

func TestCT_FtnEdnSepRefConstructor(t *testing.T) {
	v := wml.NewCT_FtnEdnSepRef()
	if v == nil {
		t.Errorf("wml.NewCT_FtnEdnSepRef must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_FtnEdnSepRef should validate: %s", err)
	}
}

func TestCT_FtnEdnSepRefMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_FtnEdnSepRef()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_FtnEdnSepRef()
	xml.Unmarshal(buf, v2)
}
