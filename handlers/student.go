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
		//BirthDate:  request.BirthDate,
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
		Nis:        request.Nis,
		BirthPlace: request.BirthPlace,
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
