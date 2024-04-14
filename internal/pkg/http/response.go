package http

import "strconv"

// Meta defines meta for api format
type Meta struct {
	Version string `json:"version"`
	APIEnv  string `json:"api_env"`
}

type BaseResponse struct {
	ResponseCode string       `json:"resp_code"`
	ResponseDesc ResponseDesc `json:"resp_desc"`
	Meta         Meta         `json:"meta"`
}

type SuccessResponse struct {
	BaseResponse
	JSONResponse
}

// ErrorResponse ..
type ErrorResponse struct {
	BaseResponse
	HTTPStatus int `json:"-"`
}

func (e *ErrorResponse) Error() string {
	return e.ResponseDesc.EN
}
func NewErrorResponse(statusCode int, desID, desEN string) *ErrorResponse {
	return &ErrorResponse{
		BaseResponse: BaseResponse{
			ResponseCode: strconv.Itoa(statusCode),
			ResponseDesc: ResponseDesc{
				ID: desID,
				EN: desEN,
			},
		},
		HTTPStatus: statusCode,
	}
}

// ResponseDesc defines details data response
type ResponseDesc struct {
	ID string `json:"id"`
	EN string `json:"en"`
}
