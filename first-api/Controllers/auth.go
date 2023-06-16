package Controllers

import (
	"first-api/models"
	"first-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignIn(c *gin.Context) {
	var loginData models.LoginData
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Perform user authentication logic here
	if !isValidCredentials(loginData.Username, loginData.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Assuming authentication is successful, generate a JWT token
	tokenString, err := utils.GenerateToken(loginData.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Attach the token to the response header
	c.Header("Authorization", "Bearer "+tokenString)

	// Return a success response
	c.JSON(http.StatusOK, gin.H{"message": "Sign-in successful"})
}

// Validate the user's credentials
// Replace this with your own authentication logic
func isValidCredentials(username, password string) bool {
	// Example validation logic
	return username == "ozone" && password == "Password"
}
