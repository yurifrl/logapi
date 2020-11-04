package file

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math"
	"strings"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/afero"
	"github.com/yurifrl/logapi"
)

var (
	chunkSize = 4 * 1024
	lineSize  = 250 * 1024
	linesPool = sync.Pool{New: func() interface{} {
		lines := make([]byte, lineSize)
		return lines
	}}
	stringPool = sync.Pool{New: func() interface{} {
		lines := ""
		return lines
	}}
	EOF = errors.New("EOF")
)

// File ...
// @TODO: Use go interface for store and for parser
type File struct {
	fs       afero.Fs
	logger   *logrus.Logger
	store    logapi.FileStore
	lastSync time.Time
}

// New creates new file
func New(fs afero.Fs, logger *logrus.Logger, store logapi.FileStore) *File {
	f := &File{
		fs:     fs,
		logger: logger,
		store:  store,
	}

	return f
}

// Sync Will index the file into the store
func (f *File) Sync(fileName string) error {
	file, err := f.fs.Open(fileName)
	if err != nil {
		return fmt.Errorf("File could not be opened `%v`", err)
	}

	defer file.Close()

	// Read the first line of the file
	reader := bufio.NewReader(file)

	if err = f.process(reader); err != nil {
		return nil
	}

	f.lastSync = time.Now()

	return err
}

func (f *File) process(file *bufio.Reader) (err error) {
	// wait group to keep track off all threads

	var wg sync.WaitGroup

	// Build chuck to process
	for {
		buf := linesPool.Get().([]byte)
		n, err := file.Read(buf)
		buf = buf[:n]

		if n == 0 {
			if err == io.EOF {
				break
			}
			if err != nil {
				f.logger.Error(err)
				break
			}
			return err
		}

		nextUntillNewline, err := file.ReadBytes('\n')

		if err != io.EOF {
			buf = append(buf, nextUntillNewline...)
		}

		wg.Add(1)
		go func() {
			// Process each chunk concurrently
			f.processChunk(buf)
			wg.Done()
		}()
	}

	wg.Wait()
	return err
}

func (f *File) processChunk(chunk []byte) {

	//another wait group to process every chunk further
	var wg2 sync.WaitGroup

	logs := stringPool.Get().(string)
	logs = string(chunk)

	linesPool.Put(chunk)

	// Split the string by "\n", so that we have slice of logs
	logsSlice := strings.Split(logs, "\n")

	stringPool.Put(logs)

	// Process the bunch of chunckSize logs in thread
	n := len(logsSlice)
	numOfThreads := n / chunkSize

	if n%chunkSize != 0 {
		numOfThreads++
	}

	// Traverse the chunk
	for i := 0; i < (numOfThreads); i++ {
		wg2.Add(1)
		// process each chunk in separate chunk
		// passing the indexes for processing
		untilIndex := i * chunkSize
		endIndex := int(math.Min(float64((i+1)*chunkSize), float64(len(logsSlice))))
		go func(s int, e int) {
			defer wg2.Done() //to avoid deadlocks
			for i := s; i < e; i++ {
				text := logsSlice[i]
				if len(text) == 0 {
					continue
				}

				// Save
				if err := f.Save(text); err != nil {
					// If error, log but continue
					f.logger.Error(err)
				}
			}
		}(untilIndex, endIndex)
	}

	//wait for a chunk to finish
	wg2.Wait()
	logsSlice = nil
}

func (f *File) Save(text string) (err error) {
	// Do something with the logs here
	parsedText, err := Parse(text)
	if err != nil {
		return fmt.Errorf("Failed to parse the string: `%v` with error `%v`", text, err)
	}

	if !parsedText.IsError() {
		return nil
	}

	// Save if it's an error
	if err := f.store.Bump(parsedText.Details()); err != nil {
		return fmt.Errorf("Failed to store with error `%v`", err)
	}

	return err
}
