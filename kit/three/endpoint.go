package three

import (
	. "github.com/myesui/uuid"
	"github.com/go-kit/kit/endpoint"
	"context"
)

type uuidRequest struct {
	Namespace *UUID `json:"namespace"`
	Names []interface{} `json:"names"`
}

type uuidResponse struct {
	Id UUID `json:"uuid,string"`
}

func makeUuidEndpoint(service Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(uuidRequest)
		return uuidResponse{service.UUID(req.Namespace, req.Names...)}, nil
	}
}
