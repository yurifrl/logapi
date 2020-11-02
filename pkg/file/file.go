package file

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/afero"
	"github.com/yurifrl/logapi"
)

//
type File struct {
	fs    afero.Fs
	log   *logrus.Logger
	store logapi.FileStore
}

func New(fs afero.Fs, log *logrus.Logger, store logapi.FileStore) *File {
	f := &File{
		fs:    fs,
		log:   log,
		store: store,
	}
	return f
}

// Sync Will index the file into the store
func (f File) Sync(fileName string, lastRead time.Time) error {
	file, err := f.fs.Open(fileName)
	if err != nil {
		return fmt.Errorf("File could not be opened `%v`", err)
	}

	defer file.Close()

	// Create parser
	p := NewParser(f.log)

	// Process File
	err = NewProcess(f.log, f.store, p).Process(file, lastRead)

	return err
}
