package routes

import (
	"diary/controller"

	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/diary", controller.GetAllDiary).Methods("GET", "OPTIONS")
	//router.HandleFunc("/api/diary", controller.CreateDiary).Methods("POST", "OPTIONS")
	//router.HandleFunc("/api/diary/{id}", controller.SaveDiary).Methods("PUT", "OPTIONS")
	//router.HandleFunc("/api/undoDiary/{id}", middleware.undoDiary).Methods("PUT", "OPTIONS")
	//router.HandleFunc("/api/deleteDiary/{id}", controller.deleteDiary).Methods("DELETE", "OPTIONS")
	return router
}
