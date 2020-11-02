package file

import (
	"fmt"
	"time"

	"github.com/yurifrl/logapi"
)

// Input Is the object value returned, it represent one line of input parsed
type Input struct {
	text  string
	items map[string]logapi.FileParserItem

	// This demands knloge of the inplementated dynamic parsers
	time    time.Time
	details []string
	payload string
	error   bool
}

// IsError returns true is the log line is
func (p *Input) IsError() bool {
	return p.error
}

// Details return the details of a line
func (p *Input) Details() []string {
	return p.details
}

// NewInput Creates a new input
func NewInput(text string) *Input {
	p := &Input{
		text:  text,
		items: make(map[string]logapi.FileParserItem),
	}
	return p
}

func (p *Input) Add(name string, item logapi.FileParserItem) (err error) {
	item.Set(p.text)
	p.items[name] = item
	var buffer interface{}

	// This is ugly, I will no spend to much time here, let it like this for now
	switch name {
	case "time":
		buffer, err = item.Get()
		p.time = buffer.(time.Time)
	case "details":
		buffer, err = item.Get()
		p.details = buffer.([]string)
	case "payload":
		buffer, err = item.Get()
		p.payload = buffer.(string)
	case "error":
		buffer, err = item.Get()
		p.error = buffer.(bool)
	default:
		err = fmt.Errorf("Invalid option passed")
	}

	return err
}

// Get access the raw parser data
func (p *Input) Get(name string) (interface{}, error) {
	s, err := p.items[name].Get()
	if err != nil {
		return nil, fmt.Errorf("Parsing error: `%v` on `%s`", err, name)
	}
	return s, nil
}
