package handlers

import (
	"encoding/json"
	"microservice/dto/notificationdto"
	dto "microservice/dto/resultdto"
	"microservice/models"
	"microservice/repositories"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type notificationHandler struct {
	NotificationRepository repositories.NotificationRepository
}

func HandlerNotification(NotificationRepository repositories.NotificationRepository) *notificationHandler {
	return &notificationHandler{NotificationRepository}
}

func (h *notificationHandler) AddNotification(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(notificationdto.Notification)
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

	requestForm := models.Notification{
		UserID:  request.UserID,
		AdminID: request.AdminID,
		Type:    request.Type,
		Status:  "todo",
	}

	data, err := h.NotificationRepository.AddNotification(requestForm)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Data: data}
	json.NewEncoder(w).Encode(response)
}
