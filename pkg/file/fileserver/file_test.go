package fileserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/go-chi/chi"
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"

	"github.com/yurifrl/logapi/mocks"
)

func TestFileServerGetAll(t *testing.T) {
	// Create test server
	r := chi.NewRouter()
	server := httptest.NewServer(r)
	defer server.Close()

	// Mock Store and server
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockFileStore := mocks.NewMockFileStore(mockCtrl)
	mockFile := mocks.NewMockFile(mockCtrl)

	// Create Item
	item := make(map[string]string)
	item["foo"] = "1"

	// Expectations
	mockFileStore.EXPECT().GetAll().Return(item, nil)

	// Setup
	err := Setup(logrus.New(), r, mockFileStore, mockFile)
	if err != nil {
		t.Error(err)
	}

	// Make request and validate we get back proper response
	e := httpexpect.New(t, server.URL)
	e.GET("/files").Expect().Status(http.StatusOK).JSON().Equal(&item)
}
