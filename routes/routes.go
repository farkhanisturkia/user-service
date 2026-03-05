package routes

import (
	// "go-react-vue/backend/controllers"
	// "go-react-vue/backend/middlewares"
	"learn-microservices/user-service/database"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	//initialize gin
	router := gin.Default()

	// set up CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:  []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders: []string{"Content-Length"},
	}))

	router.GET("/health", func(c *gin.Context) {
		sqlDB, err := database.DB.DB()
		if err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"status": "db instance error",
				"error":  err.Error(),
			})
			return
		}

		if err := sqlDB.Ping(); err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"status": "db not ready",
				"error":  err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, User!",
		})
	})

	// route register
	// router.POST("/api/register", controllers.Register)

	// // route login
	// router.POST("/api/login", controllers.Login)

	// // route users
	// router.GET("/api/users", controllers.FindUsers)

	// // route user create
	// router.POST("/api/users", middlewares.AuthMiddleware(), controllers.CreateUser)

	// // route user by id
	// router.GET("/api/users/:id", middlewares.AuthMiddleware(), controllers.FindUserById)

	// // route user update
	// router.PUT("/api/users/:id", middlewares.AuthMiddleware(), controllers.UpdateUser)

	// // route user delete
	// router.DELETE("/api/users/:id", middlewares.AuthMiddleware(), controllers.DeleteUser)

	// // route courses
	// router.GET("/api/courses", controllers.FindCourses)

	return router
}
