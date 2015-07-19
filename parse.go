package main

import (
	"fmt"
	"encoding/json"
)

type Require struct {
	Rtype string
	Rname string
}

type TestCase struct {
	Name string
	License string
	Requires []Require
}

const test_json_str = `{"Name": "Network-performance", "License": "GPL","Requires": [{"Rtype": "os", "Rname": "CentOS7.0"}]}`

func parse(ts_str string) (ts_demo TestCase) {
	json.Unmarshal([]byte(ts_str), &ts_demo)

	return ts_demo
}

func debug(ts_demo TestCase) {
//	fmt.Println(ts_demo.Requires[0].Rtype)
	fmt.Println(ts_demo.Name)
}

func main() {
	var ts_demo TestCase
        ts_demo = parse(test_json_str)
	debug(ts_demo)
}

