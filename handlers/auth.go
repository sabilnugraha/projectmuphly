package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"microservice/dto/admindto"
	authdto "microservice/dto/admindto"
	dto "microservice/dto/resultdto"
	"microservice/dto/userdto"
	"microservice/models"
	"microservice/pkg/bcrypt"
	jwtToken "microservice/pkg/jwt"
	"microservice/repositories"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handlerAuth struct {
	AuthRepository repositories.AuthRepository
}

func HandlerAuth(AuthRepository repositories.AuthRepository) *handlerAuth {
	return &handlerAuth{AuthRepository}
}

func (h *handlerAuth) AddAdmin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(admindto.AdminRequest)
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

	password, err := bcrypt.HashingPassword(request.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	user := models.Admin{
		UserName:    request.UserName,
		Password:    password,
		Position:    request.Position,
		BankAccount: request.BankAccount,
		BankName:    request.BankName,
	}

	data, err := h.AuthRepository.AddAdmin(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: data}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerAuth) CreateStudentAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	nis, _ := strconv.Atoi(r.FormValue("nis"))
	request := userdto.StudentAccountRequest{
		Password: r.FormValue("password"),
		Nis:      nis,
	}

	validation := validator.New()
	err := validation.Struct(request)

	password, err := bcrypt.HashingPassword(request.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	student, err := h.AuthRepository.GetStudentId(id)

	student.Nis = request.Nis

	student.Password = password

	data, err := h.AuthRepository.CreateStudentAccount(student)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	notif := h.AuthRepository.UpdateStatusByUserId(id, "done")

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: data, Notification: notif}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerAuth) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(admindto.AdminRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	user := models.Admin{
		UserName: request.UserName,
		Password: request.Password,
	}

	// Check email
	user, err := h.AuthRepository.Login(user.UserName)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Check password
	isValid := bcrypt.CheckPasswordHash(request.Password, user.Password)
	if !isValid {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "wrong email or password"}
		json.NewEncoder(w).Encode(response)
		return
	}

	//generate token

	claims := jwt.MapClaims{}
	claims["id"] = user.Id
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix() // 2 hours expired
	claims["position"] = user.Position

	token, errGenerateToken := jwtToken.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Println(errGenerateToken)
		fmt.Println("Unauthorize")
		return
	}

	loginResponse := authdto.AdminResponse{
		UserName: user.UserName,
		Token:    token,
		Position: user.Position,
	}

	w.Header().Set("Content-Type", "application/json")
	response := dto.SuccessResult{Status: "success", Data: loginResponse}
	json.NewEncoder(w).Encode(response)

}

func (h *handlerAuth) FindLastNis(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	user, err := h.AuthRepository.FindLastNis()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: user}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerAuth) LoginStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(userdto.StudentAccountRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	user := models.User{
		Nis:      request.Nis,
		Password: request.Password,
	}

	// Check email
	user, err := h.AuthRepository.LoginStudent(user.Nis)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Check password
	isValid := bcrypt.CheckPasswordHash(request.Password, user.Password)
	if !isValid {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "wrong nis"}
		json.NewEncoder(w).Encode(response)
		return
	}

	//generate token

	claims := jwt.MapClaims{}
	claims["id"] = user.Id
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix() // 2 hours expired

	token, errGenerateToken := jwtToken.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Println(errGenerateToken)
		fmt.Println("Unauthorize")
		return
	}

	loginResponse := userdto.StudentAccountResponse{
		Nis:   user.Nis,
		Token: token,
	}

	w.Header().Set("Content-Type", "application/json")
	response := dto.SuccessResult{Status: "success", Data: loginResponse}
	json.NewEncoder(w).Encode(response)

}

func (h *handlerAuth) CheckAuth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	// Check User by Id
	user, err := h.AuthRepository.GetAdmin(userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	CheckAuthResponse := authdto.CheckAuthResponse{
		ID:       userId,
		UserName: user.UserName,
	}

	w.Header().Set("Content-Type", "application/json")
	response := dto.SuccessResult{Status: "success", Data: CheckAuthResponse}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerAuth) CheckAuthStudentId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	// Check User by Id
	user, err := h.AuthRepository.GetStudentId(userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	CheckAuthResponse := userdto.CheckAuthStudentResponse{
		ID:       userId,
		FullName: user.FullName,
	}

	w.Header().Set("Content-Type", "application/json")
	response := dto.SuccessResult{Status: "success", Data: CheckAuthResponse}
	json.NewEncoder(w).Encode(response)
}
