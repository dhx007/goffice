package sml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/sml"
)

func TestCT_LevelGroupConstructor(t *testing.T) {
	v := sml.NewCT_LevelGroup()
	if v == nil {
		t.Errorf("sml.NewCT_LevelGroup must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_LevelGroup should validate: %s", err)
	}
}

func TestCT_LevelGroupMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_LevelGroup()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_LevelGroup()
	xml.Unmarshal(buf, v2)
}
