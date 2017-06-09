package kit

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Middleware struct {
	next UUIDKit
}

func (o *Middleware) embed(service UUIDKit) {
	o.next = service
}

func (o Middleware) Next() UUIDKit {
	return o.next
}

func AddMiddleware(service, next UUIDKit) UUIDKit {
	service.embed(next)
	return service
}

type Route struct {
	Name    string
	Method  string
	Pattern string
	Handler http.Handler
	Queries []string
}

type Routes []Route

func AddRoutes(routes ...Route) http.Handler {

	router := mux.NewRouter()

	for _, route := range routes {

		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.Handler).
			Queries(route.Queries...)
	}

	return router
}
