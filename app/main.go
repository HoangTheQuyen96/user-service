package main

import (
	"fmt"

	"github.com/HoangTheQuyen96/user-service/config"
	"github.com/HoangTheQuyen96/user-service/database/provider/redis"
	"github.com/HoangTheQuyen96/user-service/database/provider/tidb"
	"github.com/HoangTheQuyen96/user-service/domain"
	"github.com/HoangTheQuyen96/user-service/user/entrypoints/http"
	"github.com/HoangTheQuyen96/user-service/user/repository"
	"github.com/HoangTheQuyen96/user-service/user/usescases"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Cfg

	db := tidb.CreateDB(cfg.GetString("dbs.tidb.dsn"))

	redisClient := redis.CreateRedisClient(cfg.GetString("dbs.redis.address"))

	fmt.Println(redisClient)

	db.AutoMigrate(&domain.User{})

	userRepo := repository.NewUserRepository(db)

	uc := usescases.NewUserUsecase(userRepo)

	router := gin.Default()

	http.NewUserHandler(router, uc)

	router.Run(":" + cfg.GetString("server.port"))
}
