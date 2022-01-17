package Handlers

import (
	"errors"

	"fmt"
	"gorm.io/gorm"
	"io/ioutil"
	"main-api-store-management/app/database"
	"main-api-store-management/app/lib"
	"main-api-store-management/mux/muxlib"
	"net/http"
	"strings"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		muxlib.HttpOptionsResponseHeaders(w)
		return
	}

	muxlib.InitLog(r)

	db, err := database.GetDatabase()
	if err != nil {
		fmt.Println("GetDatabase - LoginHandler error:", err)
		muxlib.HttpError500(w)
		return
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("ioutil.ReadAll  - LoginHandler error:", err)
		muxlib.HttpError400(w, "invalid request structure")
		return
	}

	userInput, err := lib.ParsUserInputFrom(reqBody)
	if err != nil {
		fmt.Println("parseUserInput - LoginHandler error:", err)
		muxlib.HttpError400(w, "invalid user input")
		return
	}

	if len(userInput.Password) == 0 || len(userInput.Username) == 0 {
		fmt.Println("empty password or username - LoginHandler error:", err)
		muxlib.HttpError400(w, "empty password or username")
		return
	}

	user, err := db.GetUserByUsername(userInput.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println(" GetUser - LoginHandler error:", err)
			muxlib.HttpError404(w, "not found username")
			return
		}
		fmt.Println("GetUser - LoginHandler error:", err)
		muxlib.HttpError500(w)
		return
	}

	check := lib.CheckPasswordHash(userInput.Password, user.Password)
	if !check {
		fmt.Println("CheckPasswordHash - LoginHandler")
		muxlib.HttpError400(w, "password is incorrect")
		return
	}

	token, err := muxlib.CreateToken(user.Username, user.Email)
	if err != nil {
		fmt.Println("CreateToken - LoginHandler error:", err)
		muxlib.HttpError500(w)
		return
	}

	response := struct {
		AccessToken string `json:"access_token"`
	}{
		AccessToken: token,
	}

	jsonBytes, err := muxlib.ConvertToJsonBytes(response)
	if err != nil {
		fmt.Println("json.Marshal - LoginHandler error:", err)
		muxlib.HttpError500(w)
		return
	}

	muxlib.HttpSuccessResponse(w, http.StatusOK, jsonBytes)

}

func SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		muxlib.HttpOptionsResponseHeaders(w)
		return
	}

	muxlib.InitLog(r)

	db, err := database.GetDatabase()
	if err != nil {
		fmt.Println("GetDatabase - SignUpHandler error:", err)
		muxlib.HttpError500(w)
		return
	}
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("ioutil.ReadAll  - SignUpHandler error:", err)
		muxlib.HttpError400(w, "invalid request structure")
		return
	}
	userInput, err := lib.ParsUserInputFrom(reqBody)
	if err != nil {
		fmt.Println("parseUserInput - SignUpHandler error:", err)
		muxlib.HttpError400(w, "page should be number or image field not found")
		return
	}

	//validate := validator.New()
	//err = validate.Struct(userInput)
	//if err != nil {
	//	fmt.Println(" validate - SignUpHandler error: ", err)
	//	muxlib.HttpError400(w, "invalid user input")
	//	return
	//}

	userInput.Password, err = lib.HashPassword(userInput.Password)
	if err != nil {
		fmt.Println(" HashPassword - SignUpHandler error: ", err)
		lib.HttpError500(w)
		return
	}

	err = db.CreateUser(userInput)
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
			fmt.Println(" Duplicate - SignUpHandler error: ", err)
			muxlib.HttpError400(w, "username or email or phone_number should be unique")
			return
		}
		fmt.Println(" CreateNewUser- SignUpHandler - ", err)
		muxlib.HttpError500(w)
		return
	}

	userInput.Password = ""
	jsonBytes, err := lib.ConvertToJsonBytes(userInput)
	if err != nil {
		fmt.Println("json.Marshal - SignUpHandler error:", err)
		muxlib.HttpError500(w)
		return
	}

	muxlib.HttpSuccessResponse(w, http.StatusCreated, jsonBytes)
}
