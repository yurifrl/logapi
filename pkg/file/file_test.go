package file

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"github.com/spf13/afero"
	"github.com/yurifrl/logapi"
	"github.com/yurifrl/logapi/mocks"
)

var (
	filePath = "src/log.txt"
)

func createLog(errors int) string {
	t, _ := time.Parse(time.RFC3339, "2000-07-02T17:54:14.290Z")

	details := []string{"api-gateway", "ffd3082fe09d"}
	log := []string{}
	for ; errors != 0; errors-- {
		t = t.Local().Add(time.Second * 10)
		d := strings.Join(details, " ")
		s := fmt.Sprintf("%s [%s]: ... [error] ... %v", t.Format(time.RFC3339), d, errors)
		log = append(log, s)
	}
	return strings.Join(log, "\n")
}

func newFile(store logapi.FileStore, input string) (*File, error) {
	fs := afero.NewMemMapFs()
	f := File{
		fs:    fs,
		log:   logrus.New(),
		store: store,
	}
	fs.MkdirAll("src", 0755)
	afero.WriteFile(fs, filePath, []byte(input), 0644)
	return &f, nil
}

func TestSomething(t *testing.T) {
	// Mock Store and server
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockFileStore := mocks.NewMockFileStore(mockCtrl)

	// Define expectations
	mockFileStore.EXPECT().Bump(gomock.Any()).Return(nil).MinTimes(10)

	// Setup stubs
	lastRead, _ := time.Parse(time.RFC3339, "2001-07-02T17:54:14.290Z")

	// Create Test items
	var tests = []struct {
		input    string
		content  string
		expected string
	}{
		{input: filePath, content: createLog(10), expected: "[api-gateway]:  1 errors\n[ffd3082fe09d]: 1/1 errors"},
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			f, _ := newFile(mockFileStore, tc.content)
			err := f.Sync(tc.input, lastRead)
			if err != nil {
				t.Errorf("expected '%s', got '%s'", nil, err)
			}
		})
	}
}
