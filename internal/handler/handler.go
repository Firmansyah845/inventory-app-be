package handler

import (
	"awesomeProjectSamb/internal/health_checks"
	"awesomeProjectSamb/internal/warehouse"
	"database/sql"
	"encoding/json"
	"net/http"
)

type Handler struct {
	healthCheckService health_checks.Service
	warehouseService   warehouse.WarehouseService
}

func NewHandler(mysqlDb *sql.DB) *Handler {

	healthCheckService := health_checks.NewService(mysqlDb)

	warehouseService := warehouse.NewService(warehouse.NewDBStore(mysqlDb))
	return &Handler{
		healthCheckService: *healthCheckService,
		warehouseService:   warehouseService,
	}
}

type Response struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"contents"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func asJsonResponse(w http.ResponseWriter, httpStatus int, message string, data any) {
	response := Response{
		Message:    message,
		Data:       data,
		StatusCode: httpStatus,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)

	_ = json.NewEncoder(w).Encode(response)
}

func asNoFormatJsonResponse(w http.ResponseWriter, httpStatus int, _ string, data any) {
	//response := Response{
	//	Message:    message,
	//	Data:       data,
	//	StatusCode: httpStatus,
	//}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)

	_ = json.NewEncoder(w).Encode(data)
}

func asXmlResponse(w http.ResponseWriter, httpStatus int, _ string, data []byte) {
	w.Header().Set("Content-Type", "application/xml")
	w.WriteHeader(httpStatus)
	_, _ = w.Write(data)
}

func asErrorResponse(w http.ResponseWriter, httpStatus int, message string) {
	response := ErrorResponse{
		Message: message,
		Code:    httpStatus,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)

	_ = json.NewEncoder(w).Encode(response)
}

func asCustomErrorResponse(w http.ResponseWriter, httpStatus int, code int, message string) {
	response := ErrorResponse{
		Message: message,
		Code:    code,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)

	_ = json.NewEncoder(w).Encode(response)
}

func asInternalErrorResponse(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	_, _ = w.Write([]byte(err.Error()))
}
