package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	profiledto "party/dto/profile"
	dto "party/dto/result"
	"party/models"
	"party/repositories"
	"strconv"

	"context"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gorilla/mux"
)

type handler struct {
	ProfileRepositories repositories.ProfileRepositories
}

// var path_file = "http://localhost:5000/uploads/"

func HandlerProfile(ProfileRepositories repositories.ProfileRepositories) *handler {
	return &handler{ProfileRepositories}
}

func (h *handler) GetProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	profile, err := h.ProfileRepositories.GetProfile(id)

	profile.Image = os.Getenv("PATH_FILE") + profile.Image

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "Success", Data: profile}
	json.NewEncoder(w).Encode(response)
}

func (h *handler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dataContex := r.Context().Value("dataFile")
	filepath := ""
	if dataContex != nil {
		filepath = dataContex.(string)
	}

	request := profiledto.UpdateProfileRequest{
		FullName: r.FormValue("FullName"),
		Email:    r.FormValue("email"),
		Image:    r.FormValue("image"),
	}

	var ctx = context.Background()
	// var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	// var API_KEY = os.Getenv("API_KEY")
	// var API_SECRET = os.Getenv("API_SECRET")
	var CLOUD_NAME = "dss399smz"
	var API_KEY = "147652571825927"
	var API_SECRET = "eJoljtlnNkgIlNTsR_bM7Ji2RwA"

	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	// Upload file to Cloudinary ...
	resp, err := cld.Upload.Upload(ctx, filepath, uploader.UploadParams{Folder: "party"})

	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("error gaeys")
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	profile, err := h.ProfileRepositories.GetProfile(int(id))
	fmt.Println(profile)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if request.FullName != "" {
		profile.FullName = request.FullName
	}
	if request.Email != "" {
		profile.Email = request.Email
	}

	if filepath != "" {
		profile.Image = resp.SecureURL
	}

	data, err := h.ProfileRepositories.UpdateProfile(profile)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "Succes", Data: UpdateRespone(data)}
	json.NewEncoder(w).Encode(response)
}

func UpdateRespone(u models.Profile) profiledto.UpdateProfileResponse {
	return profiledto.UpdateProfileResponse{
		ID:       u.ID,
		FullName: u.FullName,
		Email:    u.Email,
		Image:    u.Image,
	}
}
