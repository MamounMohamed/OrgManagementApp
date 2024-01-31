package controllers

import (
	"fmt"
	"orgmanagementapp/pkg/database/mongodb/models"
	userdatabase "orgmanagementapp/pkg/database/mongodb/repository"
	"regexp"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var (
	accessSecretKey  = []byte("accessSecretKey")
	refreshSecretKey = []byte("refreshSecretKey")
)

type Claims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

func ValidateEmail(email string) error {
	// Check if the email is not empty
	if strings.TrimSpace(email) == "" {
		return fmt.Errorf("Email cannot be empty")
	}

	// Use a regular expression to validate the email format
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	if matched, _ := regexp.MatchString(emailRegex, email); !matched {
		return fmt.Errorf("Invalid email format")
	}
	return nil
}

func ValidateName(name string) error {
	// Check if the name is not empty
	if strings.TrimSpace(name) == "" {
		return fmt.Errorf("Name cannot be empty")
	}
	return nil
}

func ValidatePassword(password string) error {
	// Check if the password is not empty
	if strings.TrimSpace(password) == "" {
		return fmt.Errorf("Password cannot be empty")
	}
	return nil
}

func HashPassword(password string) (string, error) {
	// Generate a hashed version of the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CheckPassword(password, hashedPassword string) bool {
	// Compare the password with its hashed version
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func validateUser(user models.User) error {

	if err := ValidateName(user.Name); err != nil {
		fmt.Println("Invalid Name:", err)
		return err
	}

	if err := ValidateEmail(user.Email); err != nil {
		fmt.Println("Invalid Email:", err)
		return err
	}

	if err := ValidatePassword(user.Password); err != nil {
		fmt.Println("Invalid Password:", err)
		return err
	}
	return nil
}

func ValidateToken(ref_token string) bool {
	_, err := userdatabase.GetUserByToken(ref_token)
	if err != nil {
		return false
	}
	return true

}

func CreateUser(newUser models.User) error {
	err := validateUser(newUser)
	if err != nil {
		return err
	}

	hashedPassword, err := HashPassword(newUser.Password)
	if err != nil {
		return err
	}

	newUser.Password = hashedPassword
	dbErr := userdatabase.CreateUser(newUser)
	return dbErr
}

func Signin(email string, password string) (error, string, string) {
	user, err := userdatabase.GetUserByEmail(email)
	if err != nil {
		return err, "", ""
	}
	if user == nil {
		return fmt.Errorf("user not found"), "", ""
	}

	if !CheckPassword(password, user.Password) {
		return fmt.Errorf("Passwords didn't match"), "", ""
	}

	acc_token, err := CreateAccessToken(user.ID.String())
	if err != nil {
		return err, "", ""
	}
	refresh_token, err2 := CreateRefreshToken(user.ID.String())
	if err2 != nil {
		return err2, "", ""
	}

	userdatabase.UpdateTokens(email, acc_token, refresh_token)
	UpdateTokens(user, acc_token, refresh_token)
	return nil, acc_token, refresh_token
}

func UpdateTokens(user *models.User, acc_token string, ref_token string) {
	user.AccsesToken = acc_token
	user.RefreshToken = ref_token
}

func CreateAccessToken(userID string) (string, error) {
	accessClaims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(), // Token expires in 1 hour
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	tokenString, err := accessToken.SignedString(accessSecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
func UpdateAccsesToken(ref_token string) (*models.User, error) {
	user, err := userdatabase.GetUserByToken(ref_token)

	if user == nil {
		return nil, fmt.Errorf("Refresh token not found")
	}

	if err != nil {
		return nil, err
	}

	new_acc_token, err := CreateAccessToken(user.ID.String())
	if err != nil {
		return nil, err
	}

	userdatabase.UpdateAccsesToken(ref_token, new_acc_token)
	user.AccsesToken = new_acc_token
	return user, nil
}
func CreateRefreshToken(userID string) (string, error) {
	refreshClaims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(), // Token expires in 7 days
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	tokenString, err := refreshToken.SignedString(refreshSecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
