package router

import (
	"github.com/gorilla/mux"
	"main-api-store-management/mux/Handlers"
	"net/http"
)

func Init() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/signup", Handlers.SignUp).Methods(http.MethodPost, http.MethodOptions)

	r.HandleFunc("/login", Handlers.Login).Methods(http.MethodPost, http.MethodOptions)

	r.HandleFunc("/store", Handlers.CreateStore).Methods(http.MethodPost, http.MethodOptions)

	r.HandleFunc("/store", Handlers.GetAllStores).Methods(http.MethodGet, http.MethodOptions)

	r.HandleFunc("/store/{storeID}", Handlers.DeleteStore).Methods(http.MethodDelete, http.MethodOptions)

	r.NotFoundHandler = http.HandlerFunc(Handlers.NotFoundController)

	return r
}
