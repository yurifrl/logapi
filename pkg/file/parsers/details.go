package parsers

import (
	"regexp"
	"strings"
)

type Details struct {
	value []string
}

func (d *Details) Get() (interface{}, error) {
	return d.value, nil
}

func (d *Details) Set(text string) error {
	r := regexp.MustCompile(`\[([^\[\]]*)\]`)

	submatchall := r.FindAllString(text, -1)
	for _, elm := range submatchall {
		elm = strings.Trim(elm, "[")
		elm = strings.Trim(elm, "]")
		d.value = append(d.value, elm)
	}
	d.value = strings.Split(d.value[0], " ")

	return nil
}
