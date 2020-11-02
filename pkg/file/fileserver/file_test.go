package fileserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/yurifrl/logapi/gorestapi"
	"github.com/yurifrl/logapi/mocks"
)

func FileServerGetAll(t *testing.T) {

	// Create test server
	r := chi.NewRouter()
	server := httptest.NewServer(r)
	defer server.Close()

	// Mock Store and server
	ts := new(mocks.)
	err := Setup(r, ts)
	assert.Nil(t, err)

	// Create Item
	i := []*gorestapi.Thing{
		&gorestapi.Thing{
			ID:   "id1",
			Name: "name1",
		},
		&gorestapi.Thing{
			ID:   "id2",
			Name: "name2",
		},
	}

	// Mock call to item store
	ts.On("ThingFind", mock.AnythingOfType("*context.valueCtx")).Once().Return(i, nil)

	// Make request and validate we get back proper response
	e := httpexpect.New(t, server.URL)
	e.GET("/things").Expect().Status(http.StatusOK).JSON().Array().Equal(&i)

	// Check remaining expectations
	ts.AssertExpectations(t)

}
