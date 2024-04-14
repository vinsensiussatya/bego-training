package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/vinsensiussatya/bego-training/version"

	"github.com/rs/zerolog/log"
)

type CustomWriter struct {
}

var successResp = func() SuccessResponse {
	return SuccessResponse{
		BaseResponse: BaseResponse{
			ResponseCode: fmt.Sprint(http.StatusOK),
			ResponseDesc: ResponseDesc{
				EN: "OK",
				ID: "OK",
			},
		},
	}
}

func (c *CustomWriter) WriteJSON(w http.ResponseWriter, data any, pagination *PaginationResponse) {
	voData := reflect.ValueOf(data)
	var arrayData []any
	resp := successResp()
	if voData.Kind() != reflect.Slice {
		if voData.Kind() == reflect.Array {
			resp.Data = []any{data}
		} else {
			resp.Data = data
		}
	} else {
		if voData.Len() != 0 {
			resp.Data = data
		} else {
			resp.Data = arrayData
		}
	}

	resp.Pagination = pagination
	resp.Meta = Meta{
		APIEnv:  version.Environment,
		Version: version.Version,
	}

	writeSuccessResponse(w, resp)
}

// WriteError sending error response based on err type
func (c *CustomWriter) WriteError(w http.ResponseWriter, err error) {
	var errorResponse = &ErrorResponse{}
	if !errors.As(err, &errorResponse) {
		errorResponse = ErrUnprocessableEntity(err.Error())
	}
	errorResponse.Meta = Meta{
		APIEnv:  version.Environment,
		Version: version.Version,
	}
	writeErrorResponse(w, errorResponse)
}

func writeResponse(w http.ResponseWriter, response any, contentType string, httpStatus int) {
	res, err := json.Marshal(response)
	if err != nil {
		log.Err(err).Interface("response", response).Msg("failed to marshal")

		w.WriteHeader(http.StatusInternalServerError)
		if _, err = w.Write([]byte("Failed to marshal")); err != nil {
			log.Err(err).Msg("failed to write response")
		}
		return
	}

	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(httpStatus)
	_, err = w.Write(res)
	if err != nil {
		log.Err(err).Msg("failed to write response")
	}
}

func writeSuccessResponse(w http.ResponseWriter, response SuccessResponse) {
	writeResponse(w, response, "application/json", http.StatusOK)
}

func writeErrorResponse(w http.ResponseWriter, errorResponse *ErrorResponse) {
	writeResponse(w, errorResponse, "application/json", errorResponse.HTTPStatus)
}
