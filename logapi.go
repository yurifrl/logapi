package logapi

//go:generate mockgen -destination=mocks/mock_file_store.go -package=mocks github.com/yurifrl/logapi FileStore
// FileStore ...
type FileStore interface {
	Bump(key string) error
	GetAll() (map[string]int, error)
}

type FileParser interface {
	Parse(text string) (FileParserInput, error)
}

type FileParserInput interface {
	Add(name string, item FileParserItem) error
	IsError() bool
	Details() []string
}

type FileParserItem interface {
	Set(text string) error
	Get() (interface{}, error)
}
