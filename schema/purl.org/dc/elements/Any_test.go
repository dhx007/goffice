// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package elements_test

import (
	"encoding/xml"
	"testing"

	"goffice/schema/purl.org/dc/elements"
)

func TestAnyConstructor(t *testing.T) {
	v := elements.NewAny()
	if v == nil {
		t.Errorf("elements.NewAny must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed elements.Any should validate: %s", err)
	}
}

func TestAnyMarshalUnmarshal(t *testing.T) {
	v := elements.NewAny()
	buf, _ := xml.Marshal(v)
	v2 := elements.NewAny()
	xml.Unmarshal(buf, v2)
}
