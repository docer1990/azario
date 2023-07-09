package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/docer1990/azario/db"
	"github.com/docer1990/azario/models"
	jwt "github.com/golang-jwt/jwt/v5"
)

func withJWTAuth(handlerFunc http.HandlerFunc, s db.Storage, applyUserCheck bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("calling JWT Auth middlware")

		tokenString := r.Header.Get("x-api-token")
		token, err := validateJWT(tokenString)
		if err != nil {
			permissionDenied(w)
			return
		}
		if !token.Valid {
			permissionDenied(w)
			return
		}

		if applyUserCheck {
			userCheck(w, r, token, s)
		}

		if err != nil {
			WriteJSON(w, http.StatusForbidden, ApiError{Error: "invalid token"})
			return
		}

		handlerFunc(w, r)
	}
}

func createJWT(account *models.Account) (string, error) {
	// Create the Claims
	claims := &jwt.MapClaims{
		"expiresAt": 15000,
		"accountID": int64(account.ID),
	}

	secret := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}

func validateJWT(tokenString string) (*jwt.Token, error) {
	secret := os.Getenv("JWT_SECRET")

	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secret), nil
	})

}

func userCheck(w http.ResponseWriter, r *http.Request, token *jwt.Token, s db.Storage) error {
	userID, err := getID(r)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, ApiError{Error: "bad request"})
		return err
	}

	account, err := s.GetAccountById(userID)
	if err != nil {
		WriteJSON(w, http.StatusNotFound, ApiError{Error: "not found"})
		return err
	}

	claims := token.Claims.(jwt.MapClaims)
	if account.ID != int64(claims["accountID"].(float64)) {
		permissionDenied(w)
		return fmt.Errorf("invalid account ID")
	}

	return nil
}
