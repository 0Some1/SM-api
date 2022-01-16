package main

import (
	"fmt"
	"main-api-store-management/app/database"
	"main-api-store-management/app/lib"
	"main-api-store-management/mux/router"
	"net/http"
)

func main() {
	_, err := database.GetDatabase()
	if err != nil {
		fmt.Println(err)
	}
	router := router.Init()
	fmt.Println("Server started on port: " + lib.SERVER_PORT)
	err = http.ListenAndServe(":"+lib.SERVER_PORT, router)

	if err != nil {
		fmt.Println("main - problem in run server - error: ", err)
	}
}
