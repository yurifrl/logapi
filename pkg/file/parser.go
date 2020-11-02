package file

import (
	"github.com/sirupsen/logrus"
	"github.com/yurifrl/logapi"
	"github.com/yurifrl/logapi/pkg/file/parsers"
)

// Parser is the core type for parsing
type Parser struct {
	log *logrus.Logger
}

// New Creates a new parser
func NewParser(log *logrus.Logger) *Parser {
	p := &Parser{
		log: log,
	}
	return p
}

// Parse parses a line
func (p *Parser) Parse(text string) (input logapi.FileParserInput, err error) {
	input = NewInput(text)

	err = input.Add("time", &parsers.Date{})
	if err != nil {
		return nil, err
	}
	err = input.Add("payload", &parsers.Payload{})
	if err != nil {
		return nil, err
	}
	err = input.Add("details", &parsers.Details{})
	if err != nil {
		return nil, err
	}
	err = input.Add("error", &parsers.Error{})
	if err != nil {
		return nil, err
	}

	return input, nil
}
