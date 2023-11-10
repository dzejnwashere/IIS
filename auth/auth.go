package auth

import (
	"IIS/db"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
)
import "github.com/golang-jwt/jwt/v5"

type Permission int

const (
	AdminPerm       Permission = 0
	UnprotectedPerm            = -1
)

const key = "ReplaceThisBeforeProduction!!!!"

func HasPermission(request *http.Request, perm Permission) bool {
	if perm == UnprotectedPerm {
		return true
	}
	cook, err := request.Cookie("iisauth")
	if err != nil {
		return false
	}
	token, err := jwt.Parse(cook.Value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(key), nil
	})
	if err != nil {
		return false
	}
	subject, err := token.Claims.GetSubject()
	if err != nil {
		return false
	}
	subjectI, err := strconv.ParseInt(subject, 10, 64)
	if err != nil {
		return false
	}
	return (db.GetPermissions(subjectI) & (1 << perm)) != 0
}

// Authenticate will return a valid JWT token in expected format when provided login information is correct, presentable error otherwise
func Authenticate(username string, password string) (string, error) {
	id, hash, err := db.GetUserIdPasswordHash(username)
	if err != nil {
		return "", fmt.Errorf("user `%s' not found", username)
	}
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return "", fmt.Errorf("invalid password")
	}
	signedString, err := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{"sub": id}).SignedString(key)
	if err != nil {
		return "", err
	}
	return signedString, nil
}
