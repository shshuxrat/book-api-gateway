package api

import (
	"book-api-gateway/api/docs"
	"book-api-gateway/api/handlers/v1"
	"book-api-gateway/config"
	"book-api-gateway/pkg/logger"
	"book-api-gateway/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type RouterOptions struct {
	Log      logger.Logger
	Cfg      config.Config
	Services services.ServicesI
}

// SetUpRouter godoc
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func New(opt *RouterOptions) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AllowHeaders = append(config.AllowHeaders, "*")

	docs.SwaggerInfo.Title = "Api Gateway"
	docs.SwaggerInfo.Version = "1.0"

	router.Use(cors.New(config))

	handlerV1 := handlers.NewHandler(&handlers.HandlerOptions{
		Log:      opt.Log,
		Cfg:      opt.Cfg,
		Services: opt.Services,
	})

	apiV1 := router.Group("/v1")
	//book_category
	apiV1.POST("/book_category", handlerV1.CreateBookCategory)
	apiV1.GET("/book_category", handlerV1.GetAllBookCategory)
	apiV1.GET("/book_category/:book_category_id", handlerV1.GetBookCategory)
	apiV1.PUT("/book_category", handlerV1.UpdateBookCategory)
	apiV1.DELETE("/book_category/:book_category_id", handlerV1.DeleteBookCategory)

	//book
	apiV1.POST("/book", handlerV1.CreateBook)
	apiV1.GET("/book", handlerV1.GetAllBook)
	apiV1.GET("/book/:book_id", handlerV1.GetBook)
	apiV1.PUT("/book", handlerV1.UpdateBook)
	apiV1.DELETE("/book/:book_id", handlerV1.DeleteBook)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	return router

}
