package sml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/sml"
)

func TestAG_RevDataConstructor(t *testing.T) {
	v := sml.NewAG_RevData()
	if v == nil {
		t.Errorf("sml.NewAG_RevData must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.AG_RevData should validate: %s", err)
	}
}

func TestAG_RevDataMarshalUnmarshal(t *testing.T) {
	v := sml.NewAG_RevData()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewAG_RevData()
	xml.Unmarshal(buf, v2)
}
