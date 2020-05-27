package utils

import (
    "encoding/json"
    "github.com/mitchellh/mapstructure"
    "testing"
)

type People struct {
    Name string `json:"name_title"`
    Age  int    `json:"age_size"`
}

type Person struct {
    Name   string `json:"name"`
    Age    int    `json:"age"`
    School School `json:"school"`
}

type School struct {
    Addr string `json:"addr"`
    //His  string `json:"his"`
}

func TestJsonToStructDemo(t *testing.T) {
    jsonStr := `
        {
                "name_title": "jqw",
                "age_size":12,
                "school": {"addr":"wenshanlu", "his":"ming"}
        }
        `
    var people Person
    if err := json.Unmarshal([]byte(jsonStr), &people); err != nil {
        t.Error(err)
    }
    t.Log(people)
}

func TestStructToJsonDemo(t *testing.T) {
    p := People{
        Name: "jqw",
        Age:  18,
    }

    jsonBytes, err := json.Marshal(p)
    if err != nil {
        t.Error(err)
    }
    t.Log(string(jsonBytes))
}

func TestJsonToMapDemo(t *testing.T) {
    jsonStr := `
        {
                "name": "jqw",
                "school": {"name":"zzz"}
        }
        `
    var mapResult map[string]interface{}
    err := json.Unmarshal([]byte(jsonStr), &mapResult)
    if err != nil {
        t.Error("JsonToMapDemo err: ", err)
    }
    t.Log(mapResult)
}

func TestMapToJsonDemo(t *testing.T) {
    mapInstances := []map[string]interface{}{}
    instance_1 := map[string]interface{}{"name": "John", "age": 10}
    instance_2 := map[string]interface{}{"name": "Alex", "age": 12}
    mapInstances = append(mapInstances, instance_1, instance_2)

    jsonStr, err := json.Marshal(mapInstances)

    if err != nil {
        t.Error("MapToJsonDemo err: ", err)
    }
    t.Log(string(jsonStr))
}

func TestMapToStructDemo(t *testing.T) {
    mapInstance := make(map[string]interface{})
    subMap := map[string]interface{}{"addr": "wensanlu", "his": "ming"}
    mapInstance["Name"] = "jqw"
    mapInstance["Age"] = 18
    mapInstance["school"] = subMap

    var people Person
    err := mapstructure.Decode(mapInstance, &people)
    if err != nil {
        t.Error(err)
    }
    t.Log(people)
}
