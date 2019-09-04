package main

import (
	"encoding/json"
	"fmt"
)

// type SensorReading struct {
// 	Name       string `json:"name"`
// 	Capacity   int    `json:"capacity"`
// 	Time       string `json:"time"`
// 	Infomation Info   `json:"info"`
// }

// type Info struct {
// 	Description string `json:"desc"`
// }

func main() {
	fmt.Println("hello")

	jsonString := `{"name" : "battery sensor", "capacity": 40, "time": 
	"2019-01-21T19:07:28Z", "info":{
		 "desc": "a sendsor reading"
	}}`

	var reading map[string]interface{}
	err := json.Unmarshal([]byte(jsonString), &reading)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", reading)
}
