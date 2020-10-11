package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

// ErrorResponse : error response model
type ErrorResponse struct {
	StatusCode   int    `json:"status"`
	ErrorMessage string `json:"message"`
}

// ErrorHandler : convert error to json response
func ErrorHandler(err error, w http.ResponseWriter) {
	// log
	log.Fatal(err)
	// create error response
	var response = ErrorResponse{
		ErrorMessage: err.Error(),
		StatusCode:   http.StatusInternalServerError,
	}
	// create json response
	message, _ := json.Marshal(response)
	// write response header
	w.WriteHeader(response.StatusCode)
	// write message
	w.Write(message)
}

// AllCarts : gett all carts from database
func AllCarts(w http.ResponseWriter, request *http.Request) {
	// set response ehader
	w.Header().Set("Content-Type", "application/json")
	// create cart storing array
	var carts []Cart
	// get mongodb collection
	cartsCollection := db.Collection("carts")
	// get all data from collection
	cur, err := cartsCollection.Find(context.TODO(), bson.M{})
	// check if error
	if err != nil {
		// handle error
		ErrorHandler(err, w)
		// return
		return
	}
	// defer closing db connection
	defer cur.Close(context.TODO())
	// iterate over cursor results
	for cur.Next(context.TODO()) {
		// create single cart
		var cart Cart
		// decode cart
		err := cur.Decode(&cart)
		// check if decoding failed
		if err != nil {
			log.Fatal(err)
		}
		// add decoded cart to array
		carts = append(carts, cart)
	}
	// check if cursor error
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	// encode carts array and return response
	json.NewEncoder(w).Encode(carts)
}

// CreateCart : create a cart in db
func CreateCart(w http.ResponseWriter, request *http.Request) {
	// set header
	w.Header().Set("Content-Type", "application/json")
	// create cart
	var cart Cart
	// decode request params
	_ = json.NewDecoder(request.Body).Decode(&cart)
	// serialize cart
	cart.Serialize()
	// get mongodb collection
	cartsCollection := db.Collection("carts")
	// insert cart into database
	result, err := cartsCollection.InsertOne(context.TODO(), cart)
	// check if error during insertion
	if err != nil {
		// handle error
		ErrorHandler(err, w)
		// return
		return
	}
	json.NewEncoder(w).Encode(result)
}

// GetCart : get a cart in db by id
func GetCart(w http.ResponseWriter, request *http.Request) {
	// set header
	w.Header().Set("Content-Type", "application/json")
	// init cart variable
	var cart Cart
	// get params from request
	var params = mux.Vars(request)
	// convert string to ObjectID
	id, _ := primitive.ObjectIDFromHex(params["id"])
	// create filter
	filter := bson.M{"_id": id}
	// get carts collection
	cartsCollection := db.Collection("carts")
	// get cart from db
	err := cartsCollection.FindOne(context.TODO(), filter).Decode(&cart)
	// check query error
	if err != nil {
		// handle error
		ErrorHandler(err, w)
		// return
		return
	}
	json.NewEncoder(w).Encode(cart)
}

// UpdateCart : update a cart in db by id
func UpdateCart(w http.ResponseWriter, request *http.Request) {
	// set header
	w.Header().Set("Content-Type", "application/json")
	// get carts collection
	cartsCollection := db.Collection("carts")
	// init cart variable
	var cart Cart
	// get params from request
	var params = mux.Vars(request)
	// convert string to object id
	id, _ := primitive.ObjectIDFromHex(params["id"])
	// create filter
	filter := bson.M{"_id": id}
	// get update params
	_ = json.NewDecoder(request.Body).Decode(&cart)
	// prepare update model
	update := bson.D{
		{"$set", bson.D{
			{"Name", cart.Name},
			{"Admin", cart.Admin},
			{"CoAdmins", cart.CoAdmins},
			{"MenuLinks", cart.MenuLinks},
		}},
	}
	// run update query
	err := cartsCollection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&cart)
	// check if query succeeded
	if err != nil {
		// handle error
		ErrorHandler(err, w)
		// return
		return
	}
	json.NewEncoder(w).Encode(cart)
}

func DeleteCart(w http.ResponseWriter, request *http.Request) {
	// set header
	w.Header().Set("Content-Type", "application/json")
	// get params
	var params = mux.Vars(request)
	// string to primitive ObjectId
	id, err := primitive.ObjectIDFromHex(params["id"])
	// prepare filter
	filter := bson.M{"_id": id}
	// get collection
	cartsCollection := db.Collection("carts")
	// delete query
	deleteResult, err := cartsCollection.DeleteOne(context.TODO(), filter)
	// check if query succeeded
	if err != nil {
		// handle error
		ErrorHandler(err, w)
		// return
		return
	}
	json.NewEncoder(w).Encode(deleteResult)
}

/*
func CreateCart(w http.ResponseWriter, request *http.Request) {}
func UpdateCart(w http.ResponseWriter, request *http.Request) {}
func DeleteCart(w http.ResponseWriter, request *http.Request) {}
func AllOrders(w http.ResponseWriter, request *http.Request) {}
func CreateOrder(w http.ResponseWriter, request *http.Request) {}
func UpdateOrder(w http.ResponseWriter, request *http.Request) {}
func DeleteOrder(w http.ResponseWriter, request *http.Request) {}
*/

// InitAPI : initialize cart management API
func InitAPI() {
	// get port
	port := getEnvironment("PORT", ".env")
	// check port in environment
	if port == "" {
		port = "3000"
	}
	// create router
	router := mux.NewRouter()
	// register handlers
	router.HandleFunc("/carts", AllCarts).Methods("GET")
	router.HandleFunc("/carts/create", CreateCart).Methods("POST")
	router.HandleFunc("/carts/get/{id}", GetCart).Methods("GET")
	router.HandleFunc("/carts/update/{id}", UpdateCart).Methods("PUT")
	router.HandleFunc("/carts/delete/{id}", DeleteCart).Methods("DELETE")
	//router.HandleFunc("/carts/get/{id}/orders", GetOrders).Methods("GET")
	/*
		mux.HandleFunc("/order/all", AllOrders).Methods("GET")
		mux.HandleFunc("/order/create", CreateOrder).Methods("POST")
		mux.HandleFunc("/order/update", UpdateOrder).Methods("POST")
		mux.HandleFunc("/order/delete" DeleteOrder).Methods("POST")
	*/
	// serve
	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, router))
}
