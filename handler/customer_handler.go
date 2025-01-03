package handler

import (
	"bank/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type customerHandler struct {
	custSrv service.CustomerService
}

func NewCustomerHandler(custSrv service.CustomerService) customerHandler {
	return customerHandler{custSrv}
}

func (h customerHandler) GetCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := h.custSrv.GetCustomers()
	if err != nil {
		handlerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func (h customerHandler) GetCustomer(w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.Atoi(mux.Vars(r)["customerID"])

	customer, err := h.custSrv.GetCustomer(id)
	if err != nil {
		handlerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
}
