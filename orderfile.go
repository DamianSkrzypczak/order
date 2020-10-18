package order

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func NewOrderFileFrom(path string) (*Orderfile, error) {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	o := &Orderfile{}

	if err := yaml.Unmarshal(yamlFile, o); err != nil {
		return nil, err
	}

	return o, nil
}
