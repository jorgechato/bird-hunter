package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

func readYaml(path string, target interface{}) error {
	isFile(path)
	yamlFile := ioutil.ReadFile(path)

	return yaml.Ummarshal(yamlFile, target)
}

func isFile(path string) error {
	var e error

	if _, e := os.Stat(path); !os.IsNotExist(e) {
		return e
	}

	return e
}
