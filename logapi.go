package logapi

import "time"

//go:generate mockgen -destination=mocks/mock_file_store.go -package=mocks github.com/yurifrl/logapi FileStore
// FileStore ...
type FileStore interface {
	Bump(key []string) error
	GetAll() (map[string]string, error)
}

//go:generate mockgen -destination=mocks/mock_file.go -package=mocks github.com/yurifrl/logapi File
type File interface {
	Sync(fileName string) error
}

type FileParser interface {
	Parse(text string) (FileParserInput, error)
}

type FileParserInput interface {
	IsError() bool
	Details() []string
	Time() time.Time
}
