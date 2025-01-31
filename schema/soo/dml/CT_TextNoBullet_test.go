package dml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/dml"
)

func TestCT_TextNoBulletConstructor(t *testing.T) {
	v := dml.NewCT_TextNoBullet()
	if v == nil {
		t.Errorf("dml.NewCT_TextNoBullet must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_TextNoBullet should validate: %s", err)
	}
}

func TestCT_TextNoBulletMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_TextNoBullet()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_TextNoBullet()
	xml.Unmarshal(buf, v2)
}
