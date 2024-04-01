package api

import (
	// mid "api-gateway/api/middleware"

	// "github.com/casbin/casbin/util"
	casbinN "github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "api-gateway/api/docs"
	v1 "api-gateway/api/handlers/v1"
	"api-gateway/api/handlers/v1/tokens"
	"api-gateway/config"
	"api-gateway/pkg/logger"
	"api-gateway/services"
	"api-gateway/storage/repo"
)

type Option struct {
	Conf           config.Config
	Logger         logger.Logger
	ServiceManager services.IServiceManager
	Reds           repo.RedisStorageI
}

// @title Welcome to Library service
// @version 1.0
// @description Pro Library Api
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func New(option Option) *gin.Engine {
	casbinEnforcer, err := casbinN.NewEnforcer(option.Conf.CasbinConfigPath, option.Conf.AuthCSVPath)
	if err != nil {
		option.Logger.Error("cannot create a new enforcer", logger.Error(err))
	}
	/*_ = casbinEnforcer.LoadPolicy()

	casbinEnforcer.GetRoleManager().AddMatchingFunc("keyMatch", util.KeyMatch)
	casbinEnforcer.GetRoleManager().AddMatchingFunc("keyMatch3", util.KeyMatch3)*/

	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	jwtHandle := tokens.JWTHandler{
		SignInKey: option.Conf.SigningKey,
		Log:       option.Logger,
	}

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Logger:         option.Logger,
		ServiceManager: option.ServiceManager,
		Cfg:            option.Conf,
		Reds:           option.Reds,
		Casbin:         casbinEnforcer,
		JWTHandler:     jwtHandle,
	})

	api := router.Group("/v1")
	// api.Use(mid.NewAuth(casbinEnforcer, option.Conf))

	api.POST("/user/register", handlerV1.RegisterUser)
	api.POST("/user/verify/:email/:code", handlerV1.Verify)
	api.POST("/user/login/:email/:password", handlerV1.Login)
	api.POST("/user/create", handlerV1.CreateUser)
	api.DELETE("/user/delete/:id", handlerV1.DeleteUser)
	api.GET("/user/getall/:page/:limit", handlerV1.GetAllUsers)
	api.PUT("/user/update/:id", handlerV1.UpdateUser)

	api.POST("/book/create", handlerV1.CreateBook)
	api.PUT("/book/update/:id", handlerV1.UpdateBook)
	api.GET("/book/get/:id", handlerV1.GetBookById)
	api.DELETE("/book/delete/:id", handlerV1.DeleteBook)
	api.GET("/book/:page/:limit", handlerV1.ListBooks)

	url := ginSwagger.URL("swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}
