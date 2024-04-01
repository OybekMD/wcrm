package api

import (
	mid "api-gateway/api/middleware"

	"github.com/casbin/casbin/v2"
	"api-gateway/api/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "api-gateway/api/docs"
	v1 "api-gateway/api/handlers/v1"
	token "api-gateway/api/tokens"
	"api-gateway/config"
	"api-gateway/pkg/logger"
	"api-gateway/services"
	"api-gateway/storage/repo"
)

type Option struct {
	Conf           config.Config
	Logger         logger.Logger
	ServiceManager services.IServiceManager
	Enforcer       *casbin.Enforcer
	Reds           repo.RedisStorageI
}

// @title Welcome to User-Product service
// @version 1.0
// @description This code is written by Nodirbek in third mont exam GOLANG
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func New(option Option) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	jwtHandler := token.JWTHandler{
		SigninKey: option.Conf.SignInKey,
	}

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Logger:         option.Logger,
		ServiceManager: option.ServiceManager,
		Cfg:            option.Conf,
		Enforcer:       option.Enforcer,
		JWTHandler:     jwtHandler,
	})

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = []string{"*"}
	corsConfig.AllowBrowserExtensions = true
	corsConfig.AllowMethods = []string{"*"}
	router.Use(cors.New(corsConfig))

	router.Use(middleware.NewAuthorizer(option.Enforcer, jwtHandler, option.Conf))
	api := router.Group("/v1")

	// Registration
	api.POST("/signup", handlerV1.Signup)
	api.POST("/login", handlerV1.Login)
	api.POST("/verification", handlerV1.Verification)
	api.POST("/resetpassword", handlerV1.ResetPassword)
	api.POST("/resetcheckpassword", handlerV1.ResetCheckPassword)

	api.POST("/user", handlerV1.MockListUsers)
	api.GET("/user", handlerV1.MockReadUser)
	api.PUT("/user", handlerV1.MockUpdateUser)
	api.DELETE("/user/:id", handlerV1.MockDeleteUser)
	api.GET("/users", handlerV1.MockListUsers)

	api.POST("/product/create", handlerV1.CreateProduct)
	api.PUT("/product/update/:id", handlerV1.UpdateProduct)
	api.GET("/product/get/:id", handlerV1.GetProductById)
	api.DELETE("/product/delete/:id", handlerV1.DeleteProduct)
	// api.GET("/products/get/:id", handlerV1.GetPurchasedProductsByUserId)
	api.GET("/product/:page/:limit", handlerV1.ListProducts)
	// api.POST("/product/buy", handlerV1.BuyProduct)

	api.GET("/admin/:username/:password", handlerV1.GenerateAccessTokenForAdmin)

	url := ginSwagger.URL("swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}
