package repository

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type Database struct {
	dbName     string
	dbPath     string
	dbAuth     options.Credential
	client     *mongo.Client
	connection *mongo.Database
}

func (db *Database) ConnectDB() {
	dbName := db.dbName
	dbPath := db.dbPath
	log.Println("Connecting DB:", dbName)
	clientOptions := options.Client().ApplyURI(dbPath)

	var err error
	db.client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalln("Connection Error:", err)
	}
	err = db.client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalln("Ping Error:", err)
	}
	db.connection = db.client.Database(dbName)
	log.Println("Connected DB:", dbName)
}

func (db *Database) ConnectDBWithAuth() {
	dbName := db.dbName
	dbPath := db.dbPath
	dbAuth := db.dbAuth
	log.Println("Connecting DB:", dbName)
	clientOptions := options.Client().ApplyURI(dbPath)
	clientOptions.SetAuth(dbAuth)
	clientOptions.SetDirect(true)

	var err error
	db.client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalln("Connection Error:", err)
		return
	}
	err = db.client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalln("Ping Error:", err)
		return
	}
	db.connection = db.client.Database(dbName)
	log.Println("Connected DB:", dbName)
}

func (db Database) Disconnect() {
	dbName := db.dbName
	log.Println("Disconnecting DB:", dbName)
	if err := db.client.Disconnect(context.TODO()); err != nil {
		log.Println("Disconnecting DB Error")
		log.Println(err)
	}
	log.Println("Disconnected:", dbName)
}

type Collection struct {
	collectionName string
	dbConnection   *mongo.Database
	connection     *mongo.Collection
}

func (collection *Collection) Connect() {
	collectionName := collection.collectionName
	collectionConnected := collection.dbConnection.Collection(collectionName)
	log.Println("Connected Collection:", collectionName)

	collection.connection = collectionConnected
}

func (collection Collection) Find(filter bson.M, findOptions *options.FindOptions) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	log.Printf("Finding: %v", filter)

	cursor, err := collection.connection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		log.Println("Find Error", err)
		return nil, err
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var eachData map[string]interface{}
		// To decode into a struct, use cursor.Decode()
		err := cursor.Decode(&eachData)
		if err != nil {
			log.Println("Cursor Decode Error", err)
		}
		// do something with result...
		result = append(result, eachData)
		// log.Println(len(result))

		// To get the raw bson bytes use cursor.Current
		// raw := cursor.Current
		// log.Println(raw)
	}
	if err := cursor.Err(); err != nil {
		log.Println("Cursor Error", err)
		return nil, err
	}

	log.Println("Found")
	return result, nil
}

func (collection Collection) Add(data interface{}) error {
	log.Printf("Adding: %v", data)

	_, err := collection.connection.InsertOne(context.TODO(), data)
	if err != nil {
		log.Println("Add Error", err)
		return err
	}
	return nil
}

func (collection Collection) Update(filter bson.M, data interface{}) error {
	log.Printf("Updating: %v, by %v", filter, data)

	updateData := bson.M{"$set": data}
	_, err := collection.connection.UpdateOne(context.TODO(), filter, updateData)
	if err != nil {
		log.Println("Update Error", err)
		return err
	}
	return nil
}

func (collection Collection) Delete(filter bson.M) error {
	log.Printf("Deleting: %v", filter)

	_, err := collection.connection.DeleteMany(context.TODO(), filter)
	if err != nil {
		log.Println("Delete Error", err)
		return err
	}
	return nil
}

func (collection Collection) Count(filter bson.M) (int64, error) {
	result, err := collection.connection.CountDocuments(context.TODO(), filter)
	if err != nil {
		log.Println("Count Error", err)
		return 0, err
	}
	return result, nil
}
