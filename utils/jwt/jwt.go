package jwt

import (
	"fmt"

	gojwt "github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
)

// Usage
// func main() {
// 	// Secret key for signing and verifying tokens

// 	goTkn := NewJWT()
// 	secretKey := "keyskeys"

// 	claims := &Claims{
// 		Username: "test",
// 		Role:     "test",
// 		RegisteredClaims: gojwt.RegisteredClaims{
// 			ID:        "idx-1",
// 			Issuer:    "merchant-a",
// 			ExpiresAt: gojwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // Set expiration time
// 		},
// 	}

// 	// Generate a token
// 	token, err := goTkn.GenerateToken(claims, secretKey)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	fmt.Println("Generated Token:", token)

// 	// Validate the token and get claims
// 	c, err := goTkn.ValidateToken(token, secretKey)
// 	if err != nil {
// 		fmt.Println("Token validation error:", err)
// 		return
// 	}
// 	fmt.Println("Token Claims:")
// 	fmt.Println("Username:", c["username"])
// 	fmt.Println("Role:", c["role"])
// }

type JWT interface {
	GenerateToken(claims gojwt.Claims, secretKey string) (string, error)
	ValidateToken(tokenString string, secretKey string) (claims map[string]interface{}, err error)
}

type jwt struct{}

func NewJWT() JWT {
	return &jwt{}
}

func (j *jwt) GenerateToken(claims gojwt.Claims, secretKey string) (string, error) {
	// Create a new token object
	token := gojwt.NewWithClaims(gojwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (j *jwt) ValidateToken(tokenString string, secretKey string) (claims map[string]interface{}, err error) {
	// Parse the token
	token, err := gojwt.Parse(tokenString, func(token *gojwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		log.Error().Msgf("failed parse token %v", err)
		return nil, err
	}

	// Extract and return claims if token is valid
	if claims, ok := token.Claims.(gojwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
