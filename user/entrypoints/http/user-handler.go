package http

import (
	"net/http"

	"github.com/HoangTheQuyen96/user-service/domain"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type userHandler struct {
	uc domain.UserUsecase
}

// router gin
func NewUserHandler(router *gin.Engine, uc domain.UserUsecase) {
	handler := &userHandler{
		uc: uc,
	}

	router.POST("/users:action", func(c *gin.Context) {
		action := c.Param("action")

		switch action {
		case ":register":
			handler.Register(c)
		case ":login":
			handler.Login(c)
		default:
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid action"})
		}
	})
}

func (h *userHandler) Register(c *gin.Context) {
	var createUserRequest domain.CreateUserRequest

	if err := c.ShouldBindJSON(&createUserRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()

	err := validate.Struct(createUserRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.uc.Register(c, &createUserRequest)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *userHandler) Login(c *gin.Context) {
	var loginRequest domain.LoginRequest

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()

	err := validate.Struct(loginRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	loginResponse, err := h.uc.Login(c, &loginRequest)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, loginResponse)
}
