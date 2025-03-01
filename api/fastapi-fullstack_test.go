package api

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	middleware "github.com/oapi-codegen/nethttp-middleware"
	"github.com/oapi-codegen/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func doGet(t *testing.T, mux *chi.Mux, url string) *httptest.ResponseRecorder {
	response := testutil.NewRequest().Get(url).WithAcceptJson().GoWithHTTPHandler(t, mux)
	return response.Recorder
}
func TestPetStore(t *testing.T) {
	var err error

	// Get the swagger description of our API
	swagger, err := GetSwagger()
	require.NoError(t, err)

	// Clear out the servers array in the swagger spec, that skips validating
	// that server names match. We don't know how this thing will be run.
	swagger.Servers = nil

	// This is how you set up a basic chi router
	r := chi.NewRouter()

	// Use our validation middleware to check all requests against the
	// OpenAPI schema.
	r.Use(middleware.OapiRequestValidator(swagger))

	fastApi := NewFastAPI()
	HandlerFromMux(NewStrictHandler(fastApi, nil), r)

	t.Run("Health Check", func(t *testing.T) {
		//tag := "TagOfSpot"
		//newPet := api.NewPet{
		//	Name: "Spot",
		//	Tag:  &tag,
		//}

		//rr := testutil.NewRequest().Post("/pets").WithJsonBody(newPet).GoWithHTTPHandler(t, r).Recorder
		rr := doGet(t, r, "/api/v1/utils/health-check/")
		assert.Equal(t, http.StatusOK, rr.Code)

		var status bool
		err = json.NewDecoder(rr.Body).Decode(&status)
		assert.NoError(t, err, "error unmarshaling response")
		assert.Equal(t, true, status)
	})

	t.Run("GET /api/v1/items/", func(t *testing.T) {
		rr := doGet(t, r, "/api/v1/items/")
		assert.Equal(t, http.StatusOK, rr.Code)
	})
}
