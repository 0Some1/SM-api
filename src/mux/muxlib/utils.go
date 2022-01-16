package muxlib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"main-api-store-management/app/lib"
	"net/http"
	"time"
)

func ConvertToJsonBytes(payload interface{}) ([]byte, error) {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(payload)
	return buffer.Bytes(), err
}
func HashPassword(password string) (string, error) {
	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(passwordBytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CreateToken(username, email string) (string, error) {
	//create map claims
	claims := jwt.MapClaims{}
	//set data
	claims["username"] = username
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Minute * 60 * 24 * 7).Unix()
	//creat token without signed
	T := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	//signed token
	token, err := T.SignedString([]byte(lib.SECRET_KEY))
	if err != nil {
		return "", errors.New("error in signed")
	}
	return token, nil
}

func ExtractToken(token string) (bool, map[string]interface{}, error) {
	keyFunction := func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(lib.SECRET_KEY), nil
	}
	T, err := jwt.Parse(token, keyFunction)
	if err != nil {
		return false, make(map[string]interface{}), err
	}
	claims, ok := T.Claims.(jwt.MapClaims)
	return ok, claims, nil
}

func GetParamFromPath(r *http.Request, param string) (string, error) {
	vars := mux.Vars(r)
	if id, exists := vars[param]; exists {
		return id, nil
	}
	return "", errors.New("invalid " + param + " in path")
}
