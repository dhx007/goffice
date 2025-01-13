// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package sml_test

import (
	"encoding/xml"
	"testing"

	"goffice/schema/soo/sml"
)

func TestCT_CfRuleConstructor(t *testing.T) {
	v := sml.NewCT_CfRule()
	if v == nil {
		t.Errorf("sml.NewCT_CfRule must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_CfRule should validate: %s", err)
	}
}

func TestCT_CfRuleMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_CfRule()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_CfRule()
	xml.Unmarshal(buf, v2)
}
