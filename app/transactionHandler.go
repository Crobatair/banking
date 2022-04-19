package app

import (
	"encoding/json"
	"github.com/crobatair/banking/dto"
	"github.com/crobatair/banking/service"
	"net/http"
)

type TransactionHandler struct {
	service service.TransactionService
}

func (h TransactionHandler) MakeTransaction(w http.ResponseWriter, r *http.Request) {
	var req dto.TransactionRequestBody
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, "Error decoding request body"+err.Error())
		return
	}

	res, appErr := h.service.MakeTransaction(&req)
	if appErr != nil {
		writeResponse(w, appErr.Code, appErr.Message)
		return
	}

	writeResponse(w, http.StatusOK, res)

}
