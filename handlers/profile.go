package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	profiledto "waysbuck/dto/profile"
	dto "waysbuck/dto/result"
	"waysbuck/models"
	"waysbuck/repositories"

	"github.com/gorilla/mux"
)

type handlerProfile struct {
	profileRepository repositories.ProfileRepository
}

func HandlerProfile(ProfileRepository repositories.ProfileRepository) *handlerProfile {
	return &handlerProfile{ProfileRepository}
}

func (h *handlerProfile) GetProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var profile models.Profile
	profile, err := h.profileRepository.GetProfile(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseProfile(profile)}
	json.NewEncoder(w).Encode(response)
}

func convertResponseProfile(u models.Profile) profiledto.ProfileResponse{
	return profiledto.ProfileResponse{
		ID:			u.ID,
		Fullname:	u.Fullname,
		Image:		u.Image,
		UserID: 	u.UserID,
		User: 		u.User,
	}
}