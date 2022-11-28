package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	authdto "party/dto/auth"
	dto "party/dto/result"
	"party/models"
	"party/pkg/bcrypt"
	jwtToken "party/pkg/jwt"
	"party/repositories"
	"time"

	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt/v4"
)

type handleAuth struct {
	AuthRepositories repositories.AuthRepositories
}

func HandlerAuth(AuthRepositories repositories.AuthRepositories) *handleAuth{
	return &handleAuth{AuthRepositories}
}

func (h *handleAuth) Register(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	request := new(authdto.RegisterRequest)
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
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		}

	profile := models.Profile{
		FullName: request.FullName,
		Email: request.Email,
		Password: password,
	}


	claims := jwt.MapClaims{}
	claims["id"] = profile.ID
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix() // 2 jam expired

	token, errGenerate := jwtToken.GenerateToken(&claims)
	if errGenerate != nil {
		log.Println(errGenerate)
		fmt.Println("Unauthorize")
		return
	}
	
	data, err := h.AuthRepositories.Register(profile)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message:err.Error()}
		json.NewEncoder(w).Encode(response)
		fmt.Println(data)
	}
	RegisterResponse := authdto.RegisterResponse{
		Token: token,
	}

	w.Header().Set("Content-Type", "application/json")
	response := dto.SuccessResult{Code: "Succec", Data: RegisterResponse}
	json.NewEncoder(w).Encode(response)
}

func (h *handleAuth) Login(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	request := new(authdto.LoginRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	profile := models.Profile{
		Email:    request.Email,
		Password: request.Password,
	}

	profile, err := h.AuthRepositories.Login(profile.Email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Check password
	isValid := bcrypt.CheckPasswordHash(request.Password, profile.Password)
	if !isValid {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "wrong email or password"}
		json.NewEncoder(w).Encode(response)
		return
	}

	//generate token
	claims := jwt.MapClaims{}
	claims["id"] = profile.ID
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix() // 2 hours expired

	token, errGenerateToken := jwtToken.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Println(errGenerateToken)
		fmt.Println("Unauthorize")
		return
	}

	loginResponse := authdto.LoginResponse{
		Token: token,
	}

	w.Header().Set("Content-Type", "application/json")
	response := dto.SuccessResult{Code: "Success", Data: loginResponse}
	json.NewEncoder(w).Encode(response)

}