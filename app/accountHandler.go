package app

import (
	"encoding/json"
	"github.com/crobatair/banking/dto"
	"github.com/crobatair/banking/logger"
	"github.com/crobatair/banking/service"
	"github.com/gorilla/mux"
	"net/http"
)

type AccountHandler struct {
	service service.AccountService
}

func (h AccountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	var req dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		logger.Error("Error" + err.Error())
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		req.CustomerId = customerId
		account, appError := h.service.NewAccount(req)
		if appError != nil {
			logger.Error("Error" + appError.Message)
			writeResponse(w, appError.Code, appError.Message)
		} else {
			writeResponse(w, http.StatusCreated, account)
		}
	}
}
