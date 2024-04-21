package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	gocqlastra "github.com/datastax/gocql-astra"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gocql/gocql"
	"github.com/joho/godotenv"
	"github.com/yaninyzwitty/crud-eccomerce-app/repository"
	"github.com/yaninyzwitty/crud-eccomerce-app/service"
	"github.com/yaninyzwitty/crud-eccomerce-app/transport"
)

var session *gocql.Session

func main() {
	ctx := context.Background()
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

	// Initialize dependencies (products)

	productRepo := repository.NewProductRepository(ctx, session)
	productService := service.NewProductService(productRepo)
	productHander := transport.NewProductHandler(productService)

	// Initialize dependencies (categories)

	categoryRepo := repository.NewCategoryRepository(ctx, session)
	categoryService := service.NewCategoryService(categoryRepo)

	categoryHandler := transport.NewCategoryHandler(categoryService)

	// initialize dependencies (orders)

	orderRepo := repository.NewOrderRepository(ctx, session)
	orderService := service.NewOrderService(orderRepo)
	orderHandler := transport.NewOrderHandler(orderService)

	// Initialize router
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	// Define routes (products)
	router.Get("/products", productHander.GetProducts)
	router.Get("/products/{id}", productHander.GetProduct)
	router.Post("/products", productHander.CreateProduct)
	router.Patch("/products/{id}", productHander.UpdateProduct)
	router.Delete("/products/{id}", productHander.DeleteProduct)

	// Define routes (categories)

	router.Get("/categories", categoryHandler.GetCategories)
	router.Get("/categories/{id}", categoryHandler.GetCategory)
	router.Post("/categories", categoryHandler.CreateCategory)
	router.Patch("/categories/{id}", categoryHandler.UpdateCategory)
	router.Delete("/categories/{id}", categoryHandler.DeleteCategory)

	// define routes (orders)

	router.Get("/orders", orderHandler.GetOrders)
	router.Get("/orders/{id}", orderHandler.GetOrder)
	router.Post("/orders", orderHandler.CreateOrder)
	router.Patch("/orders/{id}", orderHandler.UpdateOrder)
	router.Delete("/orders/{id}", orderHandler.DeleteOrder)

	// Start the server
	log.Println("Server started on :3000 😂")
	log.Fatal(http.ListenAndServe(":3000", router))

}
