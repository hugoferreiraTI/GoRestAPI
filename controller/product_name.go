package controller

import (
	"go_mod_api/model"
	"go_mod_api/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productController struct { //user case
	productUseCase usecase.ProductUserCase
}

func NewProductControler(usecase usecase.ProductUserCase) productController {
	return productController{
		productUseCase: usecase,
	}
}

func (p *productController) GetProducts(ctx *gin.Context) {
	//products := []model.Product{ //mockando alguns produtos de momento para testar a requisição
	//	{
	//		Id:    1,
	//		Name:  "Batata frita",
	//		Price: 20,
	//	},
	//}

	products, err := p.productUseCase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, products)
}

func (p *productController) GetProductsById(ctx *gin.Context) {
	id := ctx.Param("prudctId")
	if id == "" {
		response := model.Response{
			Message: "Id do produto não pdoe ser nulo",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id) //transforma de string para int
	if err != nil {
		response := model.Response{
			Message: "ID do produto precisa ser um número",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	product, err := p.productUseCase.GetProductByID(productId) //procura no repository o ID

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if product == nil {
		response := model.Response{
			Message: "Produto não foi encontrado na base de dados",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}
	ctx.JSON(http.StatusOK, product)
}

func (p *productController) CreateProduct(ctx *gin.Context) {
	var product model.Product

	err := ctx.BindJSON(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedPRoduct, err := p.productUseCase.CreateProduct(product)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedPRoduct)

}
