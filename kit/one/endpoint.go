package one

import (
	. "github.com/myesui/uuid"
	"github.com/go-kit/kit/endpoint"
	"context"
)

type uuidResponse struct {
	Id UUID `json:"uuid,string"`
}

func makeUuidEndpoint(service Service) endpoint.Endpoint {
	return func(_ context.Context, _ interface{}) (interface{}, error) {
		return uuidResponse{service.UUID()}, nil
	}
}

type bulkRequest struct {
	Amount int `json:"amount"`
}

type bulkResponse struct {
	Ids []UUID `json:"uuids"`
}

func makeBulkEndpoint(service Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(bulkRequest)
		return bulkResponse{service.Bulk(req.Amount)}, nil
	}
}
