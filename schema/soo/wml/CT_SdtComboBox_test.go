package wml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/wml"
)

func TestCT_SdtComboBoxConstructor(t *testing.T) {
	v := wml.NewCT_SdtComboBox()
	if v == nil {
		t.Errorf("wml.NewCT_SdtComboBox must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_SdtComboBox should validate: %s", err)
	}
}

func TestCT_SdtComboBoxMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_SdtComboBox()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_SdtComboBox()
	xml.Unmarshal(buf, v2)
}
