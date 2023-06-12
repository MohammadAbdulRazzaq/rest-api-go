package utils

import (
	"fmt"
	"strconv"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// func GenerateToken(user_id string) (string, error) {
// 	viper.SetConfigFile("../config/master.yaml")
// 	err := viper.ReadInConfig()
// 	if err != nil {
// 		return "", fmt.Errorf("failed to read config file: %w", err)
// 	}

// 	token_lifespan, err := strconv.Atoi(viper.GetString("jwt.TOKEN_HOUR_LIFESPAN"))

// 	if err != nil {
// 		return "", err
// 	}

// 	claims := jwt.MapClaims{}
// 	claims["authorized"] = true
// 	claims["user_id"] = user_id
// 	claims["exp"] = time.Now().Add(time.Hour * time.Duration(token_lifespan)).Unix()
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

// 	return token.SignedString([]byte(viper.GetString("jwt.API_SECRET")))

// }

func TokenValid(c *gin.Context) error {
	viper.SetConfigFile("../config/master.yaml")
	tokenString := ExtractToken(c)
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(viper.GetString("jwt.API_SECRET")), nil
	})
	if err != nil {
		return err
	}
	return nil
}

func ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func ExtractTokenID(c *gin.Context) (uint, error) {
	viper.SetConfigFile("../config/master.yaml")
	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(viper.GetString("jwt.API_SECRET")), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)
		if err != nil {
			return 0, err
		}
		return uint(uid), nil
	}
	return 0, nil
}

// Generate a JWT token
func GenerateToken(username string) (string, error) {
	// Create a new token object
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims (payload)
	claims := token.Claims.(jwt.MapClaims)
	claims["user"] = username // Set the username as a claim

	// Set the secret key used to sign the token
	// Replace "secret" with your own secret key
	secret := []byte("secret")
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Validate a JWT token
func ValidateToken(authorizationHeader string) (bool, error) {
	// Replace "secret" with your own secret key
	secret := []byte("secret")

	// Extract the token from the Authorization header
	tokenString := strings.TrimPrefix(authorizationHeader, "Bearer ")

	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})
	if err != nil {
		return false, err
	}

	// Check if the token is valid
	if _, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
		return false, nil
	}

	return true, nil
}
