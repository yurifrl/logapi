package parsers

import (
	"strings"
	"time"
)

type Date struct {
	value time.Time
	error error
}

func (d *Date) Get() (interface{}, error) {
	return d.value, d.error
}

func (d *Date) Set(text string) error {
	logSlice := strings.SplitN(text, " ", 2)
	logCreationTime, err := time.Parse(time.RFC3339, logSlice[0])
	if err != nil {
		d.error = err
		return err
	}

	d.value = logCreationTime

	return nil
}
