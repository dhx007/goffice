package dml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/dml"
)

func TestCT_TextBulletSizePercentConstructor(t *testing.T) {
	v := dml.NewCT_TextBulletSizePercent()
	if v == nil {
		t.Errorf("dml.NewCT_TextBulletSizePercent must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_TextBulletSizePercent should validate: %s", err)
	}
}

func TestCT_TextBulletSizePercentMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_TextBulletSizePercent()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_TextBulletSizePercent()
	xml.Unmarshal(buf, v2)
}
