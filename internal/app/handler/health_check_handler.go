package handler

import (
	"net/http"

	"github.com/vinsensiussatya/bego-training/internal/app/service"
	myhttp "github.com/vinsensiussatya/bego-training/internal/pkg/http"
)

type IHealthCheckHandler interface {
	Readiness(w http.ResponseWriter, r *http.Request) (data myhttp.ResponseBody, err error)
	Liveness(w http.ResponseWriter, r *http.Request) (data myhttp.ResponseBody, err error)
}

type HealthCheckHandler struct {
	healthCheckService service.IHealthCheckService
}

func NewHealthCheckHandler(hcs service.IHealthCheckService) IHealthCheckHandler {
	return &HealthCheckHandler{
		healthCheckService: hcs,
	}
}

// HealthCheckReadiness godoc
// @Summary      Health Check Readiness
// @Description  get health from all dependencies
// @Tags         health
// @Accept       json
// @Produce      json
// @Param        Authorization header string true "Basic {token}"
// @Success      200  {object}  response.PingResponse
// @Failure      400  {string}  "bad request"
// @Failure      404  {string}  "not found"
// @Failure      500  {string}  "internal server error"
// @Router       /health/readiness [get]
func (h HealthCheckHandler) Readiness(_ http.ResponseWriter, r *http.Request) (data myhttp.ResponseBody, err error) {
	healthCheckResponse := h.healthCheckService.Ping(r.Context())
	data.JSON.Data = healthCheckResponse
	return
}

// HealthCheckLiveness godoc
// @Summary      Health Check Liveness
// @Description  check health server
// @Tags         health
// @Accept       json
// @Produce      json
// @Param        Authorization header string true "Basic {token}"
// @Success      200  {string}  "ok"
// @Failure      400  {string}  "bad request"
// @Failure      404  {string}  "not found"
// @Failure      500  {string}  "internal server error"
// @Router       /health/liveness [get]
func (h HealthCheckHandler) Liveness(_ http.ResponseWriter, r *http.Request) (data myhttp.ResponseBody, err error) {
	data.JSON.Data = "OK"
	return
}
