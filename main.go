package main

import (
"encoding/json"
"log"
"net/http"
"os"
"time"

    gocqlastra "github.com/datastax/gocql-astra"
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
    "github.com/gocql/gocql"
    "github.com/joho/godotenv"
    "github.com/shopspring/decimal"

)

var session *gocql.Session

// Product represents a product entity.
type Product struct {
ID gocql.UUID `json:"product_id"`
Name string `json:"name"`
Description string `json:"description"`
Price decimal.Decimal `json:"price"`
Quantity int `json:"quantity"`
}

func main() {
// Load environment variables and configure database
if err := godotenv.Load(); err != nil {
log.Fatal("failed to load the env file: ", err)
}

    token := os.Getenv("CASSANDRA_CLIENT_TOKEN")
    cluster, err := gocqlastra.NewClusterFromBundle("./secure-connect.zip", "token", token, 10*time.Second)
    if err != nil {
    	panic("unable to load the bundle")
    }

    session, err = gocql.NewSession(*cluster)
    if err != nil {
    	log.Fatalf("unable to create session: %v", err)
    }
    defer session.Close()

    // Initialize router
    router := chi.NewRouter()
    router.Use(middleware.Logger)

    // Define routes
    // router.Get("/products", )
    router.Get("/products/{id}", getProduct)
    router.Post("/products", createProduct)
    router.Patch("/products/{id}", updateProduct)
    router.Delete("/products/{id}", deleteProduct)

    // Start the server
    log.Println("Server started on :3000 😂")
    log.Fatal(http.ListenAndServe(":3000", router))

}

func getProduct(w http.ResponseWriter, r _http.Request) {
id := chi.URLParam(r, "id")
var product Product
if err := session.Query(`SELECT _ FROM products WHERE product_id = ?`, id).Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Quantity); err != nil {
w.WriteHeader(http.StatusNotFound)
return
}
json.NewEncoder(w).Encode(product)
}

func createProduct(w http.ResponseWriter, r \*http.Request) {
var product Product
if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
w.WriteHeader(http.StatusBadRequest)
return
}
defer r.Body.Close()

    if err := session.Query(`INSERT INTO products (product_id, name, description, price, quantity) VALUES (?, ?, ?, ?, ?)`,
    	gocql.TimeUUID(),
    	product.Name,
    	product.Description,
    	product.Price,
    	product.Quantity).Exec(); err != nil {
    	w.WriteHeader(http.StatusInternalServerError)
    	return
    }
    w.WriteHeader(http.StatusCreated)

}

func updateProduct(w http.ResponseWriter, r \*http.Request) {
id := chi.URLParam(r, "id")
var product Product
if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
w.WriteHeader(http.StatusBadRequest)
return
}
defer r.Body.Close()

    if err := session.Query(`UPDATE products SET name = ?, description = ?, price = ?, quantity = ? WHERE product_id = ?`,
    	product.Name,
    	product.Description,
    	product.Price,
    	product.Quantity,
    	id).Exec(); err != nil {
    	w.WriteHeader(http.StatusInternalServerError)
    	return
    }
    w.WriteHeader(http.StatusOK)

}

func deleteProduct(w http.ResponseWriter, r \*http.Request) {
id := chi.URLParam(r, "id")
if err := session.Query(`DELETE FROM products WHERE product_id = ?`, id).Exec(); err != nil {
w.WriteHeader(http.StatusInternalServerError)
return
}
w.WriteHeader(http.StatusOK)
}
