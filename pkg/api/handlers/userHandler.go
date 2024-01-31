package handler

import (
	"encoding/json"
	"net/http"
	userController "orgmanagementapp/pkg/controllers"
	models "orgmanagementapp/pkg/database/mongodb/models"
)

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}
	err = userController.CreateUser(newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response := map[string]string{"message": "User signed up successfully"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
	return
}

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	var data map[string]string

	// Decode JSON request body into 'data' map
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	email := data["email"]
	password := data["password"]

	err, acc_token, ref_token := userController.Signin(email, password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userController.Signin(email, password)
	response := map[string]string{"message": "User signed in successfully", "accses_token": acc_token, "refresh_token": ref_token}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
	return

}

func RefreshTokenHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var requestBody struct {
		RefreshToken string `json:"refresh_token"`
	}

	// Decode JSON request body
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	if !userController.ValidateToken(requestBody.RefreshToken) {
		http.Error(w, "Invalid refresh token", http.StatusUnauthorized)
		return
	}

	user, err := userController.UpdateAccsesToken(requestBody.RefreshToken)

	if err != nil {
		http.Error(w, "Error creating access token", http.StatusInternalServerError)
		return
	}
	if user == nil {
		http.Error(w, "User with this token not found", http.StatusInternalServerError)
		return

	}

	response := map[string]string{"message": "succses",
		"accses_token": user.AccsesToken, "refresh_token": requestBody.RefreshToken}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
	return

}
