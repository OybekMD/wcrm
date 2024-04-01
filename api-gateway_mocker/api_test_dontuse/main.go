package main

import (
	"api-gateway/api/api_test/handlers"
	"api-gateway/api/api_test/storage/kv"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	kv.Init(kv.NewInMemoryInst())

	router := gin.New()
	router.POST("/user", handlers.CreateUser)
	router.GET("/user", handlers.GetUser)
	router.DELETE("/user", handlers.DeleteUser)
	router.GET("/users", handlers.ListUsers)

	router.GET("/categoryicon", handlers.GetProduct)
	router.POST("/categoryicon", handlers.CreateProduct)
	router.DELETE("/categoryicon", handlers.DeleteProduct)
	router.GET("/categoryicons", handlers.ListProducts)

	router.GET("/category", handlers.GetProduct)
	router.POST("/category", handlers.CreateProduct)
	router.DELETE("/category", handlers.DeleteProduct)
	router.GET("/categorys", handlers.ListProducts)

	router.GET("/product", handlers.GetProduct)
	router.POST("/product", handlers.CreateProduct)
	router.DELETE("/product", handlers.DeleteProduct)
	router.GET("/products", handlers.ListProducts)

	router.GET("/orderproduct", handlers.GetProduct)
	router.POST("/orderproduct", handlers.CreateProduct)
	router.DELETE("/orderproduct", handlers.DeleteProduct)
	router.GET("/orderproducts", handlers.ListProducts)

	router.GET("/comment", handlers.GetProduct)
	router.POST("/comment", handlers.CreateProduct)
	router.DELETE("/comment", handlers.DeleteProduct)
	router.GET("/comments", handlers.ListProducts)

	log.Fatal(http.ListenAndServe(":9999", router))
}
