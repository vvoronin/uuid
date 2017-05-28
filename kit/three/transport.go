package three

import (
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/kit/log"
	"net/http"
	"github.com/myesui/uuid/kit"
	"context"
	. "github.com/myesui/uuid"
)

// MakeHandler makes the UUID Service handler
func MakeHandler(service Service, logger log.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
		kithttp.ServerErrorEncoder(kit.EncodeError),
	}

	routes := kit.Routes{
		kit.Route{
			Name: "UUID",
			Method: "GET",
			Pattern: "/three/v1/uuid",
			Handler: kithttp.NewServer(
				makeUuidEndpoint(service),
				decodeUuidRequest,
				kit.Encode,
				opts...
			),
			Queries: []string{
				"namespace",
				"{namespace}",
				"name",
				"{name}",
			},
		},
	}

	return kit.AddRoutes(routes...)
}

func decodeUuidRequest(_ context.Context, request *http.Request) (interface{}, error) {
	queryParams := request.URL.Query()

	value, ok := queryParams["namespace"]
	if !ok || len(value) > 1 {
		return nil, kit.ErrInvalidArgument
	}

	namespace, err := Parse(value[0])
	if err != nil {
		return nil, kit.ErrInvalidArgument
	}

	value, ok = queryParams["name"]
	if !ok {
		return nil, kit.ErrInvalidArgument
	}

	names := make([]interface{}, len(value))
	for i := range value {
		names[i] = value[i]
	}

	return &uuidRequest{namespace, names}, nil
}
