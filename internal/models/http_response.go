package models

import "github.com/diazharizky/julo-mini-wallet/internal/enum"

type HTTPResponse struct {
	Status enum.HTTPResponseStatus `json:"status"`
	Data   interface{}             `json:"data,omitempty"`
}

func SuccessResponse(data interface{}) HTTPResponse {
	return HTTPResponse{
		Status: enum.HTTPResponseStatusSuccess,
		Data:   data,
	}
}

func FailedResponse(data interface{}) HTTPResponse {
	return HTTPResponse{
		Status: enum.HTTPResponseStatusFailed,
		Data:   data,
	}
}

func FatalResponse() HTTPResponse {
	return FailedResponse(map[string]interface{}{
		"message": "Internal server error",
	})
}
