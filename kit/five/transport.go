package five

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/log"
	khttp "github.com/go-kit/kit/transport/http"
	. "github.com/myesui/uuid"
	"github.com/myesui/uuid/kit"
)

// MakeHandler makes the UUID Service handler
func MakeHandler(service Service, logger log.Logger) http.Handler {
	opts := []khttp.ServerOption{
		khttp.ServerErrorLogger(logger),
		khttp.ServerErrorEncoder(kit.EncodeError),
	}

	routes := kit.Routes{
		kit.Route{
			Name:    "UUID",
			Method:  "GET",
			Pattern: "/five/v1/uuid",
			Handler: khttp.NewServer(
				makeUuidEndpoint(service),
				decodeUuidRequest,
				kit.Encode,
				opts...,
			),
			Queries: []string{
				"namespace",
				"{namespace:[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}}",
				"name",
				"{name:*}",
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
