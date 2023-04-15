package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Wai30Yan/cna-server/pkg/model"
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
)


func (m *Repository) LogIn(w http.ResponseWriter, r *http.Request) {

}

func (m *Repository) LogOut(w http.ResponseWriter, r *http.Request) {

}

func (m *Repository) SignUp(w http.ResponseWriter, r *http.Request) {
	var admin *model.Admin
	json.NewDecoder(r.Body).Decode(&admin)
	hashedPassword, err := hashPassword(admin.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	// Store the user with hashed password in a database or any other persistent storage
	
	a, err := m.DB.SignUp(admin.UserName, hashedPassword)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(&a)
}



// Function to hash password using bcrypt
func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// Function to verify password against a hashed password
func verifyPassword(password, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err
}

func generateJWTToken(username string) (string, error) {
	// Define the expiration time for the token
	expirationTime := time.Now().Add(24 * time.Hour)

	// Create a new token with the username as the claim
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":       expirationTime.Unix(),
	})

	// Sign the token with a secret key
	tokenString, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}