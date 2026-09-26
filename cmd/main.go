package main

import (
	"go_mod_api/controller"
	"go_mod_api/db"
	"go_mod_api/repository"
	"go_mod_api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	//camada de inicialização
	dbConnection, err := db.ConnectDB()

	if err != nil {
		panic(err)
	}

	//camada de repository
	ProductRepository := repository.NewProductRepository(dbConnection)

	//camada USecase
	server := gin.Default()
	ProductUseCase := usecase.NewProductUserCase(ProductRepository)

	//camada de controllers
	ProductController := controller.NewProductControler(ProductUseCase)

	server.GET("/ping", func(ctx *gin.Context) { //testando conexão com o server
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", ProductController.GetProducts)
	server.POST("/product", ProductController.CreateProduct)
	server.GET("/product/:prudctId", ProductController.GetProductsById)

	server.Run(":8000")
}
