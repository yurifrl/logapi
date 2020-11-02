package file

import (
	"github.com/sirupsen/logrus"
	"github.com/yurifrl/logapi"
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
	input, err = NewInput(text)

	return input, err
}
