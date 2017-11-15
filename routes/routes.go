package routes

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/grantrobertsmith/votem-vote/app"
	"github.com/grantrobertsmith/votem-vote/controllers"
	"github.com/volatiletech/abcweb/abcmiddleware"
	"github.com/volatiletech/abcweb/abcserver"
)

// NewRouter creates a new router
func NewRouter(a *app.App, middlewares []abcmiddleware.MiddlewareFunc) *chi.Mux {
	router := chi.NewRouter()

	for _, middleware := range middlewares {
		router.Use(middleware)
	}

	// The common state for each route handler
	root := controllers.Root{
		Render:  a.Render,
		Session: a.Session,
	}

	// The common state for the api handler
	apiRoot := controllers.Root{
		Session: a.Session,
	}

	// 404 route handler
	notFound := abcserver.NewNotFoundHandler(a.AssetsManifest)
	router.NotFound(notFound.Handler(a.Config.Server, a.Render))

	// 405 route handler
	methodNotAllowed := abcserver.NewMethodNotAllowedHandler()
	router.MethodNotAllowed(methodNotAllowed.Handler(a.Render))

	// error middleware handles controller errors
	errMgr := abcmiddleware.NewErrorManager(a.Render)

	errMgr.Add(abcmiddleware.NewError(controllers.ErrUnauthorized, http.StatusUnauthorized, "errors/401", nil))
	errMgr.Add(abcmiddleware.NewError(controllers.ErrForbidden, http.StatusForbidden, "errors/403", nil))

	// Make a pointer to the errMgr.Errors function so it's easier to call
	e := errMgr.Errors

	main := controllers.Main{Root: root}
	router.Get("/", e(main.Home))

	api := controllers.Main{Root: apiRoot}
	router.Post("/api/authenticate", e(api.Authenticate))
	router.Post("/api/users", e(api.Register))
	router.Post("/api/vote", e(api.Vote))

	return router
}
