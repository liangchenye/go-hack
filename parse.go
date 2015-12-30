-- +goose Up

-- -----------------------------------------------------
-- Table Namespace
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS Namespace (
  id SERIAL PRIMARY KEY,
  name VARCHAR(128) NULL);


-- -----------------------------------------------------
-- Table Layer
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS Layer (
  id SERIAL PRIMARY KEY,
  name VARCHAR(128) NOT NULL UNIQUE,
  engineversion SMALLINT NOT NULL,
  parent_id INT NULL REFERENCES Layer,
  namespace_id INT NULL REFERENCES Namespace);

CREATE INDEX ON Layer (parent_id);
CREATE INDEX ON Layer (namespace_id);

-- -----------------------------------------------------
-- Table Feature
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS Feature (
  id SERIAL PRIMARY KEY,
  namespace_id INT NOT NULL REFERENCES Namespace,
  name VARCHAR(128) NOT NULL,

  UNIQUE (namespace_id, name));

-- -----------------------------------------------------
-- Table FeatureVersion
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS FeatureVersion (
  id SERIAL PRIMARY KEY,
  feature_id INT NOT NULL REFERENCES Feature,
  version VARCHAR(128) NOT NULL);

CREATE INDEX ON FeatureVersion (feature_id);


-- -----------------------------------------------------
-- Table Layer_diff_FeatureVersion
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS Layer_diff_FeatureVersion (
  layer_id INT NOT NULL REFERENCES Layer ON DELETE CASCADE,
  featureversion_id INT NOT NULL REFERENCES FeatureVersion,
  modification VARCHAR(32) NOT NULL,

  PRIMARY KEY (layer_id, featureversion_id));

CREATE INDEX ON Layer_diff_FeatureVersion (layer_id);
CREATE INDEX ON Layer_diff_FeatureVersion (featureversion_id);
CREATE INDEX ON Layer_diff_FeatureVersion (featureversion_id, layer_id);


-- -----------------------------------------------------
-- Table Vulnerability
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS Vulnerability (
  id SERIAL PRIMARY KEY,
  namespace_id INT NOT NULL REFERENCES Namespace,
  name VARCHAR(128) NOT NULL,
  description TEXT NULL,
  link VARCHAR(128) NULL,
  severity VARCHAR(32) NULL,

  UNIQUE (namespace_id, name));


-- -----------------------------------------------------
-- Table Vulnerability_FixedIn_Feature
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS Vulnerability_FixedIn_Feature (
  vulnerability_id INT NOT NULL REFERENCES Vulnerability ON DELETE CASCADE,
  feature_id INT NOT NULL REFERENCES Feature,
  version VARCHAR(128) NOT NULL,

  PRIMARY KEY (vulnerability_id, feature_id));

CREATE INDEX ON Vulnerability_FixedIn_Feature (feature_id, vulnerability_id);

-- -----------------------------------------------------
-- Table Vulnerability_Affects_FeatureVersion
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS Vulnerability_Affects_FeatureVersion (
  vulnerability_id INT NOT NULL REFERENCES Vulnerability ON DELETE CASCADE,
  featureversion_id INT NOT NULL REFERENCES FeatureVersion,

  PRIMARY KEY (vulnerability_id, featureversion_id));

CREATE INDEX ON Vulnerability_Affects_FeatureVersion (featureversion_id, vulnerability_id);


-- -----------------------------------------------------
-- Table KeyValue
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS KeyValue (
  id SERIAL PRIMARY KEY,
  key VARCHAR(128) NOT NULL UNIQUE,
  value TEXT);

-- +goose Down

DROP TABLE IF EXISTS Namespace,
                     Layer,
                     Feature,
                     FeatureVersion,
                     Layer_diff_FeatureVersion,
                     Vulnerability,
                     Vulnerability_FixedIn_Feature,
                     Vulnerability_Affects_FeatureVersion,
                     KeyValue
            CASCADE;



package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"io/ioutil"
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

func apply_os(req Require) (resource Resource){
	var default_url string
//run this first to test
// curl -i -H 'Content-Type: application/json'     -d '{"Distribution":"Ubuntu14.04","Arch":"arm64", "id": "1235"}' http://127.0.0.1:8080/os

	default_url = "http://localhost:8080/os/Ubuntu14.04"
	resp, err := http.Get(default_url)
	defer resp.Body.Close()
	if err != nil {
		// handle error
		fmt.Println("err in get")
		resource.ID = ""
		resource.Msg = "err in get os"
		resource.Status = false
	} else {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			resource.ID = ""
			resource.Msg = "err in read os"
			resource.Status = false
		} else {
//	fmt.Println(string(body))
			json.Unmarshal([]byte(body), &resource)
			resource.Msg = "ok"
			resource.Status = tru
			fmt.Println(resource)
		}	
	}

	return resource
}

func apply_container(req Require) (resource Resource){
	return resource
}

func apply_resource(req Require) (resource Resource){
	if req.Rtype == "os" {
		resource = apply_os(req)
	} else if req.Rtype == "container" {
		resource = apply_container(req)
	} else {
		fmt.Println("What is the new type? How can it pass the validation test")
	}
	return resource
}

func apply_resources(ts_demo TestCase) (resources []Resource){
	for index :=0; index < len(ts_demo.Requires); index++ {
		var resource Resource
		var req Require
		req = ts_demo.Requires[index]
		resource = apply_resource(req)
		if len(resource.ID) > 1 {
			resources = append(resources, resource)	
		}
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
	resources = apply_resources(ts_demo)
	validate, msg = ar_validation(resources)
	if !validate {
		fmt.Println(msg)
		return
	}
	debug_ar(resources)
}

