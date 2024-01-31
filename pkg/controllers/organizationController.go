package controllers

import (
	"fmt"
	"orgmanagementapp/pkg/database/mongodb/models"
	orgDatabase "orgmanagementapp/pkg/database/mongodb/repository"
	"time"
)

func generateID() string {
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	return fmt.Sprintf("%d", timestamp)
}

func CreateOrg(org *models.Organization) error {

	if orgDatabase.NameExists(org.Name) {
		return fmt.Errorf("Organization name exists")
	}

	org.OrganizationID = generateID()
	err := orgDatabase.CreateOrganization(org)
	if err != nil {
		return err
	}
	return nil
}

func GetOrganizationById(id string) (*models.Organization, error) {
	org, err := orgDatabase.GetOrganizationById(id)
	if org == nil {
		return nil, fmt.Errorf("Couldn't read organization")
	}
	if err != nil {
		return org, err
	}

	return org, nil
}
func GetAllOrganizations() ([]models.Organization, error) {
	organizations, err := orgDatabase.GetAllOrganizations()
	if organizations == nil {
		return nil, err
	}
	if err != nil {
		return organizations, err
	}
	return organizations, nil
}

func UpdateOrganization(id, name, description string) (*models.Organization, error) {
	err := orgDatabase.UpdateOrganization(id, name, description)
	if err != nil {
		return nil, err
	}
	res, err := orgDatabase.GetOrganizationById(id)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res, nil
}
func DeleteOrganization(id string) error {
	err := orgDatabase.DeleteOrganization(id)
	if err != nil {
		return err
	}
	return nil
}
func InviteUserToOrganization(id, email string) error {
	err := orgDatabase.InviteUserToOrganization(id, email)
	if err != nil {
		return err
	}
	return nil
}
