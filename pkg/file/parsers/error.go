package parsers

import (
	"strings"
)

type Error struct {
	isError bool
}

func (d *Error) Get() (interface{}, error) {
	return d.isError, nil
}

func (d *Error) Set(text string) error {
	d.isError = strings.Contains(text, "[error]")

	return nil
}
