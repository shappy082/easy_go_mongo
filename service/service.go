package service

import (
	"easy_go_mongo/model"
	"easy_go_mongo/repository"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

var (
	AppName    string
	AppVersion string
	AppENV     string
	AppPort    string
	TimeZone   string
)

func SetupConfig(configPath string) {
	viper.SetConfigType("json")
	viper.SetConfigName("config." + os.Getenv("ENV"))
	viper.AddConfigPath(configPath)

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	viper.Debug()
}

func SetupConfigParam() {
	AppName = viper.GetString("ProjectName")
	AppVersion = viper.GetString("Version")
	AppENV = viper.GetString("ENV")

	AppPort = viper.GetString("Listening.Port")
	TimeZone = viper.GetString("TimeZone")
}

func HealthCheck() model.Health {
	var HealthStatus model.Health
	HealthStatus.AppName = AppName
	HealthStatus.Version = AppVersion
	HealthStatus.Status = "I'm OK."
	HealthStatus.ENV = AppENV
	return HealthStatus
}

func GetTimestamp() string {
	date := time.Now()
	location, err := time.LoadLocation(TimeZone)

	if err != nil {
		log.Println("Get Timestamp Error ", err)
		return date.Format("2006-01-02T15:04:05Z")
	}
	return date.In(location).Format("2006-01-02T15:04:05Z")
}

// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
// !! WRITE YOUR CODE DOWN HERE !!
// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

func GetTodoList() model.ResponseMessageStatusData {
	filter := bson.M{}
	findOptions := options.Find()

	returnValue := model.ResponseMessageStatusData{}
	returnValue.Data = []string{}

	result, err := repository.CollectionTodoList.Find(filter, findOptions)
	log.Println(result, err)
	if err != nil {
		returnValue.Status = "error"
		return returnValue
	}

	returnValue.Status = "success"
	returnValue.Data = result
	return returnValue
}

func AddTodoList(data model.AddTaskData) model.AddTaskResponseData {
	//compute
	result := model.AddTaskResponseData{}

	err := repository.CollectionTodoList.Add(data)
	if err != nil {
		result.Status = "failed"
	} else {
		result.Status = "success"
	}

	//return
	return result
}

func GetAllTodoList() model.GetAllResponseData {
	//compute
	result := model.GetAllResponseData{}
	result.Status = "success"

	filter := bson.M{}
	findOption := options.Find()
	foundData, err := repository.CollectionTodoList.Find(filter, findOption)

	var resultData []model.AddTaskData
	marshalFoundData, _ := json.Marshal(foundData)
	_ = json.Unmarshal(marshalFoundData, &resultData)

	if err != nil {
		result.Status = "failed"
	} else {
		result.Result = resultData
	}

	//return
	return result
}
