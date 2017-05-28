package one

import (
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/kit/log"
	"github.com/myesui/uuid/kit"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
	"context"
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
			Pattern: "/one/v1/uuid",
			Handler: kithttp.NewServer(
				makeUuidEndpoint(service),
				kit.Decode(nil),
				kit.Encode,
				opts...
			),
		},
		kit.Route{
			Name: "BULK",
			Method: "GET",
			Pattern: "/one/v1/bulk/{amount}",
			Handler: kithttp.NewServer(
				makeBulkEndpoint(service),
				decodeBulkRequest,
				kit.Encode,
				opts...
			),
		},
	}

	return kit.AddRoutes(routes...)
}

func decodeBulkRequest(_ context.Context, request *http.Request) (interface{}, error) {
	vars := mux.Vars(request)
	value, ok := vars["amount"]
	if !ok {
		return nil, kit.ErrMissingArgument
	}

	amount, err := strconv.ParseUint(value, 10, 32)
	if err != nil {
		return nil, kit.ErrInvalidArgument
	}

	return bulkRequest{int(amount)}, nil
}