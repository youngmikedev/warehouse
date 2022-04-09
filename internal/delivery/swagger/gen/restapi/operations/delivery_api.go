// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/imranzahaev/warehouse/internal/delivery/swagger/gen/restapi/operations/auth"
	"github.com/imranzahaev/warehouse/internal/delivery/swagger/gen/restapi/operations/products"
	"github.com/imranzahaev/warehouse/internal/delivery/swagger/gen/restapi/operations/users"
)

// NewDeliveryAPI creates a new Delivery instance
func NewDeliveryAPI(spec *loads.Document) *DeliveryAPI {
	return &DeliveryAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer: runtime.JSONConsumer(),

		JSONProducer: runtime.JSONProducer(),

		ProductsGetProductProductIDHandler: products.GetProductProductIDHandlerFunc(func(params products.GetProductProductIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation products.GetProductProductID has not yet been implemented")
		}),
		ProductsGetProductsHandler: products.GetProductsHandlerFunc(func(params products.GetProductsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation products.GetProducts has not yet been implemented")
		}),
		UsersGetUserHandler: users.GetUserHandlerFunc(func(params users.GetUserParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation users.GetUser has not yet been implemented")
		}),
		AuthPostAuthLogOutHandler: auth.PostAuthLogOutHandlerFunc(func(params auth.PostAuthLogOutParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation auth.PostAuthLogOut has not yet been implemented")
		}),
		AuthPostAuthRefreshTokenHandler: auth.PostAuthRefreshTokenHandlerFunc(func(params auth.PostAuthRefreshTokenParams) middleware.Responder {
			return middleware.NotImplemented("operation auth.PostAuthRefreshToken has not yet been implemented")
		}),
		AuthPostAuthSignInHandler: auth.PostAuthSignInHandlerFunc(func(params auth.PostAuthSignInParams) middleware.Responder {
			return middleware.NotImplemented("operation auth.PostAuthSignIn has not yet been implemented")
		}),
		AuthPostAuthSignUpHandler: auth.PostAuthSignUpHandlerFunc(func(params auth.PostAuthSignUpParams) middleware.Responder {
			return middleware.NotImplemented("operation auth.PostAuthSignUp has not yet been implemented")
		}),
		ProductsPostProductHandler: products.PostProductHandlerFunc(func(params products.PostProductParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation products.PostProduct has not yet been implemented")
		}),
		ProductsPutProductProductIDHandler: products.PutProductProductIDHandlerFunc(func(params products.PutProductProductIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation products.PutProductProductID has not yet been implemented")
		}),
		UsersPutUserHandler: users.PutUserHandlerFunc(func(params users.PutUserParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation users.PutUser has not yet been implemented")
		}),

		// Applies when the "Authorization" header is set
		UsersAuthAuth: func(token string) (interface{}, error) {
			return nil, errors.NotImplemented("api key auth (UsersAuth) Authorization from header param [Authorization] has not yet been implemented")
		},
		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*DeliveryAPI API Server for Warehouse Platform */
type DeliveryAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator

	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator

	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// UsersAuthAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key Authorization provided in the header
	UsersAuthAuth func(string) (interface{}, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// ProductsGetProductProductIDHandler sets the operation handler for the get product product ID operation
	ProductsGetProductProductIDHandler products.GetProductProductIDHandler
	// ProductsGetProductsHandler sets the operation handler for the get products operation
	ProductsGetProductsHandler products.GetProductsHandler
	// UsersGetUserHandler sets the operation handler for the get user operation
	UsersGetUserHandler users.GetUserHandler
	// AuthPostAuthLogOutHandler sets the operation handler for the post auth log out operation
	AuthPostAuthLogOutHandler auth.PostAuthLogOutHandler
	// AuthPostAuthRefreshTokenHandler sets the operation handler for the post auth refresh token operation
	AuthPostAuthRefreshTokenHandler auth.PostAuthRefreshTokenHandler
	// AuthPostAuthSignInHandler sets the operation handler for the post auth sign in operation
	AuthPostAuthSignInHandler auth.PostAuthSignInHandler
	// AuthPostAuthSignUpHandler sets the operation handler for the post auth sign up operation
	AuthPostAuthSignUpHandler auth.PostAuthSignUpHandler
	// ProductsPostProductHandler sets the operation handler for the post product operation
	ProductsPostProductHandler products.PostProductHandler
	// ProductsPutProductProductIDHandler sets the operation handler for the put product product ID operation
	ProductsPutProductProductIDHandler products.PutProductProductIDHandler
	// UsersPutUserHandler sets the operation handler for the put user operation
	UsersPutUserHandler users.PutUserHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *DeliveryAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *DeliveryAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *DeliveryAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *DeliveryAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *DeliveryAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *DeliveryAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *DeliveryAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *DeliveryAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *DeliveryAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the DeliveryAPI
func (o *DeliveryAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.UsersAuthAuth == nil {
		unregistered = append(unregistered, "AuthorizationAuth")
	}

	if o.ProductsGetProductProductIDHandler == nil {
		unregistered = append(unregistered, "products.GetProductProductIDHandler")
	}
	if o.ProductsGetProductsHandler == nil {
		unregistered = append(unregistered, "products.GetProductsHandler")
	}
	if o.UsersGetUserHandler == nil {
		unregistered = append(unregistered, "users.GetUserHandler")
	}
	if o.AuthPostAuthLogOutHandler == nil {
		unregistered = append(unregistered, "auth.PostAuthLogOutHandler")
	}
	if o.AuthPostAuthRefreshTokenHandler == nil {
		unregistered = append(unregistered, "auth.PostAuthRefreshTokenHandler")
	}
	if o.AuthPostAuthSignInHandler == nil {
		unregistered = append(unregistered, "auth.PostAuthSignInHandler")
	}
	if o.AuthPostAuthSignUpHandler == nil {
		unregistered = append(unregistered, "auth.PostAuthSignUpHandler")
	}
	if o.ProductsPostProductHandler == nil {
		unregistered = append(unregistered, "products.PostProductHandler")
	}
	if o.ProductsPutProductProductIDHandler == nil {
		unregistered = append(unregistered, "products.PutProductProductIDHandler")
	}
	if o.UsersPutUserHandler == nil {
		unregistered = append(unregistered, "users.PutUserHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *DeliveryAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *DeliveryAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	result := make(map[string]runtime.Authenticator)
	for name := range schemes {
		switch name {
		case "UsersAuth":
			scheme := schemes[name]
			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, o.UsersAuthAuth)

		}
	}
	return result
}

// Authorizer returns the registered authorizer
func (o *DeliveryAPI) Authorizer() runtime.Authorizer {
	return o.APIAuthorizer
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *DeliveryAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *DeliveryAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *DeliveryAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the delivery API
func (o *DeliveryAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *DeliveryAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/product/{productId}"] = products.NewGetProductProductID(o.context, o.ProductsGetProductProductIDHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/products"] = products.NewGetProducts(o.context, o.ProductsGetProductsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/user"] = users.NewGetUser(o.context, o.UsersGetUserHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/auth/log-out"] = auth.NewPostAuthLogOut(o.context, o.AuthPostAuthLogOutHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/auth/refresh-token"] = auth.NewPostAuthRefreshToken(o.context, o.AuthPostAuthRefreshTokenHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/auth/sign-in"] = auth.NewPostAuthSignIn(o.context, o.AuthPostAuthSignInHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/auth/sign-up"] = auth.NewPostAuthSignUp(o.context, o.AuthPostAuthSignUpHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/product"] = products.NewPostProduct(o.context, o.ProductsPostProductHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/product/{productId}"] = products.NewPutProductProductID(o.context, o.ProductsPutProductProductIDHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/user"] = users.NewPutUser(o.context, o.UsersPutUserHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *DeliveryAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *DeliveryAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *DeliveryAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *DeliveryAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *DeliveryAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}
