package file

import (
	"regexp"
	"strings"
	"time"

	"github.com/yurifrl/logapi"
)

// Input Is the object value returned, it represent one line of input parsed
type Input struct {
	text string

	time    time.Time
	details []string
	trace   string
	error   bool
}

// Parse parses a line
func Parse(text string) (input logapi.FileParserInput, err error) {
	input, err = NewInput(text)

	return input, err
}

// IsError returns true is the log line is
func (i *Input) IsError() bool {
	return i.error
}

// Details return the details of a line
func (i *Input) Details() []string {
	return i.details
}

// Time return the time of a line
func (i *Input) Time() time.Time {
	return i.time
}

// NewInput Creates a new input
func NewInput(text string) (*Input, error) {
	i := &Input{
		text: text,
	}
	err := i.setAll()
	return i, err
}

func (i *Input) setAll() (err error) {
	if err = i.setDetails(); err != nil {
		return
	}
	if err = i.setTime(); err != nil {
		return
	}

	if err = i.setTrace(); err != nil {
		return
	}

	if err = i.setError(); err != nil {
		return
	}

	return err
}

// Fields

func (i *Input) setTime() error {
	logSlice := strings.SplitN(i.text, " ", 2)
	logCreationTime, err := time.Parse(time.RFC3339, logSlice[0])
	if err != nil {
		return err
	}

	i.time = logCreationTime

	return nil
}

func (i *Input) setDetails() error {
	r := regexp.MustCompile(`\[([^\[\]]*)\]`)

	submatchall := r.FindAllString(i.text, -1)
	for _, elm := range submatchall {
		elm = strings.Trim(elm, "[")
		elm = strings.Trim(elm, "]")
		i.details = append(i.details, elm)
	}
	i.details = strings.Split(i.details[0], " ")

	return nil
}

func (i *Input) setError() error {
	i.error = strings.Contains(i.text, "[error]")

	return nil
}

func (i *Input) setTrace() error {
	i.trace = strings.Split(i.text, "]:")[1]

	return nil
}
