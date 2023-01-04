package handlers

import (
	"encoding/json"
	dto "microservice/dto/resultdto"
	studentdto "microservice/dto/student"
	"microservice/models"
	"microservice/repositories"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type studentHandler struct {
	StudentRepository repositories.StudentRepository
}

func HandlerStudent(StudentRepository repositories.StudentRepository) *studentHandler {
	return &studentHandler{StudentRepository}
}

func (h *studentHandler) AddStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(studentdto.User)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	student := models.User{
		FullName: request.FullName,
		Nis:      request.Nis,
		Address:  request.Address,
		Phone:    request.Address,
		Category: request.Category,
		Status:   request.Status,
		Angkatan: request.Angkatan,
	}

	data, err := h.StudentRepository.AddStudent(student)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: data}
	json.NewEncoder(w).Encode(response)
}

func (h *studentHandler) GetStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	nis, _ := strconv.Atoi(mux.Vars(r)["nis"])

	student, err := h.StudentRepository.GetStudent(nis)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: student}
	json.NewEncoder(w).Encode(response)
}
