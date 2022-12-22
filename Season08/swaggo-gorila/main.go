package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
	_ "swaggo-gorila/docs"
	"time"
)

type Item struct {
	Description string `json:"description"`
	ItemCode    string `json:"item_code"`
	Price       int64  `json:"price"`
	Quantity    int64  `json:"quantity"`
}

type Order struct {
	OrderID      int       `json:"order_id"`
	CustomerName string    `json:"customer_name"`
	Items        []Item    `json:"items"`
	OrderAt      time.Time `json:"order_at"`
}

var orders []Order
var oldOrderId int

// createOrders godoc
// @Summary Create order
// @Description create Order include detail item
// @Tags root
// @Accept json
// @Produce json
// @Success 200 {array} Order
// @Router /order [post]
func createOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var order Order
	//err := json.Unmarshal(r.Body., &order)
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		json.NewEncoder(w).Encode("Error")
	}
	oldOrderId++
	order.OrderID = oldOrderId
	orders = append(orders, order)
	json.NewEncoder(w).Encode(order)
}

func getOrderAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Hello")
}

func getOrderById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Hello")
}

func updateOrderById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Hello")
}
func deleteOrderById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Hello")
}

// @title Orders API
// @version 0.1

// @description this is description
// @termOfService http://iketutg.com/termp

// @Company IKG
// @contact email : iketutg@gmail.com

// @license name Apache 2.0

// @host 127.0.0.1:8000
// @BasePath /
// @scheme http
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/orders", createOrders).Methods("POST")
	router.HandleFunc("/orders/", getOrderAll).Methods("GET")
	router.HandleFunc("/orders/{orderId}", getOrderById).Methods("GET")
	router.HandleFunc("/orders/{orderId}", updateOrderById).Methods("PUT")
	router.HandleFunc("/orders/{orderId}", deleteOrderById).Methods("DELETE")
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	server := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
