package routes

import (
	"rest3/cmd/server/handlers"
	"rest3/internal/product"
	"rest3/pkg/store"

	"github.com/gin-gonic/gin"
)

type Router struct {
	en *gin.Engine
}

func NewRouter(en *gin.Engine) *Router {
	return &Router{en: en}
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
	storage := store.NewStorage("./products1.json")
	repo := product.NewRepository(storage)
	service := product.NewService(repo)
	handler := handlers.NewProduct(service)

	products := routes.en.Group("/products")
	{
		products.GET("", handler.Get())
		products.POST("", handler.Create())
		products.GET("/:id", handler.GetById())
		products.GET("/search", handler.GetSearch())
		products.DELETE(":id", handler.Delete())
		products.PATCH(":id", handler.Patch())
		products.PUT(":id", handler.Put())
	}
}
