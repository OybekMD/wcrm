package api

import (
	_ "api-gateway/api/docs" // swag
	v1 "api-gateway/api/handlers/v1"

	// token "api-gateway/api/tokens"

	// "api-gateway/api/middleware"
	"api-gateway/config"
	"api-gateway/pkg/logger"
	"api-gateway/services"

	"github.com/casbin/casbin/v2"
	"github.com/gin-contrib/cors"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

// Option ...
type Option struct {
	Conf           config.Config
	Logger         logger.Logger
	ServiceManager services.IServiceManager
	Enforcer       *casbin.Enforcer
}

// @title WCRM User
// @version 0.1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func New(option Option) *gin.Engine {
	router := gin.New()

	// jwtHandler := token.JWTHandler{
	// 	SigninKey: option.Conf.SignInKey,
	// }

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Cfg:            option.Conf,
		Logger:         option.Logger,
		ServiceManager: option.ServiceManager,
		Enforcer:       option.Enforcer,
	})

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = []string{"*"}
	corsConfig.AllowBrowserExtensions = true
	corsConfig.AllowMethods = []string{"*"}
	router.Use(cors.New(corsConfig))

	// router.Use(middleware.NewAuthorizer(option.Enforcer, jwtHandler, option.Conf))
	api := router.Group("/v1")

	// Registration
	api.POST("/login", handlerV1.Login)
	api.POST("/signin", handlerV1.Signin)
	api.POST("/verification", handlerV1.Verification)
	api.POST("/resetpassword", handlerV1.ResetPassword)
	api.POST("/resetcheckpassword", handlerV1.ResetCheckPassword)

	// user
	api.POST("/user", handlerV1.CreateUser)
	api.GET("/user/:id", handlerV1.ReadUser)
	api.PUT("/user/:id", handlerV1.UpdateUser)
	api.DELETE("/user/:id", handlerV1.DeleteUser)
	api.GET("/users", handlerV1.ListUsers)

	// Fileup
	api.POST("/pdfupload", handlerV1.UploadPDFFile)
	api.POST("/imageupload", handlerV1.UploadImageFile)
	api.POST("/soundupload", handlerV1.UploadSoundFile)
	api.POST("/videoupload", handlerV1.UploadVideoFile)
	router.Static("/media", "./media")

	url := ginSwagger.URL("swagger/doc.json")
	api.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}
