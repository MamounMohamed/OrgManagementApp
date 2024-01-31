package models

type OrganizationMember struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	AccessLevel string `json:"access_level"`
}

type Organization struct {
	OrganizationID      string               `json:"organization_id,omitempty" bson:"_id,omitempty"`
	Name                string               `json:"name"`
	Description         string               `json:"description"`
	OrganizationMembers []OrganizationMember `json:"organization_members"`
}
