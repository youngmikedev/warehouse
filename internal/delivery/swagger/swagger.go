package swagger

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/imranzahaev/warehouse/internal/config"
	"github.com/imranzahaev/warehouse/internal/delivery/swagger/gen/models"
	"github.com/imranzahaev/warehouse/internal/delivery/swagger/gen/restapi"
	"github.com/imranzahaev/warehouse/internal/delivery/swagger/gen/restapi/operations"
	"github.com/imranzahaev/warehouse/internal/delivery/swagger/gen/restapi/operations/auth"
	"github.com/imranzahaev/warehouse/internal/delivery/swagger/gen/restapi/operations/users"
	"github.com/imranzahaev/warehouse/internal/domain"
	"github.com/imranzahaev/warehouse/internal/service"
	"github.com/rs/cors"
	"github.com/rs/zerolog"
)

func NewServer(services *service.Services, cfg *config.Config, logger *zerolog.Logger) (*restapi.Server, error) {
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		return nil, err
	}

	api := operations.NewDeliveryAPI(swaggerSpec)
	server := restapi.NewServer(api)

	server.GracefulTimeout = 10 * time.Second
	server.Host = cfg.HTTP.Host
	server.Port = cfg.HTTP.Port
	server.SetHandler(configureAPI(api, services, logger))
	return server, nil
}

func configureAPI(api *operations.DeliveryAPI, services *service.Services, logger *zerolog.Logger) http.Handler {
	type sessionPayload struct {
		sid int
		uid int
		// access token
		at string
	}

	api.ServeError = errors.ServeError

	api.Logger = logger.Printf

	api.UseSwaggerUI()

	api.JSONConsumer = runtime.JSONConsumer()
	api.JSONProducer = runtime.JSONProducer()

	api.UsersAuthAuth = func(token string) (interface{}, error) {
		sid, uid, err := services.User.CheckAccessToken(context.Background(), token)
		if err != nil {
			return nil, err
		}
		return &sessionPayload{sid: sid, uid: uid, at: token}, err
	}

	api.APIAuthorizer = api.Authorizer()

	api.AuthPostAuthSignInHandler = auth.PostAuthSignInHandlerFunc(func(params auth.PostAuthSignInParams) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		res, err := services.SignIn(ctx, params.Input.Email.String(), params.Input.Password.String())
		if err != nil {
			switch err {
			case domain.ErrInvalidLoginOrPassword:
				return auth.NewPostAuthSignInUnauthorized().WithPayload(domain.ErrInvalidLoginOrPassword.Error())
			case domain.ErrInternal:
				return auth.NewPostAuthSignInInternalServerError().WithPayload(domain.ErrInternal.Error())
			default:
				return auth.NewPostAuthSignInBadRequest().WithPayload(err.Error())
			}
		}

		payload := models.UserSignInResponse{
			Tokens: &models.UserTokensResponse{
				AccessToken:  res.AccessToken,
				RefreshToken: res.RefreshToken,
			},
			User: &models.UserUser{
				ID:        int64(res.User.ID),
				Email:     strfmt.Email(res.User.Email),
				Name:      res.User.Email,
				CreatedAt: strfmt.Date(res.User.CreatedAt),
			},
		}
		return auth.NewPostAuthSignInOK().WithPayload(&payload)
	})

	api.AuthPostAuthSignUpHandler = auth.PostAuthSignUpHandlerFunc(func(params auth.PostAuthSignUpParams) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		u := domain.User{
			Name:  *params.Input.Name,
			Email: params.Input.Email.String(),
		}
		if err := services.User.SignUp(ctx, u, params.Input.Password.String()); err != nil {
			switch err {
			case domain.ErrInternal:
				return auth.NewPostAuthSignUpInternalServerError().WithPayload(domain.ErrInternal.Error())
			default:
				return auth.NewPostAuthSignUpBadRequest().WithPayload(err.Error())
			}
		}

		return auth.NewPostAuthSignInOK()
	})

	api.AuthPostAuthLogOutHandler = auth.PostAuthLogOutHandlerFunc(func(params auth.PostAuthLogOutParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		sp := principal.(*sessionPayload)
		if err := services.User.LogOut(ctx, sp.at); err != nil {
			return auth.NewPostAuthLogOutInternalServerError().WithPayload(domain.ErrInternal.Error())
		}

		return auth.NewPostAuthLogOutOK()
	})

	api.AuthPostAuthRefreshTokenHandler = auth.PostAuthRefreshTokenHandlerFunc(func(params auth.PostAuthRefreshTokenParams) middleware.Responder {
		ctx := params.HTTPRequest.Context()

		at, rt, err := services.User.RefreshTokens(ctx, *params.Input.Token)
		if err != nil {
			switch err {
			// case domain.ErrSessionNotFound,
			// 	domain.ErrTokenExpired:
			// 	return auth.NewPostAuthRefreshTokenForbidden().WithPayload(err.Error())
			case domain.ErrInternal:
				return auth.NewPostAuthRefreshTokenInternalServerError().WithPayload(domain.ErrInternal.Error())
			default:
				return auth.NewPostAuthRefreshTokenForbidden().WithPayload(err.Error())
			}
		}
		return auth.NewPostAuthRefreshTokenOK().WithPayload(
			&models.UserTokensResponse{AccessToken: at, RefreshToken: rt},
		)
	})

	api.UsersGetUserHandler = users.GetUserHandlerFunc(func(params users.GetUserParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		sp := principal.(*sessionPayload)
		u, err := services.User.Get(ctx, sp.uid)
		if err != nil {
			switch err {
			case domain.ErrUserNotFound:
				return users.NewGetUserNotFound()
			default:
				return users.NewGetUserInternalServerError().WithPayload(err.Error())
			}
		}
		return users.NewGetUserOK().WithPayload(&models.UserUser{
			ID:        int64(u.ID),
			Name:      u.Name,
			Email:     strfmt.Email(u.Email),
			CreatedAt: strfmt.Date(u.CreatedAt),
		})
	})

	api.UsersPutUserHandler = users.PutUserHandlerFunc(func(params users.PutUserParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		sp := principal.(*sessionPayload)
		if err := services.User.Update(ctx, domain.User{
			ID:    sp.uid,
			Name:  params.Input.Name,
			Email: params.Input.Email.String(),
		}, params.Input.Password); err != nil {
			switch err {
			// case domain.ErrUserAlreadyExists:
			// 	return users.NewPutUserBadRequest().WithPayload(err.Error())
			case domain.ErrInternal:
				return users.NewPutUserInternalServerError().WithPayload(domain.ErrInternal.Error())
			default:
				// domain.ErrUserAlreadyExists
				return users.NewPutUserBadRequest().WithPayload(err.Error())
			}
		}

		return users.NewPutUserOK()
	})

	return setupGlobalMiddleware(api.Serve(setupMiddlewares), logger)
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler, logger *zerolog.Logger) http.Handler {
	handler = addAccessLogging(handler, logger)
	handler = cors.AllowAll().Handler(handler)

	return handler
}

func addAccessLogging(next http.Handler, logger *zerolog.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info().
			Str("method", r.Method).
			Str("url", r.URL.String()).
			Msg("Received request")
		next.ServeHTTP(w, r)
	})
}
