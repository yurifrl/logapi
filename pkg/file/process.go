package file

import (
	"bufio"
	"io"
	"math"
	"strings"
	"sync"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/yurifrl/logapi"
)

var (
	chunckSize = 4 * 1024
)

// Definitions
type Process struct {
	logger *logrus.Logger
	store  logapi.FileStore
	parser logapi.FileParser
}

func NewProcess(logger *logrus.Logger, store logapi.FileStore, parser logapi.FileParser) *Process {
	return &Process{
		logger: logger,
		store:  store,
		parser: parser,
	}
}

func (p *Process) Process(file io.Reader, start time.Time) (err error) {
	// Sync Pools to reuse memory and decrease GC usage
	linesPool := sync.Pool{New: func() interface{} {
		lines := make([]byte, 250*1024)
		return lines
	}}

	stringPool := sync.Pool{New: func() interface{} {
		lines := ""
		return lines
	}}

	r := bufio.NewReader(file)

	var wg sync.WaitGroup

	for {
		buf := linesPool.Get().([]byte)
		n, err := r.Read(buf)
		buf = buf[:n]

		if n == 0 {
			if err != nil {
				p.logger.Error(err)
				break
			}
			if err == io.EOF {
				break
			}
			return err
		}

		nextUntillNewline, err := r.ReadBytes('\n')

		if err != io.EOF {
			buf = append(buf, nextUntillNewline...)
		}

		wg.Add(1)
		go func() {
			p.processChunk(buf, &linesPool, &stringPool, start)
			wg.Done()
		}()

	}

	wg.Wait()
	return err
}

func (p *Process) processChunk(chunk []byte, linesPool *sync.Pool, stringPool *sync.Pool, start time.Time) {
	var wg2 sync.WaitGroup

	logs := stringPool.Get().(string)
	logs = string(chunk)

	linesPool.Put(chunk)

	logsSlice := strings.Split(logs, "\n")

	stringPool.Put(logs)

	chunkSize := 300
	n := len(logsSlice)
	noOfThread := n / chunkSize

	if n%chunkSize != 0 {
		noOfThread++
	}

	for i := 0; i < (noOfThread); i++ {

		wg2.Add(1)
		go func(s int, e int) {
			defer wg2.Done() //to avaoid deadlocks
			for i := s; i < e; i++ {
				text := logsSlice[i]
				if len(text) == 0 {
					continue
				}

				parsedText, err := p.parser.Parse(text)
				if err != nil {
					p.logger.Errorf("Failed to parse the string: `%v` with error `%v`", text, err)
				}
				// Save if it's an error
				if parsedText.IsError() {
					for _, d := range parsedText.Details() {
						err := p.store.Bump(d)
						if err != nil {
							p.logger.Errorf("Failed to store with error `%v`", err)
						}
					}
				}
			}

		}(i*chunkSize, int(math.Min(float64((i+1)*chunkSize), float64(len(logsSlice)))))
	}

	wg2.Wait()
	logsSlice = nil
}
