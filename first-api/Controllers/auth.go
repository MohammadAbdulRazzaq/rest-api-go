package Controllers

import (
	"net/http"

	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// generateToken generates a JWT token
func generateToken() string {
	// Define the expiration time for the token
	expirationTime := time.Now().Add(24 * time.Hour) // Token expires in 24 hours

	// Create the JWT claims, which include the expiration time
	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
	}

	// Generate the token using the claims and your JWT secret key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("AbdulrazzaQ"))
	if err != nil {
		// Handle the error if token generation fails
		// For example, you can return an empty string or an error message
		return ""
	}

	return signedToken
}

// SignIn ... User sign-in
func SignIn(c *gin.Context) {
	// Perform user authentication logic here
	// You can validate the user's credentials, generate a token, etc.

	// Assuming authentication is successful, generate a JWT token
	token := generateToken()

	// Return the token as the response
	c.JSON(http.StatusOK, gin.H{"token": token})
}
