package main

import (
	"products-api/controller"
	"products-api/db"
	"products-api/repository"
	usecase "products-api/useCase"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	//Camada de banco de dados
	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	//Camada de repositórios
	ProductRepository := repository.NewProductRepository(dbConnection)

	//Camada de use cases
	ProductUseCase := usecase.NewProductsUseCase(ProductRepository)

	//Camada de controllers
	ProductController := controller.NewProductController(ProductUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", ProductController.GetProduct)
	server.POST("/products", ProductController.CreateProduct)
	server.GET("/products/:id", ProductController.GetProductByID)

	server.Run(":8080")
}
