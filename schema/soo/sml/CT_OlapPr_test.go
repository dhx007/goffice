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

func TestCT_OlapPrConstructor(t *testing.T) {
	v := sml.NewCT_OlapPr()
	if v == nil {
		t.Errorf("sml.NewCT_OlapPr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_OlapPr should validate: %s", err)
	}
}

func TestCT_OlapPrMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_OlapPr()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_OlapPr()
	xml.Unmarshal(buf, v2)
}
