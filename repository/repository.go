package repository

import (
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	DatabaseEasyGo     Database
	CollectionTodoList Collection
)

func ConnectDB() {
	ConnectEasyGoDB()
}

func ConnectEasyGoDB() {
	dbNameTarget := "Databases.EasyGoDB"
	dbName := viper.GetString(dbNameTarget + ".DBName")
	dbPath := viper.GetString(dbNameTarget+".Protocol") + "://" +
		viper.GetString(dbNameTarget+".IP") + ":" +
		viper.GetString(dbNameTarget+".Port")
	dbAuth := options.Credential{
		AuthMechanism: viper.GetString(dbNameTarget + ".AuthMechanism"),
		Username:      viper.GetString(dbNameTarget + ".User"),
		Password:      viper.GetString(dbNameTarget + ".Password"),
		AuthSource:    viper.GetString(dbNameTarget + ".AuthSource"),
	}

	DatabaseEasyGo = Database{
		dbName: dbName,
		dbPath: dbPath,
		dbAuth: dbAuth,
	}

	if dbAuth.Username != "" && dbAuth.Password != "" {
		DatabaseEasyGo.ConnectDBWithAuth()
	} else {
		DatabaseEasyGo.ConnectDB()
	}

	collectionNameTodoList := viper.GetString(dbNameTarget + ".Collections.TodoList")
	CollectionTodoList = Collection{
		dbConnection:   DatabaseEasyGo.connection,
		collectionName: collectionNameTodoList,
	}

	CollectionTodoList.Connect()
}

func DisconnectDB() {
	DatabaseEasyGo.Disconnect()
}
