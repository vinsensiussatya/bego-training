package http

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/schema"
	"github.com/rs/zerolog/log"
)

const DateLayout = "2006-01-02"

type RequestBinder struct {
	decoder *schema.Decoder
	valid   *validator.Validate
}

func NewRequestBinder() *RequestBinder {
	// todo consider go-playground/form ?
	decoder := schema.NewDecoder()
	decoder.RegisterConverter(time.Time{}, func(s string) reflect.Value {
		parsed, err := time.Parse(DateLayout, s)
		if err != nil {
			return reflect.ValueOf(nil)
		}
		return reflect.ValueOf(parsed)
	})
	decoder.IgnoreUnknownKeys(true)
	return &RequestBinder{
		decoder: decoder,
		valid:   validator.New(),
	}
}

func (b *RequestBinder) BindRequest(ctx context.Context, r *http.Request, req any) error {
	if r.Method == http.MethodGet {
		err := r.ParseForm()
		if err != nil {
			log.Ctx(ctx).Err(err).Send()
			return ErrBadRequest()
		}
		err = b.decoder.Decode(req, r.Form)
		if err != nil {
			log.Ctx(ctx).Err(err).Interface("request_form", &r.Form).Send()
			return ErrBadRequest()
		}
	} else {
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			log.Ctx(ctx).Err(err).Send()
			return ErrBadRequest()
		}
	}

	err := b.valid.Struct(req)
	if err != nil {
		log.Ctx(ctx).Err(err).Interface("req", req).Send()
		return ErrBadRequestCustomDesc(err.Error(), err.Error())
	}
	return nil
}
