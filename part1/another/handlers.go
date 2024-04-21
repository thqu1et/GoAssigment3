package another

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

var jwtKey = []byte("secret_key")

var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

type Credentilas struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func Login(writer http.ResponseWriter, request *http.Request) {
	var credentials Credentilas
	err := json.NewDecoder(request.Body).Decode(&credentials)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	expectedPassword, ok := users[credentials.Username]
	if !ok || expectedPassword != credentials.Password {
		writer.WriteHeader(http.StatusUnauthorized)
	}

	experationTime := time.Now().Add(time.Minute * 5)

	claims := Claims{
		Username: credentials.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: experationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(writer, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: experationTime,
	})
}

func Home(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			writer.WriteHeader(http.StatusUnauthorized)
			return
		}
		writer.WriteHeader(http.StatusBadRequest)
	}
	tokenStr := cookie.Value
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			writer.WriteHeader(http.StatusUnauthorized)
			return
		}
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	if !tkn.Valid {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	writer.Write([]byte(fmt.Sprintf("Hello %s", claims.Username)))
}

func Refresh(writer http.ResponseWriter, request *http.Request) {

}
