package logapi

import "time"

//go:generate mockgen -destination=mocks/mock_file_store.go -package=mocks github.com/yurifrl/logapi FileStore
// FileStore ...
type FileStore interface {
	Bump(key string) error
	GetAll() (map[string]int, error)
}

type FileSync interface {
	Sync(fileName string, lastRead time.Time) error
}

type FileParser interface {
	Parse(text string) (FileParserInput, error)
}

type FileParserInput interface {
	IsError() bool
	Details() []string
}

type FileParserItem interface {
	Set(text string) error
	GetTime() (time.Time, error)
	GetDetails() ([]string, error)
	GetTrace() (string, error)
	GetError() (bool, error)
}
