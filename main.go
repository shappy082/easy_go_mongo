package main

import (
	"easy_go_mongo/api"
	"easy_go_mongo/repository"
	"easy_go_mongo/service"
	"log"
	"net/http"
)

func main() {
	log.Println("Hello World : Easy Go Mongo")

	service.SetupConfig("./config")
	service.SetupConfigParam()

	repository.ConnectDB()
	defer repository.DisconnectDB()

	http.HandleFunc("/api/health", api.Health)
	http.HandleFunc("/api/get_todo", api.GetTodoListAPI)

	// http://localhost:80/api/addTask
	http.HandleFunc("/addTask", api.AddTaskAPI)       // POST
	http.HandleFunc("/getAllTask", api.GetAllTaskAPI) // GET

	http.ListenAndServe(":"+service.AppPort, nil)
}
