// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"log"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/rs/cors"
	"github.com/youngmikedev/warehouse/internal/delivery/swagger/gen/restapi/operations"
	"github.com/youngmikedev/warehouse/internal/delivery/swagger/gen/restapi/operations/auth"
	"github.com/youngmikedev/warehouse/internal/delivery/swagger/gen/restapi/operations/users"
)

//go:generate swagger generate server --target ../../gen --name Delivery --spec ../../swagger.yaml --principal interface{} --exclude-main

func configureFlags(api *operations.DeliveryAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.DeliveryAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "Authorization" header is set
	if api.UsersAuthAuth == nil {
		api.UsersAuthAuth = func(token string) (interface{}, error) {
			return nil, errors.NotImplemented("api key auth (UsersAuth) Authorization from header param [Authorization] has not yet been implemented")
		}
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()

	if api.UsersGetUserHandler == nil {
		api.UsersGetUserHandler = users.GetUserHandlerFunc(func(params users.GetUserParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation users.GetUser has not yet been implemented")
		})
	}
	if api.AuthPostAuthLogOutHandler == nil {
		api.AuthPostAuthLogOutHandler = auth.PostAuthLogOutHandlerFunc(func(params auth.PostAuthLogOutParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation auth.PostAuthLogOut has not yet been implemented")
		})
	}
	if api.AuthPostAuthRefreshTokenHandler == nil {
		api.AuthPostAuthRefreshTokenHandler = auth.PostAuthRefreshTokenHandlerFunc(func(params auth.PostAuthRefreshTokenParams) middleware.Responder {
			return middleware.NotImplemented("operation auth.PostAuthRefreshToken has not yet been implemented")
		})
	}
	if api.AuthPostAuthSignInHandler == nil {
		api.AuthPostAuthSignInHandler = auth.PostAuthSignInHandlerFunc(func(params auth.PostAuthSignInParams) middleware.Responder {
			return middleware.NotImplemented("operation auth.PostAuthSignIn has not yet been implemented")
		})
	}
	if api.AuthPostAuthSignUpHandler == nil {
		api.AuthPostAuthSignUpHandler = auth.PostAuthSignUpHandlerFunc(func(params auth.PostAuthSignUpParams) middleware.Responder {
			return middleware.NotImplemented("operation auth.PostAuthSignUp has not yet been implemented")
		})
	}
	if api.UsersPutUserHandler == nil {
		api.UsersPutUserHandler = users.PutUserHandlerFunc(func(params users.PutUserParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation users.PutUser has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	handler = addLogging(handler)
	handler = cors.AllowAll().Handler(handler)

	return handler
}

func addLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("received request:", r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}
