package routes

import (
	"rest2/cmd/handlers"
	"rest2/internal/domain"
	"rest2/internal/product"

	"github.com/gin-gonic/gin"
)

type Router struct {
	db *[]domain.Product
	en *gin.Engine
}

func NewRouter(en *gin.Engine, db *[]domain.Product) *Router {
	return &Router{en: en, db: db}
}

func (routes *Router) SetRoutes() {
	routes.SetProduct()
}

func (routes *Router) SetProduct() {
	// instances
	// var productsList = []domain.Product{}
	//repo := product.NewRepository(productsList)
	//service := product.NewService(repo)
	//productHandler := handler.NewProductHandler(service)

	repo := product.NewRepository(routes.db, 3)
	service := product.NewService(repo)
	handler := handlers.NewProduct(service)

	products := routes.en.Group("/products")
	products.GET("", handler.Get())
	products.POST("", handler.Create())
	products.GET("/:id", handler.GetById())
	products.GET("/search", handler.GetSearch())
}
