package mongodb

import (
	"context"
	"fmt"
	models "orgmanagementapp/pkg/database/mongodb/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func NameExists(name string) bool {
	var org models.Organization
	err := GetMongoClient().Database("new-db").Collection("Organizations").FindOne(context.Background(), bson.M{"name": name}).Decode(&org)
	if err == mongo.ErrNoDocuments {
		return false
	}
	return true
}

func GetOrganizationByName(name string) (*models.Organization, error) {
	var res models.Organization
	err := GetMongoClient().Database("new-db").Collection("Organizations").FindOne(context.Background(), bson.M{"name": name}).Decode(&res)
	if err != nil {
		return nil, fmt.Errorf("couldn't retrieve organization check your given name")
	}
	return &res, nil
}
func CreateOrganization(org *models.Organization) error {
	_, err := GetMongoClient().Database("new-db").Collection("Organizations").InsertOne(context.Background(), org)
	if err != nil {
		return fmt.Errorf("Error in creating organization")
	}
	fmt.Printf("Organization '%s' created successfully.\n", org.Name)

	return nil
}

func GetOrganizationById(id string) (*models.Organization, error) {
	var res models.Organization
	err := GetMongoClient().Database("new-db").Collection("Organizations").FindOne(context.Background(), bson.M{"_id": id}).Decode(&res)
	if err != nil {
		return nil, fmt.Errorf("Couldn't get organization check id")
	}
	return &res, nil
}

func GetAllOrganizations() ([]models.Organization, error) {
	cursor, err := GetMongoClient().Database("new-db").Collection("Organizations").Find(context.Background(), bson.M{})

	if err != nil {
		return nil, fmt.Errorf("Error in retrieving organizations")
	}
	defer cursor.Close(context.Background())
	var organizations []models.Organization
	for cursor.Next(context.Background()) {
		var org models.Organization
		if err := cursor.Decode(&org); err != nil {
			return nil, fmt.Errorf("Error in retrieving organizations")
		}
		organizations = append(organizations, org)
	}

	return organizations, nil

}
func UpdateOrganization(id, name, description string) error {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"name": name, "description": description}}

	res, err := GetMongoClient().Database("new-db").Collection("Organizations").UpdateOne(context.Background(), filter, update)
	if err != nil {
		return fmt.Errorf("error in updating organization check the id")
	}
	if res.ModifiedCount == 0 {
		return fmt.Errorf("Couldn't find id")
	}

	return err
}

func DeleteOrganization(id string) error {
	res, err := GetMongoClient().Database("new-db").Collection("Organizations").DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		return fmt.Errorf("couldn't delete organization")
	}
	if res.DeletedCount == 0 {
		return fmt.Errorf("couldn't delete organization check id already exists")
	}
	return nil
}

func InviteUserToOrganization(organization_id, userEmail string) error {
	user, err := GetUserByEmail(userEmail)
	if user == nil {
		return fmt.Errorf("User Not Found")
	}
	if err != nil {
		return err
	}
	member := models.OrganizationMember{
		Email:       userEmail,
		Name:        user.Name,
		AccessLevel: "member",
	}

	org, err := GetOrganizationById(organization_id)
	if err != nil {
		return fmt.Errorf("organizationId not found")
	}

	var members []models.OrganizationMember
	members = org.OrganizationMembers
	members = append(members, member)

	filter := bson.M{"_id": organization_id}
	update := bson.M{
		"$set": bson.M{"organizationmembers": members},
	}

	_, err = GetMongoClient().Database("new-db").Collection("Organizations").UpdateOne(context.Background(), filter, update)
	if err != nil {
		return fmt.Errorf("Error in updating organization")
	}
	return nil
}
