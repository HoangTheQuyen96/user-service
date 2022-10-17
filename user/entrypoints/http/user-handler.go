package http

import (
	"net/http"

	"github.com/HoangTheQuyen96/user-service/domain"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	uc domain.UserUsecase
}

// router gin
func NewUserHandler(router *gin.Engine, uc domain.UserUsecase) {
	handler := &userHandler{
		uc: uc,
	}

	router.POST("/users:register", handler.Register)
}

func (h *userHandler) Register(c *gin.Context) {
	var createUserRequest domain.CreateUserRequest

	if err := c.ShouldBindJSON(&createUserRequest); err != nil {
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
