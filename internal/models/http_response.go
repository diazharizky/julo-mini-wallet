package models

type HTTPResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func SuccessResponse(data interface{}) HTTPResponse {
	return HTTPResponse{
		Status: "success",
		Data:   data,
	}
}
