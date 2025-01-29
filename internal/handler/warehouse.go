package handler

import (
	"awesomeProjectSamb/internal/warehouse"
	"encoding/json"
	"net/http"
)

func (h *Handler) IncomingGoods(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var incomingData warehouse.IncomingData

	if err := json.NewDecoder(r.Body).Decode(&incomingData); err != nil {
		asErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	err := h.warehouseService.IncomingGoods(ctx, incomingData)
	if err != nil {
		asErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	asJsonResponse(w, http.StatusOK, "success", nil)
}

func (h *Handler) OutgoingGoods(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var incomingData warehouse.OutgoingData

	if err := json.NewDecoder(r.Body).Decode(&incomingData); err != nil {
		asErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	err := h.warehouseService.OutgoingGoods(ctx, incomingData)
	if err != nil {
		asErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	asJsonResponse(w, http.StatusOK, "success", nil)
}

func (h *Handler) StockReport(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	result, err := h.warehouseService.StockReport(ctx)
	if err != nil {
		asErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	asJsonResponse(w, http.StatusOK, "success", result)
}
