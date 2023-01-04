package handlers

import (
	"encoding/json"
	"microservice/dto/journaldto"
	dto "microservice/dto/resultdto"
	"microservice/models"
	"microservice/repositories"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
)

type journalHandler struct {
	JournalRepository repositories.JournalRepository
}

func HandlerJournal(JournalRepository repositories.JournalRepository) *journalHandler {
	return &journalHandler{JournalRepository}
}

func (h *journalHandler) ReadJournals(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	journals, err := h.JournalRepository.ReadJournals()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	// Create Embed Path File on Image property here ...
	// for i, p := range products {
	// 	products[i].Image = path_file + p.Image
	// }

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: journals}
	json.NewEncoder(w).Encode(response)
}

func (h *journalHandler) InputJournal(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// get image filepath
	//dataContex := r.Context().Value("dataFile")
	//filepath := dataContex.(string) // add this code

	opsinput, _ := strconv.Atoi(r.FormValue("opsinput"))
	opsoutput, _ := strconv.Atoi(r.FormValue("opsoutput"))
	idtrans, _ := strconv.Atoi(r.FormValue("idtrans"))
	monthlyinput, _ := strconv.Atoi(r.FormValue("monthlyinput"))
	monthlyoutput, _ := strconv.Atoi(r.FormValue("monthlyoutput"))
	mahadinput, _ := strconv.Atoi(r.FormValue("mahadinput"))
	mahadoutput, _ := strconv.Atoi(r.FormValue("mahadoutput"))

	request := new(journaldto.JournalRequest)
	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// data form pattern submit to pattern entity db product
	journal := models.Journal{
		IdTransaction: idtrans,
		Desc:          r.FormValue("desc"),
		OpsInput:      opsinput,
		OpsOutput:     opsoutput,
		MonthlyInput:  monthlyinput,
		MonthlyOutput: monthlyoutput,
		MahadInput:    mahadinput,
		MahadOutput:   mahadoutput,
	}

	// if request.OpsInput > 0 {
	// 	journal.OpsInput = request.OpsInput
	// }
	// if request.OpsOutput > 0 {
	// 	journal.OpsOutput = request.OpsOutput
	// }
	// if request.MonthlyInput > 0 {
	// 	journal.MonthlyInput = request.MonthlyInput
	// }
	// if request.MonthlyOutput > 0 {
	// 	journal.MonthlyOutput = request.MonthlyOutput
	// }
	// if request.MahadInput > 0 {
	// 	journal.MahadInput = request.MahadInput
	// }
	// if request.MahadOutput > 0 {
	// 	journal.MahadOutput = request.MahadOutput
	// }

	data, err := h.JournalRepository.InputJournal(journal)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	journal, _ = h.JournalRepository.GetJournal(journal.Id)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: data}
	json.NewEncoder(w).Encode(response)
}
