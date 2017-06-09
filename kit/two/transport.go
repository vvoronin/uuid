package two

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-kit/kit/log"
	khttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
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
			Pattern: "/two/v1/uuid/{type:[0-9]}",
			Handler: khttp.NewServer(
				makeUuidEndpoint(service),
				decodeUuidRequest,
				kit.Encode,
				opts...,
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
