package main

import (
	"os"

	"github.com/boolow5/GinBoilerPlate/conf"
	"github.com/boolow5/GinBoilerPlate/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	conf.InitConfig()
	port := os.Getenv("PORT")
	r := gin.Default() //.New()
	//r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Use(conf.CORSMiddleware())

	if port == "" {
		port = "8000"
	}

	// the jwt middleware
	adminAuthMiddleware := conf.NewAuthMiddleware(5)
	L1Auth := conf.NewAuthMiddleware(1)
	L2Auth := conf.NewAuthMiddleware(2)
	L3Auth := conf.NewAuthMiddleware(3)
	L4Auth := conf.NewAuthMiddleware(4)
	L5Auth := conf.NewAuthMiddleware(5)

	// routes
	r.POST("/login", adminAuthMiddleware.LoginHandler)

	api := r.Group("/api/v1")

	// routes that need authentication
	authorized1 := api.Group("/l1")
	authorized2 := api.Group("/l2")
	authorized3 := api.Group("/l3")
	authorized4 := api.Group("/l4")
	authorized5 := api.Group("/l5")
	authorized1.Use(L1Auth.MiddlewareFunc())
	authorized2.Use(L2Auth.MiddlewareFunc())
	authorized3.Use(L3Auth.MiddlewareFunc())
	authorized4.Use(L4Auth.MiddlewareFunc())
	authorized5.Use(L5Auth.MiddlewareFunc())
	{
		authorized5.GET("/refresh_token", adminAuthMiddleware.RefreshHandler)
		// index endpoint
		authorized5.GET("/", controllers.Index)
		// role endpoints
		authorized2.GET("/roles", controllers.AllRoles)
		authorized2.POST("/role", controllers.AddRole)
		authorized2.PUT("/role", controllers.UpdateRole)
		authorized2.DELETE("/role", controllers.DeleteRole)
		// users endpoints
		authorized3.GET("/users", controllers.AllUsers)
		authorized3.GET("/user", controllers.GetUser)
		authorized2.POST("/user", controllers.AddUser)
		authorized2.PUT("/user", controllers.UpdateUser)
		authorized3.PUT("/user/password", controllers.UpdateUserPassword)
		authorized2.DELETE("/user", controllers.DeleteUser)
	}

	//endless.ListenAndServe(":"+port, r)
	r.Run(":8080")
}
