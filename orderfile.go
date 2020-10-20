package order

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Orderfile represents single instance of orderfile
type Orderfile struct {
	Version string `yaml:"version"`
	Orders  Orders `yaml:"orders"`
}

func NewOrderFileFrom(path string) (*Orderfile, error) {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	o := &Orderfile{}

	if err := yaml.UnmarshalStrict(yamlFile, o); err != nil {
		return nil, err
	}

	return o, nil
}

func (o *Orderfile) GetOrder(orderName string) (*Order, bool) {
	order, ok := o.Orders[orderName]
	return &order, ok
}

func (o *Orderfile) ListOrdersNames() []string {
	names := make([]string, 0, len(o.Orders))
	for name := range o.Orders {
		names = append(names, name)
	}

	return names
}

// Orders represents orders section of file
type Orders map[string]Order

// Order represents single command order
type Order struct {
	Description string `yaml:"description"`
	Script      []Cmd  `yaml:"script"`
}

// Cmd represents single script command
type Cmd string
