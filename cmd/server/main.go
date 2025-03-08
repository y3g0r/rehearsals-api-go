// This is an example of implementing the Pet Store from the OpenAPI documentation
// found at:
// https://github.com/OAI/OpenAPI-Specification/blob/master/examples/v3.0/petstore.yaml

package main

import (
	"context"
	"fmt"
	"github.com/go-chi/cors"
	"github.com/y3g0r/rehearsals-api-go/api"
	"github.com/y3g0r/rehearsals-api-go/internal/config"
	"github.com/y3g0r/rehearsals-api-go/internal/domain"
	"github.com/y3g0r/rehearsals-api-go/internal/repository"
	"github.com/y3g0r/rehearsals-api-go/internal/service"
	"os"

	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	cfg := config.LoadConfig()

	swagger, err := api.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}

	// Clear out the servers array in the swagger spec, that skips validating
	// that server names match. We don't know how this thing will be run.
	swagger.Servers = nil

	// dependency injection
	userRepo := repository.NewInMemoryUserRepository()
	cryptoSvc := service.NewCryptoService([]byte(cfg.SecretKey))
	userService := service.NewUserService(userRepo, cryptoSvc)

	// create dummy user
	password, _ := cryptoSvc.HashPassword("123456789")
	dummyUser := domain.NewSimpleUserWithEmail(
		"dfcade99-9451-4f9a-8f1d-44015d4c4ff7",
		password,
		"dummy@example.com",
	)

	_ = userRepo.CreateUser(context.Background(), dummyUser)

	// Create an instance of our handler which satisfies the generated interface
	fastApi := api.NewFastAPI(userService, cryptoSvc)
	fastApiStrictHandler := api.NewStrictHandler(fastApi, nil)

	// This is how you set up a basic chi router
	r := chi.NewRouter()

	// Enable CORS
	// TODO: make CORS configurable with environment variables
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"}, // Use * to allow all origins
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	// Use our validation middleware to check all requests against the
	// OpenAPI schema.
	//r.Use(middleware.OapiRequestValidator(swagger))

	// We now register our petStore above as the handler for the interface
	api.HandlerFromMux(fastApiStrictHandler, r)

	s := &http.Server{
		Handler: r,
		Addr:    cfg.ServerAddress,
	}

	// And we serve HTTP until the world ends.
	log.Fatal(s.ListenAndServe())
}
