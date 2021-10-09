package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rishavnaskar/insta/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://admin:R1i2s3hav@cluster0.nodr5.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
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

//Model helper

// insert a user
func insertOneUser(user models.User) {
	inserted, err := collection1.InsertOne(context.Background(), user)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 user in db with id: ", inserted.InsertedID)
}

// get a user
func getOneUser(userId string) {
	id, _ := primitive.ObjectIDFromHex(userId)
	filter := bson.M{"_id": id}

	result := collection1.FindOne(context.Background(), filter)
	var user bson.M
	result.Decode(&user)

	fmt.Println("User details ", result)
}

// insert a post
func insertOnePost(post models.Post) {
	inserted, err := collection2.InsertOne(context.Background(), post)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 user in db with id: ", inserted.InsertedID)
}

// get a post
func getOnePost(postId string) {
	id, _ := primitive.ObjectIDFromHex(postId)
	filter := bson.M{"_id": id}

	result := collection2.FindOne(context.Background(), filter)
	var post bson.M
	result.Decode(&post)

	fmt.Println("Post details ", result)
}

//Controllers - files

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	insertOneUser(user)
	json.NewEncoder(w).Encode(user)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var post models.Post
	_ = json.NewDecoder(r.Body).Decode(&post)
	insertOnePost(post)
	json.NewEncoder(w).Encode(post)
}

func GetAUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")

	params := mux.Vars(r)
	getOneUser(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func GetAPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")

	params := mux.Vars(r)
	getOnePost(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}
