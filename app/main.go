package main

import (
	"github.com/HoangTheQuyen96/user-service/database/provider/tidb"
	"github.com/HoangTheQuyen96/user-service/domain"
	"github.com/HoangTheQuyen96/user-service/user/entrypoints/http"
	"github.com/HoangTheQuyen96/user-service/user/repository"
	"github.com/HoangTheQuyen96/user-service/user/usescases"
	"github.com/gin-gonic/gin"
)

func main() {
	db := tidb.CreateDB()

	db.AutoMigrate(&domain.User{})

	userRepo := repository.NewUserRepository(db)

	uc := usescases.NewUserUsecase(userRepo)

	router := gin.Default()

	http.NewUserHandler(router, uc)

	router.Run(":8080")
}
