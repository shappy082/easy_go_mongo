package model

type Health struct {
	AppName   string `json:"app_name"`
	Version   string `json:"version"`
	Status    string `json:"status"`
	ENV       string `json:"env"`
	Timestamp string `json:"timestamp"`
}

type ResponseMessageStatusData struct {
	Status    string      `json:"response_status"`
	Message   string      `json:"response_message"`
	Data      interface{} `json:"response_data"`
	Timestamp string      `json:"response_timestamp"`
}

type AddTaskResponseData struct {
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}

type AddTaskRequest struct {
	ID          string `json:"id"`
	Subject     string `json:"subject"`
	Time        string `json:"time"`
	Description string `json:"description"`
}

type AddTaskData struct {
	ID          string `json:"id"`
	Subject     string `json:"subject"`
	Time        string `json:"time"`
	Description string `json:"description"`
	CreateDate  string `json:"create_date" bson:"create_date"`
	UpdateDate  string `json:"update_date" bson:"update_date"`
}

type GetAllResponseData struct {
	Status    string        `json:"status"`
	Result    []AddTaskData `json:"result"`
	Timestamp string        `json:"timestamp"`
}
