package two

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	. "github.com/myesui/uuid"
)

type uuidRequest struct {
	SystemId `json:"type"`
}

type uuidResponse struct {
	Id UUID `json:"uuid,string"`
}

func makeUuidEndpoint(service Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(uuidRequest)
		return uuidResponse{service.UUID(req.SystemId)}, nil
	}
}
