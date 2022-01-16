package Handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"io/ioutil"
	"main-api-store-management/app/database"
	"main-api-store-management/app/lib"
	"main-api-store-management/app/models"
	"main-api-store-management/mux/muxlib"
	"net/http"
	"strings"
)

func CreateStore(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		muxlib.HttpOptionsResponseHeaders(w)
		return
	}

	lib.InitLog(r)

	db, err := database.GetDatabase()
	if err != nil {
		fmt.Println("GetDatabase - CreateStoreHandler error:", err)
		muxlib.HttpError500(w)
		return
	}

	user, err := muxlib.GetBearerUser(db, r.Header)
	if err != nil {
		fmt.Println("GetBearerUser - CreateStoreHandler error:", err)
		muxlib.HttpError401(w, err.Error())
		return
	}

	store := new(models.Store)
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("ioutil.ReadAll  - CreateStoreHandler error:", err)
		muxlib.HttpError400(w, "invalid request structure")
		return
	}
	err = json.Unmarshal(reqBody, store)
	if err != nil {
		fmt.Println("json.Unmarshal - CreateStoreHandler error:", err)
		muxlib.HttpError400(w, "invalid request structure")
		return
	}
	err = db.CreateStore(store, user)
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
			fmt.Println(" Duplicate - CreateStoreHandler error: ", err)
			muxlib.HttpError400(w, "username or email should be unique")
			return
		}
		fmt.Println(" CreateStoreHandler- CreateStoreHandler - ", err)
		muxlib.HttpError500(w)
		return
	}

	jsonBytes, err := muxlib.ConvertToJsonBytes(store)
	if err != nil {
		fmt.Println("json.Marshal - CreateStoreHandler error:", err)
		muxlib.HttpError500(w)
		return
	}

	muxlib.HttpSuccessResponse(w, http.StatusCreated, jsonBytes)

}

func GetAllStores(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		muxlib.HttpOptionsResponseHeaders(w)
		return
	}

	lib.InitLog(r)

	db, err := database.GetDatabase()
	if err != nil {
		fmt.Println("GetDatabase - CreateStoreHandler error:", err)
		muxlib.HttpError500(w)
		return
	}

	user, err := muxlib.GetBearerUser(db, r.Header)
	if err != nil {
		fmt.Println("GetBearerUser - CreateStoreHandler error:", err)
		muxlib.HttpError401(w, err.Error())
		return
	}

	stores, err := db.GetAllStores(user)
	if err != nil {
		fmt.Println("GetAllStores - CreateStoreHandler error:", err)
		muxlib.HttpError500(w)
		return
	}

	jsonBytes, err := muxlib.ConvertToJsonBytes(stores)
	if err != nil {
		fmt.Println("json.Marshal - CreateStoreHandler error:", err)
		muxlib.HttpError500(w)
		return
	}
	muxlib.HttpSuccessResponse(w, http.StatusOK, jsonBytes)
}

func DeleteStore(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		muxlib.HttpOptionsResponseHeaders(w)
		return
	}

	lib.InitLog(r)

	db, err := database.GetDatabase()
	if err != nil {
		fmt.Println("GetDatabase - DeleteStoreHandler error:", err)
		muxlib.HttpError500(w)
		return
	}

	_, err = muxlib.GetBearerUser(db, r.Header)
	if err != nil {
		fmt.Println("GetBearerUser - DeleteStoreHandler error:", err)
		muxlib.HttpError401(w, err.Error())
		return
	}

	storeID, err := muxlib.GetParamFromPath(r, "storeID")
	if err != nil {
		fmt.Println("GetParamFromPath - DeleteStoreHandler error:", err)
		muxlib.HttpError400(w, err.Error())
		return
	}

	err = db.DeleteStoreByID(storeID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println(" DeleteStoreByID - DeleteStoreHandler error: ", err)
			muxlib.HttpError400(w, "store not found")
			return
		}
		fmt.Println("DeleteStoreByID - DeleteStoreHandler error:", err)
		muxlib.HttpError500(w)
		return
	}

	muxlib.HttpSuccessResponse(w, http.StatusNoContent, nil)
}
