package utils

import (
    "os"
    "time"
    "math/rand"
    "github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userID uint) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "userID": userID,
        "exp":    time.Now().Add(time.Hour * 72).Unix(),
    })

    return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func ValidateJWT(tokenString string) (jwt.MapClaims, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return []byte(os.Getenv("JWT_SECRET")), nil
    })
    if err != nil {
        return nil, err
    }

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok || !token.Valid {
        return nil, err
    }

    return claims, nil
}

func GenerateRandomString(n int) string {
    const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    result := make([]byte, n)
    for i := range result {
        result[i] = letters[rand.Intn(len(letters))]
    }
    return string(result)
}
