package api

import (
	"easy_go_mongo/model"
	"easy_go_mongo/service"
	"encoding/json"
	"net/http"
)

const (
	contentTypeKey   = "Content-Type"
	contentTypeValue = "application/json"
)

func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(contentTypeKey, contentTypeValue)

	serviceData := service.HealthCheck()
	serviceData.Timestamp = service.GetTimestamp()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(serviceData)
}

// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
// !! WRITE YOUR CODE DOWN HERE !!
// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

func GetTodoListAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(contentTypeKey, contentTypeValue)

	serviceData := service.GetTodoList()
	serviceData.Timestamp = service.GetTimestamp()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(serviceData)
}

func isPOST(r *http.Request) bool {
	return r.Method == "POST"
}
func isGET(r *http.Request) bool {
	return r.Method == "GET"
}

func AddTaskAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(contentTypeKey, contentTypeValue)
	if !isPOST(r) {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	requestBody := model.AddTaskRequest{}
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	data := model.AddTaskData{
		ID:          requestBody.ID,
		Subject:     requestBody.Subject,
		Time:        requestBody.Time,
		Description: requestBody.Description,
		CreateDate:  service.GetTimestamp(),
		UpdateDate:  service.GetTimestamp(),
	}

	serviceData := service.AddTodoList(data)
	serviceData.Timestamp = service.GetTimestamp()

	statusCode := http.StatusCreated
	if serviceData.Status == "failed" {
		statusCode = http.StatusBadRequest
	}
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(serviceData)
}

func GetAllTaskAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(contentTypeKey, contentTypeValue)
	if !isGET(r) {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	serviceData := service.GetAllTodoList()
	serviceData.Timestamp = service.GetTimestamp()

	statusCode := http.StatusOK
	if serviceData.Status == "failed" {
		statusCode = http.StatusBadRequest
	}
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(serviceData)
}
