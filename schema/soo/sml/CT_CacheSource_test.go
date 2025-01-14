package sml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/sml"
)

func TestCT_CacheSourceConstructor(t *testing.T) {
	v := sml.NewCT_CacheSource()
	if v == nil {
		t.Errorf("sml.NewCT_CacheSource must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_CacheSource should validate: %s", err)
	}
}

func TestCT_CacheSourceMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_CacheSource()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_CacheSource()
	xml.Unmarshal(buf, v2)
}
