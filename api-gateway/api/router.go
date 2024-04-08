package api

import (
	_ "api-gateway/api/docs" // swag
	v1 "api-gateway/api/handlers/v1"
	"api-gateway/api/middleware"
	token "api-gateway/api/tokens"
	// "api-gateway/kafka"

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
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func New(option Option) *gin.Engine {
	router := gin.New()

	jwtHandler := token.JWTHandler{
		SigninKey: option.Conf.SignInKey,
	}

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

	router.Use(middleware.NewAuthorizer(option.Enforcer, jwtHandler, option.Conf))
	api := router.Group("/v1")
	
	// Casbin role
	api.GET("/rbac/roles", handlerV1.ListRoles)
	api.GET("/rbac/list-role-policies", handlerV1.ListPolicies)
	api.POST("/rbac/add-user-role", handlerV1.CreateRole)

	// Registration
	api.POST("/login", handlerV1.Login)
	api.POST("/signup", handlerV1.Signup)
	api.POST("/verification", handlerV1.Verification)
	api.POST("/resetpassword", handlerV1.ResetPassword)
	api.POST("/resetcheckpassword", handlerV1.ResetCheckPassword)

	// User
	api.POST("/user", handlerV1.CreateUser)
	api.GET("/user/:id", handlerV1.ReadUser)
	api.PUT("/user/:id", handlerV1.UpdateUser)
	api.DELETE("/user/:id", handlerV1.DeleteUser)
	api.GET("/users", handlerV1.ListUsers)

	// CategoryIcon
	api.POST("/categoryicon", handlerV1.CreateCategoryIcon)
	api.GET("/categoryicon/:id", handlerV1.ReadCategoryIcon)
	api.PUT("/categoryicon/:id", handlerV1.UpdateCategoryIcon)
	api.DELETE("/categoryicon/:id", handlerV1.DeleteCategoryIcon)
	api.GET("/categoryicons", handlerV1.ListCategoryIcons)

	// Category
	api.POST("/category", handlerV1.CreateCategory)
	api.GET("/category/:id", handlerV1.ReadCategory)
	api.PUT("/category/:id", handlerV1.UpdateCategory)
	api.DELETE("/category/:id", handlerV1.DeleteCategory)
	api.GET("/categorys", handlerV1.ListCategorys)

	// Product
	api.POST("/product", handlerV1.CreateProduct)
	api.GET("/product/:id", handlerV1.ReadProduct)
	api.PUT("/product/:id", handlerV1.UpdateProduct)
	api.DELETE("/product/:id", handlerV1.DeleteProduct)
	api.GET("/products", handlerV1.ListProducts)
	api.GET("/listproductwithcomments", handlerV1.ListProductWithComment)

	// Orderproduct
	api.POST("/orderproduct", handlerV1.CreateOrderproduct)
	api.GET("/orderproduct/:id", handlerV1.ReadOrderproduct)
	api.PUT("/orderproduct/:id", handlerV1.UpdateOrderproduct)
	api.DELETE("/orderproduct/:id", handlerV1.DeleteOrderproduct)
	api.GET("/orderproducts", handlerV1.ListOrderproducts)

	// Comment
	api.POST("/comment", handlerV1.CreateComment)
	api.GET("/comment/:id", handlerV1.ReadComment)
	api.PUT("/comment", handlerV1.UpdateComment)
	api.DELETE("/comment/:id", handlerV1.DeleteComment)
	api.GET("/comments", handlerV1.ListComments)
	api.GET("/commentsbyproductid", handlerV1.ListCommentsByProductId)

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
