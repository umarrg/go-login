package routes

import (
	"encoding/json"
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
	"github.com/umarrg/microservice/model"
)

var username = "umar"
var password = "secret"

func Login(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	user := model.User{}
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		json.NewEncoder(w).Encode(model.Exception{Message: err.Error()})
		return
	}
	if ValidateCredentials(user) {
		json.NewEncoder(w).Encode(model.Res{Data: user, Token: SignedTokenString(user)})
		return
	}
	json.NewEncoder(w).Encode(model.Exception{Message: "invalid credentials"})
}

func ValidateCredentials(user model.User) bool {
	if user.Username == username && user.Password == password {
		return true
	}
	return false
}

func SignedTokenString(user model.User) string {
	var st string

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"password": user.Password,
	})

	tokenString, err := token.SignedString([]byte(st))
	if err != nil {
		log.Println(err)
	}
	return tokenString
}
