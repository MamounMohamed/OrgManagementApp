package handler

import (
	"encoding/json"
	"net/http"
	organizationController "orgmanagementapp/pkg/controllers"
	models "orgmanagementapp/pkg/database/mongodb/models"
	"strings"
)

func CreateOrganizationHandler(w http.ResponseWriter, r *http.Request) {

	var org *models.Organization

	err := json.NewDecoder(r.Body).Decode(&org)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}
	err = organizationController.CreateOrg(org)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]string{"message": "Organization added successfully",
		"organization_id": org.OrganizationID}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
	return

}

func ReadOrganizationHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "Invalid organization ID", http.StatusBadRequest)
		return
	}
	organizationID := parts[2]
	res, err := organizationController.GetOrganizationById(organizationID)
	if res == nil {
		http.Error(w, "Organization Not Found", http.StatusInternalServerError)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&res)
	return

}

func ReadAllOrganizationsHandler(w http.ResponseWriter, r *http.Request) {
	res, err := organizationController.GetAllOrganizations()
	if res == nil {
		http.Error(w, "No Organizations Found", http.StatusInternalServerError)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&res)
	return

}

func UpdateOrganizationHandler(w http.ResponseWriter, r *http.Request) {
	// Implementation for UpdateOrganization endpoint
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "Invalid organization ID", http.StatusBadRequest)
		return
	}
	organizationID := parts[2]
	var data map[string]string
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}
	name := data["name"]
	description := data["description"]
	res, err := organizationController.UpdateOrganization(organizationID, name, description)
	if res == nil {
		http.Error(w, "Couldn't retrive or update organization", http.StatusInternalServerError)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&res)
	return

}

func DeleteOrganizationHandler(w http.ResponseWriter, r *http.Request) {
	// Implementation for DeleteOrganization endpoint
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "Invalid organization ID", http.StatusBadRequest)
		return
	}
	organizationID := parts[2]
	err := organizationController.DeleteOrganization(organizationID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response := map[string]string{"message": "Organization deleted successfully"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
	return
}

func InviteUserToOrganizationHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "Invalid organization ID", http.StatusBadRequest)
		return
	}
	organizationID := parts[2]
	var data map[string]string
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}
	email := data["email"]
	err = organizationController.InviteUserToOrganization(organizationID, email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response := map[string]string{"message": "user invited successfully"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
	return

}
