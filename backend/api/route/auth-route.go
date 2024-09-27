package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/naol86/addis-software-starter/project/backend/api/controller"
	"github.com/naol86/addis-software-starter/project/backend/config"
	"github.com/naol86/addis-software-starter/project/backend/internal/repository"
	"github.com/naol86/addis-software-starter/project/backend/internal/usecase"
	"go.mongodb.org/mongo-driver/v2/mongo"
)



func initAuthRoute(env *config.Env, timeout time.Duration, db *mongo.Database, r *gin.RouterGroup) {
	userRepository := repository.NewUserRepository(db, "users")
	userUsecase := usecase.NewUserUseCase(timeout, userRepository)
	userController := controller.UserController{
		UserUsecase: userUsecase,
		Env: env,
	}
	r.POST("/signup", userController.Signup)
	r.POST("/signin", userController.Signin)
}