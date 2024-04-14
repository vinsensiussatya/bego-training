package http

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
)

type HandlerOption func(*Handler)

type ResponseBody struct {
	File *FileResponse
	JSON JSONResponse
}

type FileResponse struct {
	*os.File
	FileName string
}

type JSONResponse struct {
	Data       any                 `json:"data,omitempty"`
	Pagination *PaginationResponse `json:"pagination,omitempty"`
}

func NewJSONResponse(data any, pagination *PaginationResponse) ResponseBody {
	return ResponseBody{
		JSON: JSONResponse{
			Data:       data,
			Pagination: pagination,
		}}
}

func NewFileResponse(file *os.File, fileName string) ResponseBody {
	return ResponseBody{File: &FileResponse{
		File:     file,
		FileName: fileName,
	}}
}

type Handler struct {
	// H is handler, with return any as data object, *string for token next page, error for error type
	H func(w http.ResponseWriter, r *http.Request) (ResponseBody, error)
	CustomWriter
	ServiceName string
}

func NewHTTPHandler(opts ...HandlerOption) func(handler func(w http.ResponseWriter, r *http.Request) (ResponseBody, error)) Handler {
	return func(handler func(w http.ResponseWriter, r *http.Request) (ResponseBody, error)) Handler {
		h := Handler{H: handler, CustomWriter: CustomWriter{}}

		// Option paremeters values:
		for _, opt := range opts {
			opt(&h)
		}

		return h
	}
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	respBody, err := h.H(w, r)
	if err != nil {
		h.WriteError(w, err)
		return
	}

	if respBody.File != nil {
		h.writeFile(r.Context(), w, respBody.File.File, respBody.File.FileName)
		h.removeFile(r.Context(), respBody.File.File)
		return
	}

	h.WriteJSON(w, respBody.JSON.Data, respBody.JSON.Pagination)
}

// writeFile write temp file to response
func (h Handler) writeFile(ctx context.Context, w http.ResponseWriter, file *os.File, fileName string) {
	fileBytes, err := os.ReadFile(file.Name())
	if err != nil {
		log.Ctx(ctx).Err(err).Send()
		err = ErrInternal()
		h.WriteError(w, err)
		return
	}

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%q", fileName))
	w.Header().Set("Content-Type", "application/octet-stream")
	_, err = w.Write(fileBytes)
	if err != nil {
		log.Ctx(ctx).Err(err).Send()
		err = ErrInternal()
		h.WriteError(w, err)
		return
	}
}

func (h Handler) removeFile(ctx context.Context, file *os.File) {
	_ = file.Close()
	if err := os.Remove(file.Name()); err != nil {
		log.Ctx(ctx).Err(err).Send()
	}
}
