package handlers

import (
	"encoding/json"
	billingdto "microservice/dto/billing"
	dto "microservice/dto/resultdto"

	"microservice/models"
	"microservice/repositories"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type billingHandler struct {
	BillingRepository repositories.BillingRepository
}

type categoryHandler struct {
	StudentRepository repositories.StudentRepository
}

func HandlerBilling(BillingRepository repositories.BillingRepository) *billingHandler {
	return &billingHandler{BillingRepository}
}

func HandlerCategory(StudentRepository repositories.StudentRepository) *categoryHandler {
	return &categoryHandler{StudentRepository}
}

func (h *billingHandler) AddBilling(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(billingdto.BillingRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	validate := validator.New()
	err := validate.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	var spp int
	var ops int
	var dorm int
	if request.Category == "A" {
		ops = 50000
		spp = 200000
		dorm = 400000
	} else if request.Category == "B" {
		ops = 50000
		spp = 100000
		dorm = 400000
	} else if request.Category == "C" {
		ops = 50000
		spp = 0
		dorm = 400000
	} else if request.Category == "D" {
		ops = 0
		spp = 0
		dorm = 0
	}

	requestForm := models.BillingMonthly{
		UserID:    request.IdUser,
		Monthly:   spp,
		Ops:       ops,
		Dormitory: dorm,
		Month:     request.Month,
	}

	validatee := validator.New()
	errr := validatee.Struct(requestForm)
	if errr != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.BillingRepository.AddBilling(requestForm)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Data: data}
	json.NewEncoder(w).Encode(response)
}
