package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

func readYaml(path string, target interface{}) error {
	isFile(path)
	yamlFile, _ := ioutil.ReadFile(path)

	return yaml.Unmarshal(yamlFile, target)
}

func isFile(path string) error {
	var e error

	if _, e := os.Stat(path); !os.IsNotExist(e) {
		return e
	}

	return e
}
