package mongodb

import (
	"context"
	"errors"
	"fmt"
	models "orgmanagementapp/pkg/database/mongodb/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser(user models.User) error {
	repo := GetMongoClient().Database("new-db").Collection("Users")

	// Check if the user with the same email already exists
	existingUser, err := GetUserByEmail(user.Email)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return errors.New("User with the same email already exists")
	}

	// Insert the new user
	_, err = repo.InsertOne(context.Background(), user)
	if err != nil {
		return fmt.Errorf("Error creating user: %v", err)
	}

	fmt.Printf("User '%s' created successfully.\n", user.Name)
	return nil
}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	repo := GetMongoClient().Database("new-db").Collection("Users")

	err := repo.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, nil // User not found
	}
	if err != nil {
		return nil, fmt.Errorf("Error getting user: %v", err)
	}

	return &user, nil
}

func UpdateTokens(email, acc_token, ref_token string) error {
	filter := bson.D{{Key: "email", Value: email}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "accsestoken", Value: acc_token}, {Key: "refreshtoken", Value: ref_token}}}}
	_, err := GetMongoClient().Database("new-db").Collection("Users").UpdateOne(context.Background(), filter, update)
	if err != nil {
		return fmt.Errorf("Couldn't update token")
	}
	return nil
}

func GetTokens(email string) (*models.User, error) {
	user, err := GetUserByEmail(email)
	if err != nil || user == nil {
		return nil, fmt.Errorf("Couldn't get tokens")
	}
	return user, nil

}
func GetUserByToken(ref_token string) (*models.User, error) {
	var user models.User
	filter := bson.M{"refreshtoken": ref_token}
	err := GetMongoClient().Database("new-db").Collection("Users").FindOne(context.Background(), filter).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, fmt.Errorf("No user found with refresh token")
	}
	if err != nil {
		return nil, fmt.Errorf("Error getting user: %v", err)
	}

	return &user, nil
}

func UpdateAccsesToken(ref_token, acc_token string) error {
	filter := bson.D{{Key: "refreshtoken", Value: ref_token}}
	update := bson.D{{Key: "$set", Value: bson.M{"accsestoken": acc_token}}}
	res, err := GetMongoClient().Database("new-db").Collection("Users").UpdateOne(context.Background(), filter, update)
	if err != nil || res.ModifiedCount == 0 {
		return fmt.Errorf("Couldn't update accses token")
	}
	return nil
}
