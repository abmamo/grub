package main

import (
	"context"
	"encoding/json"
	"fmt"
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
	// get mongodb collection
	cartsCollection := db.Collection("carts")
	// create cart storing array
	var carts []Cart
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
	// get mongodb collection
	cartsCollection := db.Collection("carts")
	// create cart
	var cart Cart
	// decode request params
	_ = json.NewDecoder(request.Body).Decode(&cart)
	// serialize cart
	cart.Serialize()
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
	// get carts collection
	cartsCollection := db.Collection("carts")
	// init cart variable
	var cart Cart
	// get params from request
	var params = mux.Vars(request)
	// convert string to ObjectID
	cartId, _ := primitive.ObjectIDFromHex(params["cartId"])
	// create filter
	filter := bson.M{"_id": cartId}
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
	cartId, _ := primitive.ObjectIDFromHex(params["cartId"])
	// create filter
	filter := bson.M{"_id": cartId}
	// get update params
	_ = json.NewDecoder(request.Body).Decode(&cart)
	// prepare update model
	update := bson.D{
		{"$set", bson.D{
			{"name", cart.Name},
			{"admin", cart.Admin},
			{"coadmins", cart.CoAdmins},
			{"menulinks", cart.MenuLinks},
			{"expires", cart.Expires},
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

// DeleteCart : delete a cart in db by id
func DeleteCart(w http.ResponseWriter, request *http.Request) {
	// set header
	w.Header().Set("Content-Type", "application/json")
	// get collection
	cartsCollection := db.Collection("carts")
	// get params
	var params = mux.Vars(request)
	// string to primitive ObjectId
	cartId, err := primitive.ObjectIDFromHex(params["cartId"])
	// prepare filter
	filter := bson.M{"_id": cartId}
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

// AllOrders : get all orders in a cart by card id
func AllOrders(w http.ResponseWriter, request *http.Request) {
	// set header
	w.Header().Set("Content-Type", "application/json")
	// get carts collection
	cartsCollection := db.Collection("carts")
	// init cart variable
	var cart Cart
	// get params from request
	var params = mux.Vars(request)
	// convert string to object id
	cartId, _ := primitive.ObjectIDFromHex(params["cartId"])
	// create filter
	filter := bson.M{"_id": cartId}
	// get cart from db
	err := cartsCollection.FindOne(context.TODO(), filter).Decode(&cart)
	// check query error
	if err != nil {
		// handle error
		ErrorHandler(err, w)
		// return
		return
	}
	fmt.Println(cart.Orders)
	// write orders
	json.NewEncoder(w).Encode(cart.Orders)
}

// CreateOrder : adds an order to a given cart
func CreateOrder(w http.ResponseWriter, request *http.Request) {
	// set header
	w.Header().Set("Content-Type", "application/json")
	// get carts collection
	cartsCollection := db.Collection("carts")
	// init cart variable
	var cart Cart
	// get params from request
	var params = mux.Vars(request)
	// convert string to ObjectID
	cartId, _ := primitive.ObjectIDFromHex(params["cartId"])
	// create filter
	filter := bson.M{"_id": cartId}
	// get cart from db
	findErr := cartsCollection.FindOne(context.TODO(), filter).Decode(&cart)
	// check query error
	if findErr != nil {
		// handle error
		ErrorHandler(findErr, w)
		// return
		return
	}
	// create order
	var order Order
	// decode request params
	_ = json.NewDecoder(request.Body).Decode(&order)
	// update order
	order.ID = primitive.NewObjectID()
	// add order to card
	cart.Orders = append(cart.Orders, order)
	// update cart in db
	// prepare update model
	update := bson.D{
		{"$set", bson.D{
			{"orders", cart.Orders},
		}},
	}
	// run update query
	updateErr := cartsCollection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&cart)
	// check if query succeeded
	if updateErr != nil {
		// handle error
		ErrorHandler(updateErr, w)
		// return
		return
	}
	json.NewEncoder(w).Encode(cart)
}

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
	router.HandleFunc("/carts/get/{cartId}", GetCart).Methods("GET")
	router.HandleFunc("/carts/update/{cartId}", UpdateCart).Methods("PUT")
	router.HandleFunc("/carts/delete/{cartId}", DeleteCart).Methods("DELETE")
	router.HandleFunc("/carts/get/{cartId}/orders", AllOrders).Methods("GET")
	router.HandleFunc("/carts/update/{cartId}/orders/create", CreateOrder).Methods("POST")
	//router.HandleFunc("/carts/update/{cartId}/orders/update/{orderId}", UpdateOrder).Methods("PUT")
	//router.HandleFunc("/carts/delete/{cartId}/orders/delete/{orderId}", DeleteOrder).Methods("DELETE")
	// serve
	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, router))
}
