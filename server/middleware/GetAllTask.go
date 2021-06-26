package middleware

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// get all task from the DB and return it
func getAllTask() []primitive.M {
	//uses abson package to get the data from the MongoDB
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var results []primitive.M
	for cur.Next(context.Background()) {
		//bson.M is used where M is an unordered, 
		//concise representation of a BSON document. 
		//It should generally be used to Serialize BSON. 
		//All the documents returned in primitive.M type. 
		//Package primitive contains types similar to Go primitives 
		//for BSON types can do not have direct Go primitive representations.
		var result bson.M
		e := cur.Decode(&result)
		if e != nil {
			log.Fatal(e)
		}
		// fmt.Println("cur..>", cur, "result", reflect.TypeOf(result), reflect.TypeOf(result["_id"]))
		results = append(results, result)

	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.Background())
	return results
}

// GetAllTask get all the task route
func GetAllTask(w http.ResponseWriter, r *http.Request) {
	//set the header to tackle the cors issue
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	payload := getAllTask()
	json.NewEncoder(w).Encode(payload)
}