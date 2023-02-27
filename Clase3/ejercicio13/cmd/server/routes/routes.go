package routes

import (
	"os"
	"rest3/cmd/server/handlers"
	"rest3/cmd/server/middlewares"
	"rest3/docs"
	"rest3/internal/product"
	"rest3/pkg/store"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	routes.en.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	products := routes.en.Group("/products")
	// routes.en.Use(middlewares.TokenAuthMiddleware())
	products.Use(middlewares.TokenAuthMiddleware())
	{
		products.GET("", handler.Get(), middlewares.ResponseMiddleware())
		products.POST("", handler.Create())
		products.GET("/:id", handler.GetById())
		products.GET("/search", handler.GetSearch())
		products.DELETE(":id", handler.Delete())
		products.PATCH(":id", handler.Patch())
		products.PUT(":id", handler.Put())
	}
}
