package handlers

import (
	"encoding/json"
	dto "microservice/dto/resultdto"
	studentdto "microservice/dto/student"
	"microservice/models"
	"microservice/repositories"
	"net/http"
	"strconv"
	"time"

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

	// request := new(studentdto.User)
	// if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
	// 	json.NewEncoder(w).Encode(response)
	// 	return
	// }

	nis, _ := strconv.Atoi(r.FormValue("nis"))
	nisn, _ := strconv.Atoi(r.FormValue("nisn"))
	nik, _ := strconv.Atoi(r.FormValue("nik"))
	angkgatan, _ := strconv.Atoi(r.FormValue("angkatan"))
	birthdate, _ := time.Parse("2006-01-02", r.FormValue("birthdate"))
	request := studentdto.User{
		FullName:   r.FormValue("fullname"),
		Nis:        nis,
		Address:    r.FormValue("address"),
		Phone:      r.FormValue("phone"),
		Category:   r.FormValue("category"),
		Status:     r.FormValue("status"),
		Angkatan:   angkgatan,
		Nisn:       nisn,
		Nik:        nik,
		NickName:   r.FormValue("nickname"),
		BirthPlace: r.FormValue("birthplace"),
		BirthDate:  birthdate,
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
		FullName:   request.FullName,
		NickName:   request.NickName,
		Address:    request.Address,
		Nik:        request.Nik,
		Nisn:       request.Nisn,
		BirthPlace: request.BirthPlace,
		BirthDate:  request.BirthDate,
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

func (h *studentHandler) AddPhoto(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	dataContex := r.Context().Value("dataFile")
	filepath := dataContex.(string) // add this code
	request := studentdto.AddPhoto{

		Photo: filepath,
	}

	validation := validator.New()
	err := validation.Struct(request)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// var ctx = context.Background()
	// var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	// var API_KEY = os.Getenv("API_KEY")
	// var API_SECRET = os.Getenv("API_SECRET")

	// // Add your Cloudinary credentials ...
	// cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	// // Upload file to Cloudinary ...
	// resp, err := cld.Upload.Upload(ctx, filepath, uploader.UploadParams{Folder: "waysbeans"})

	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	photo, err := h.StudentRepository.GetStudent(id)

	if filepath != "false" {
		photo.Photo = filepath
	}

	data, err := h.StudentRepository.AddPhoto(photo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: data}
	json.NewEncoder(w).Encode(response)
}

func (h *studentHandler) GetNIS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	nis, _ := strconv.Atoi(mux.Vars(r)["nis"])

	student, err := h.StudentRepository.GetNIS(nis)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Create Embed Path File on Image property here ...
	// product.Image = path_file + product.Image

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: student}
	json.NewEncoder(w).Encode(response)
}

func (h *studentHandler) AddGroupClass(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	gc, _ := strconv.Atoi(r.FormValue("groupclass"))
	startclass, _ := time.Parse("2006-01-02", r.FormValue("startclass"))
	endclass := startclass.AddDate(3, 0, 0)
	requestForm := studentdto.RequestGroupclass{
		Groupclass: gc,
		Startclass: startclass,
		Endclass:   endclass,
	}
	validation := validator.New()
	err := validation.Struct(requestForm)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	student := models.Groupclass{
		Groupclass: requestForm.Groupclass,
		Startclass: requestForm.Startclass,
		Endclass:   requestForm.Endclass,
	}

	data, err := h.StudentRepository.AddGroupClass(student)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Data: data}
	json.NewEncoder(w).Encode(response)
}

func (h *studentHandler) AddSubClass(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	gc, _ := strconv.Atoi(r.FormValue("groubclass"))
	requestForm := studentdto.RequestSubClass{
		Groupclass: gc,
		Subclass:   r.FormValue("subclass"),
	}
	validation := validator.New()
	err := validation.Struct(requestForm)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	student := models.SubClass{
		Groupclass: requestForm.Groupclass,
		Subclass:   requestForm.Subclass,
	}

	data, err := h.StudentRepository.AddSubClass(student)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Data: data}
	json.NewEncoder(w).Encode(response)
}

func (h *studentHandler) AddStudentToGroupClass(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(r.FormValue("userid"))
	class, _ := strconv.Atoi(r.FormValue("classid"))

	requestForm := studentdto.RequestStudentClass{
		UserId:  id,
		ClassId: class,
	}
	validation := validator.New()
	err := validation.Struct(requestForm)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	student := models.ClassUser{
		UserID:       uint(requestForm.UserId),
		GroupclassID: uint(requestForm.ClassId),
	}

	data, err := h.StudentRepository.AddStudentToGroupClass(student)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Data: data}
	json.NewEncoder(w).Encode(response)
}

func (h *studentHandler) AddStudentToSubClass(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(r.FormValue("userid"))
	class, _ := strconv.Atoi(r.FormValue("subclassid"))

	requestForm := studentdto.RequestStudentSubClass{
		UserId:     id,
		SubClassId: class,
	}
	validation := validator.New()
	err := validation.Struct(requestForm)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	student := models.SubclassUser{
		UserID:     uint(requestForm.UserId),
		SubClassID: uint(requestForm.SubClassId),
	}

	data, err := h.StudentRepository.AddStudentToSubClass(student)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Data: data}
	json.NewEncoder(w).Encode(response)
}
