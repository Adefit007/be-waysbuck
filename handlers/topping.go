package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	dto "waysbuck/dto/result"
	toppingsdto "waysbuck/dto/topping"
	"waysbuck/models"
	"waysbuck/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type handlerTopping struct {
	ToppingRepository repositories.ToppingRepository
}

func HandlerTopping(ToppingRepository repositories.ToppingRepository) *handlerTopping {
	return &handlerTopping{ToppingRepository}
}

func (h *handlerTopping) FindToppings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	topping, err := h.ToppingRepository.FindToppings()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: topping}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTopping) GetTopping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	topping, err := h.ToppingRepository.GetTopping(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: topping}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTopping) CreateTopping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")

	request := new(toppingsdto.ToppingRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError,Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	topping := models.Topping{
		Title :		request.Title,
		Price: 		request.Price,
		Image:		request.Image,	
	}

	topping, err = h.ToppingRepository.CreateTopping(topping)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError,Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	topping, _ =h.ToppingRepository.GetTopping(topping.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK,Data: topping}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTopping) UpdateTopping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")

	request := new(toppingsdto.UpdateTopping)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest,Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	topping, err := h.ToppingRepository.GetTopping(int(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest,Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	if (request.Title) != "" {
		topping.Title = request.Title
	}

	if request.Price != 0 {
		topping.Price =request.Price
	}

	if request.Image != "" {
		topping.Image = request.Image
	}

	data, err := h.ToppingRepository.UpdateTopping(topping)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError,Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseTopping(data)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTopping) DeleteTopping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	topping, err := h.ToppingRepository.GetTopping(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest,Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	data, err := h.ToppingRepository.DeleteTopping(topping)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK,Data: data}
	json.NewEncoder(w).Encode(response)
}

func convertResponseTopping(u models.Topping) toppingsdto.ToppingRequest {
	return toppingsdto.ToppingRequest{
		Title: u.Title,
		Price: u.Price,
		Image: u.Image,
	}
}
