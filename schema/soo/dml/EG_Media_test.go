package dml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/dml"
)

func TestEG_MediaConstructor(t *testing.T) {
	v := dml.NewEG_Media()
	if v == nil {
		t.Errorf("dml.NewEG_Media must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.EG_Media should validate: %s", err)
	}
}

func TestEG_MediaMarshalUnmarshal(t *testing.T) {
	v := dml.NewEG_Media()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewEG_Media()
	xml.Unmarshal(buf, v2)
}
