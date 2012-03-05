// Copyright 2012 Alca Società Cooperativa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package client

import (
	"testing"
	"encoding/json"
)

func TestJsonData(t *testing.T) {
    rawData := "{ \"attr1\": \"string1\", \"attr2\": 2, \"attr3\": 5.34, \"attr4\": { \"name\": \"obj1\" } }"
	var jd JsonData = make(JsonData)

    if err := json.Unmarshal(([]byte)(rawData), &jd); err != nil {
		t.Fatal()
	}

	if jd.GetString("attr1") != "string1" ||
		jd.GetInt("attr2") != 2 ||
		jd.GetFloat("attr3") != 5.34 ||
		jd.GetObject("attr4").GetString("name") != "obj1" {
		t.Fatal()
	}	
}