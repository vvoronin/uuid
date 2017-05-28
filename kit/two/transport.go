package two

import (
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/kit/log"
	. "github.com/myesui/uuid"
	"net/http"
	"github.com/myesui/uuid/kit"
	"context"
	"github.com/gorilla/mux"
	"strconv"
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
			Pattern: "/two/v1/uuid/{type:[0-9]}",
			Handler: kithttp.NewServer(
				makeUuidEndpoint(service),
				decodeUuidRequest,
				kit.Encode,
				opts...
			),
		},
	}

	return kit.AddRoutes(routes...)
}

func decodeUuidRequest(_ context.Context, request *http.Request) (interface{}, error) {
	vars := mux.Vars(request)
	value, ok := vars["type"]
	if !ok {
		return nil, kit.ErrMissingArgument
	}

	systemId, err := strconv.ParseUint(value, 10, 8)
	if err != nil || SystemId(systemId) <= SystemIdNone || SystemId(systemId) >= SystemIdUnknown {
		return nil, kit.ErrInvalidArgument
	}

	return uuidRequest{SystemId(systemId)}, nil
}
