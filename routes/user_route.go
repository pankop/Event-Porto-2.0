package routes

import (
	"github.com/Caknoooo/go-gin-clean-starter/controller"
	"github.com/Caknoooo/go-gin-clean-starter/middleware"
	"github.com/Caknoooo/go-gin-clean-starter/service"
	"github.com/gin-gonic/gin"
)

func User(route *gin.Engine, userController controller.UserController, jwtService service.JWTService) {
	routes := route.Group("/api/auth")
	{
		// User
		routes.POST("", userController.Register)
		routes.POST("/login", userController.Login)
		routes.POST("/verify", userController.VerifyEmail)
		routes.POST("/sendmail", userController.SendVerificationEmail)
		routes.GET("", middleware.Authenticate(jwtService), middleware.OnlyAllow("admin"), userController.GetAllUser)
		routes.GET("/me", middleware.Authenticate(jwtService), userController.Me)
		routes.PUT("/update", middleware.Authenticate(jwtService), userController.Update)
		routes.POST("/reset", userController.ResetPassword)
		routes.POST("/forget", userController.ForgetPassword)
		routes.DELETE("", middleware.Authenticate(jwtService), userController.Delete)
	}
}
