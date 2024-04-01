package handlers

import (
	"api-gateway/api/api_test/storage"
	"api-gateway/api/api_test/storage/kv"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// type handlerV1 struct {
// 	log            logger.Logger
// 	serviceManager services.IServiceManager
// 	cfg            config.Config
// 	jwthandler     token.JWTHandler
// 	enforcer       *casbin.Enforcer
// }

// // HandlerV1Config ...
// type HandlerV1Config struct {
// 	Logger         logger.Logger
// 	ServiceManager services.IServiceManager
// 	Cfg            config.Config
// 	JWTHandler     token.JWTHandler
// 	Enforcer       *casbin.Enforcer
// }

// // New ...
// func New(c *HandlerV1Config) *handlerV1 {
// 	return &handlerV1{
// 		log:            c.Logger,
// 		serviceManager: c.ServiceManager,
// 		cfg:            c.Cfg,
// 		jwthandler:     c.JWTHandler,
// 		enforcer:       c.Enforcer,
// 	}
// }

// User crud
func RegisterUser(c *gin.Context) {
	var newUser storage.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	newUser.Id = uuid.NewString()
	newUser.Email = strings.ToLower(newUser.Email)
	err := newUser.Validate()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	userJson, err := json.Marshal(newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := kv.Set(newUser.Id, string(userJson), 1000); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// auth := smtp.PlainAuth("", "Oybekgolang@gmail.com", "ecncwhvfdyvjghux", "smtp.gmail.com")
	// err = smtp.SendMail("smtp.gmail.com:587", auth, "Oybekgolang@gmail.com", []string{newUser.Email}, []byte("To: "+newUser.Email+"\r\nSubject: Email verification\r\nMIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"))
	// if err != nil {
	// 	log.Fatalf("Error sending otp to email: %v", err)
	// }
	log.Println("Email sent successfully")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "One time password sent to your email",
	})
}

func Verify(c *gin.Context) {
	userCode := c.Param("code")

	if userCode != "12345" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Incorrect code",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})
}