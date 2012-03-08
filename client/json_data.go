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
type JsonMap map[string]interface{}

func (j JsonMap) GetString(attr string) string {
	return j[attr].(string)
}

func (j JsonMap) GetBool(attr string) bool {
	return j[attr].(bool)
}

func (j JsonMap) GetFloat(attr string) float64 {
	return j[attr].(float64)
}

func (j JsonMap) GetInt(attr string) int {
	return int(j[attr].(float64))
}

func (j JsonMap) GetMap(attr string) JsonMap {
	return (JsonMap)(j[attr].(map[string]interface{}))
}

func (j JsonMap) GetArray(attr string) JsonArray {
	return (JsonArray)(j[attr].([]interface{}))
}

type JsonArray []interface{}

func (j JsonArray) GetString(idx int) string {
	return j[idx].(string)
}

func (j JsonArray) GetBool(idx int) bool {
	return j[idx].(bool)
}

func (j JsonArray) GetFloat(idx int) float64 {
	return j[idx].(float64)
}

func (j JsonArray) GetInt(idx int) int {
	return int(j[idx].(float64))
}

func (j JsonArray) GetObject(idx int) JsonMap {
	return (JsonMap)(j[idx].(map[string]interface{}))
}

func (j JsonArray) GetArray(idx int) JsonArray {
	return (JsonArray)(j[idx].([]interface{}))
}
