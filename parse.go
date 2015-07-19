package main

import (
	"fmt"
	"encoding/json"
)

type Require struct {
	Rtype string
	Rname string
}

type Resource struct {
	Req Require
//TODO: put following to a struct and make a hash?
	ID  string	//returned 
	Status bool	//whether it is available
	Msg string	//return value from server
}

type TestCase struct {
	Name string
	License string
	Group string
	Requires []Require
}

const test_json_str = `{"Name": "Network-performance", "License": "GPL",
"Requires": [{"Rtype": "os", "Rname": "CentOS7.0"}, {"Rtype": "container", "Rname": "DockerV1.0"}
]}`

func parse(ts_str string) (ts_demo TestCase) {
	json.Unmarshal([]byte(ts_str), &ts_demo)

	return ts_demo
}

func ts_validation(ts_demo TestCase) (validate bool, err_string string){
	if len(ts_demo.Name) > 0 {
	} else {
		err_string = "Cannot find the name"
		return false, err_string
	}

	if len(ts_demo.Requires) > 0 {
	} else {
		err_string = "Cannot find the Required resource"
		return false, err_string
	}
	return true, "OK"
}

func debug_ts(ts_demo TestCase) {
	fmt.Println(ts_demo.Name)
	fmt.Println(ts_demo.Group)
	fmt.Println(ts_demo.Requires)
}

func apply_resource(ts_demo TestCase) (resources []Resource){
	for index :=0; index < len(ts_demo.Requires); index++ {
		var resource Resource
		resource.Req = ts_demo.Requires[index]
		if resource.Req.Rtype == "os"{
			fmt.Println("os")
//TODO, here we should add the restful api to get the result from test sever!
		} else if resource.Req.Rtype == "container" {
			
//TODO, here we should add the restful api to get the result from container pool!
		} else {
			fmt.Println("What is the new type? How can it pass the validation test")
			continue
		}
		resources = append(resources, resource)
	}
	return resources
}

func ar_validation(ar_demo []Resource) (validate bool, err_string string){
	return true, "OK"
}

func debug_ar(ar_demo []Resource) {
	fmt.Println(ar_demo)
}

func main() {
	var ts_demo TestCase
	var validate bool
	var msg string

        ts_demo = parse(test_json_str)
	validate, msg = ts_validation(ts_demo)
	if !validate {
		fmt.Println(msg)
		return
	}
	debug_ts(ts_demo)

	var resources []Resource
//TODO: async in the future
	resources = apply_resource(ts_demo)
	validate, msg = ar_validation(resources)
	if !validate {
		fmt.Println(msg)
		return
	}
	debug_ar(resources)
}

