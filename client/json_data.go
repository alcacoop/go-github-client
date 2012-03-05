// Copyright 2012 Alca Società Cooperativa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package client

/* 
JsonData contains some helpers to help developers accessing a generic parsed json.

As an example:

  rawData := "{ \"attr1\": \"string1\", \"attr2\": 2, \"attr3\": 5.34, \"attr4\": { \"name\": \"obj1\" } }"
  var jd JsonData = make(JsonData)

  _ = json.Unmarshal(([]byte)(rawData), &jd)
  jd.GetString("attr1")
  jd.GetInt("attr2") 
  jd.GetFloat("attr3")
  jd.GetObject("attr4").GetString("name")

NOTE: JsonData.Get* Function can raise an error due to incompatible interface 
conversions.

*/
type JsonData map[string]interface{}

func (j JsonData) GetString(attr string) string {
	return j[attr].(string)
}

func (j JsonData) GetBool(attr string) bool {
	return j[attr].(bool)
}

func (j JsonData) GetFloat(attr string) float64 {
	return j[attr].(float64)
}

func (j JsonData) GetInt(attr string) int {
	return int(j[attr].(float64))
}

func (j JsonData) GetObject(attr string) JsonData {
	return (JsonData)(j[attr].(map[string]interface{}))
}
