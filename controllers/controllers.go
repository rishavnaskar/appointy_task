package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Sanskrita2001/insta/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://Sanskrita:sanskrita@cluster0.d0vor.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
const dbName = "instagram"
const col1Name = "user"
const col2name = "posts"

var collection1 *mongo.Collection
var collection2 *mongo.Collection

//connect with mongoDB
func init() {
	//client option
	clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success")

	collection1 = client.Database(dbName).Collection(col1Name)
	collection2 = client.Database(dbName).Collection(col2name)

	//collection instance
	fmt.Println("Collection instances are ready")
}

// insert 1 record
func insertOneMovie(movie models.User) {
	inserted, err := collection1.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 user in db with id: ", inserted.InsertedID)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	insertOneMovie(user)
	json.NewEncoder(w).Encode(user)
}
