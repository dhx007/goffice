package sml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/sml"
)

func TestCT_DiscretePrConstructor(t *testing.T) {
	v := sml.NewCT_DiscretePr()
	if v == nil {
		t.Errorf("sml.NewCT_DiscretePr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_DiscretePr should validate: %s", err)
	}
}

func TestCT_DiscretePrMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_DiscretePr()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_DiscretePr()
	xml.Unmarshal(buf, v2)
}
