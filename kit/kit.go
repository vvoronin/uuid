package kit

import (
	"errors"
	"net/http"
	"encoding/json"
	kithttp "github.com/go-kit/kit/transport/http"
	"context"
)

var (
	// ErrInvalidArgument is returned when there is something wrong with expected values
	ErrInvalidArgument = errors.New("invalid argument")

	// ErrInvalidArgument is returned when there is something wrong with expected values
	ErrMissingArgument = errors.New("missing argument")

	// ErrUnknown is returned when there is some unknown error
	ErrUnknown = errors.New("unknown error")
)

type Kit interface {
	embed(service Kit)
	String() string
}

func Decode(objectFunc func() interface{}) kithttp.DecodeRequestFunc {
	return func(ctx context.Context, request *http.Request) (interface{}, error) {
		if (objectFunc == nil) {
			return nil, nil
		}
		object := objectFunc()
		if err := json.NewDecoder(request.Body).Decode(object); err != nil {
			return nil, err
		}
		return object, nil
	}
}

type serviceError interface {
	error() error
}

func Encode(ctx context.Context, response http.ResponseWriter, object interface{}) error {
	if e, ok := response.(serviceError); ok && e.error() != nil {
		EncodeError(ctx, e.error(), response)
		return nil
	}

	return json.NewEncoder(response).Encode(object)
}

// encode errors
func EncodeError(_ context.Context, err error, response http.ResponseWriter) {
	response.Header().Set("Content-Type", "application/json; charset=utf-8")

	switch err {
	case ErrUnknown:
		response.WriteHeader(http.StatusNotFound)
	case ErrInvalidArgument, ErrMissingArgument:
		response.WriteHeader(http.StatusBadRequest)
	default:
		response.WriteHeader(http.StatusInternalServerError)
	}

	json.NewEncoder(response).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func Make(service Kit) Kit {
	service.embed(service)
	return service
}
