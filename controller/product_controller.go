package controller

import (
	"net/http"
	"products-api/model"
	usecase "products-api/useCase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productController struct {
	productsUseCase usecase.ProductsUseCase
}

func NewProductController(useCase usecase.ProductsUseCase) productController {
	return productController{
		productsUseCase: useCase,
	}
}

func (p *productController) GetProduct(ctx *gin.Context) {

	products, err := p.productsUseCase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"products": products,
	})
}

func (p *productController) CreateProduct(ctx *gin.Context) {
	var product model.Product
	err := ctx.BindJSON(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	insertedProduct, err := p.productsUseCase.CreateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, insertedProduct)
}

func (p *productController) GetProductByID(ctx *gin.Context) {
	id_param := ctx.Param("id")
	id, err := strconv.Atoi(id_param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	product, err := p.productsUseCase.GetProductByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if product.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Product not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, product)
}
