package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// task complete method, update task's status to true
func taskComplete(task string) {
	fmt.Println(task)
	//MongoDB assigns ids to the data in ObjectID format. 
	//To get the ObjectID from the task id (string), 
	//we are using primitive package’s method ObjectIDFromHex
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": true}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count: ", result.ModifiedCount)
}

// TaskComplete update task route
//update request where it will update the task’s status according to task ID
func TaskComplete(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	//To get the params from the URL, we are using mux package
	params := mux.Vars(r)
	//Using mux, send task id as a string to the taskComplete function.
	taskComplete(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}