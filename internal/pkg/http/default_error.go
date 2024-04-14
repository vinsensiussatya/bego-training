package http

import (
	"net/http"
	"strconv"
)

func ErrInternal() *ErrorResponse {
	return &ErrorResponse{
		BaseResponse: BaseResponse{
			ResponseCode: strconv.Itoa(http.StatusInternalServerError),
			ResponseDesc: ResponseDesc{
				ID: "Terjadi kesalahan pada sistem kami",
				EN: "Internal error",
			},
		},
		HTTPStatus: http.StatusInternalServerError,
	}
}

func ErrUnprocessableEntity(err string) *ErrorResponse {
	return &ErrorResponse{
		BaseResponse: BaseResponse{
			ResponseCode: strconv.Itoa(http.StatusUnprocessableEntity),
			ResponseDesc: ResponseDesc{
				ID: err,
				EN: err,
			},
		},
		HTTPStatus: http.StatusUnprocessableEntity,
	}
}

func ErrUnauthorized() *ErrorResponse {
	return &ErrorResponse{
		BaseResponse: BaseResponse{
			ResponseCode: strconv.Itoa(http.StatusUnauthorized),
			ResponseDesc: ResponseDesc{
				ID: "Anda tidak diijinkan",
				EN: "You are not authorized",
			},
		},
		HTTPStatus: http.StatusUnauthorized,
	}
}

var ErrForbidden *ErrorResponse = &ErrorResponse{
	BaseResponse: BaseResponse{
		ResponseCode: strconv.Itoa(http.StatusForbidden),
		ResponseDesc: ResponseDesc{
			ID: "Anda tidak diijinkan",
			EN: "You are not authorized",
		},
	},
	HTTPStatus: http.StatusForbidden,
}

func ErrBadRequest() *ErrorResponse {
	return &ErrorResponse{
		BaseResponse: BaseResponse{
			ResponseCode: strconv.Itoa(http.StatusBadRequest),
			ResponseDesc: ResponseDesc{
				ID: "Request Anda tidak valid",
				EN: "Invalid request",
			},
		},
		HTTPStatus: http.StatusBadRequest,
	}
}

func ErrBadRequestCustomDesc(id, en string) *ErrorResponse {
	return &ErrorResponse{BaseResponse: BaseResponse{
		ResponseCode: strconv.Itoa(http.StatusBadRequest),
		ResponseDesc: ResponseDesc{
			ID: "Request Anda tidak valid: " + id,
			EN: "Invalid request: " + en,
		},
	},
		HTTPStatus: http.StatusBadRequest,
	}
}

func ErrNotFound() *ErrorResponse {
	return &ErrorResponse{
		BaseResponse: BaseResponse{
			ResponseCode: strconv.Itoa(http.StatusNotFound),
			ResponseDesc: ResponseDesc{
				ID: "Tidak ditemukan",
				EN: "Not found",
			},
		},
		HTTPStatus: http.StatusNotFound,
	}
}
