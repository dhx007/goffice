package pml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/pml"
)

func TestEG_SlideListChoiceConstructor(t *testing.T) {
	v := pml.NewEG_SlideListChoice()
	if v == nil {
		t.Errorf("pml.NewEG_SlideListChoice must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.EG_SlideListChoice should validate: %s", err)
	}
}

func TestEG_SlideListChoiceMarshalUnmarshal(t *testing.T) {
	v := pml.NewEG_SlideListChoice()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewEG_SlideListChoice()
	xml.Unmarshal(buf, v2)
}
