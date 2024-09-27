package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/naol86/addis-software-starter/project/backend/api/middleware"
	"github.com/naol86/addis-software-starter/project/backend/config"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func SetUpRoute(env *config.Env, timeout time.Duration, db *mongo.Database, r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	protectedRoute := r.Group("/", middleware.JwtAuthMiddleWare(env.AccessTokenSecret))

	protectedRoute.GET("/protected", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Protected route",})
	})

	authRoute := r.Group("/auth")
	initAuthRoute(env, timeout, db, authRoute)


}