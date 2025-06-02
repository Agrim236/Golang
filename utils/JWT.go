// filepath: d:\Golang.notesAPI\utils\jwt.go
package utils

import (
    "os"
    "time"
    "github.com/golang-jwt/jwt/v5"
)

// GenerateJWT creates a new JWT token for a user
func GenerateJWT(userID uint) (string, error) {
    // Set expiration time
    expirationTime := time.Now().Add(24 * time.Hour)
    
    // Create claims
    claims := jwt.MapClaims{
        "id":  userID,
        "exp": expirationTime.Unix(),
    }
    
    // Create token with claims
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    
    // Get JWT secret from environment
    jwtSecret := os.Getenv("JWT_SECRET")
    if jwtSecret == "" {
        jwtSecret = "your-secret-key" // Fallback secret (not recommended for production)
    }
    
    // Generate signed token
    tokenString, err := token.SignedString([]byte(jwtSecret))
    if err != nil {
        return "", err
    }
    
    return tokenString, nil
}