package parsers

import "strings"

type Payload struct {
	value string
}

func (d *Payload) Get() (interface{}, error) {
	return d.value, nil
}

func (d *Payload) Set(text string) error {
	d.value = strings.Split(text, "]:")[1]

	return nil
}
